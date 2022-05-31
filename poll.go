package main

import (
	"github.com/yuin/gopher-lua"
	"reflect"
	"time"
)

type Event interface {
	handle(l *lua.LState, val lua.LValue)
}

type EventTimeout struct{}

func (evt EventTimeout) handle(l *lua.LState, val lua.LValue) {
	l.SetField(val, "type", lua.LString("timeout"))
}

type EventInterrupt struct{}

func (evt EventInterrupt) handle(l *lua.LState, val lua.LValue) {
	l.SetField(val, "type", lua.LString("interrupt"))
}

func doPoll(l *lua.LState, clients []*Client) int {
	cases := make([]reflect.SelectCase, 0, len(clients)+2)

	for _, client := range clients {
		if client.state != csConnected {
			continue
		}

		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(client.queue),
		})
	}

	offset := len(cases)

	if offset < 1 {
		return 0
	}

	cases = append(cases, reflect.SelectCase{
		Dir:  reflect.SelectRecv,
		Chan: reflect.ValueOf(signalChannel),
	})

	if l.GetTop() > 1 {
		timeout := time.After(time.Duration(float64(l.ToNumber(2)) * float64(time.Second)))

		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(timeout),
		})
	}

	idx, value, _ := reflect.Select(cases)

	var evt Event
	tbl := l.NewTable()

	if idx > offset {
		evt = EventTimeout{}
	} else if idx == offset {
		evt = EventInterrupt{}
	} else {
		evt = value.Interface().(Event)
		l.SetField(tbl, "client", clients[idx].userdata)
	}

	evt.handle(l, tbl)

	l.Push(tbl)
	return 1
}
