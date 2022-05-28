package main

import (
	"github.com/anon55555/mt"
	"github.com/yuin/gopher-lua"
	"reflect"
	"time"
)

func doPoll(l *lua.LState, clients []*Client) (*Client, *mt.Pkt, bool) {
	var timeout time.Duration
	hasTimeout := false
	if l.GetTop() > 1 {
		timeout = time.Duration(float64(l.ToNumber(2)) * float64(time.Second))
		hasTimeout = true
	}

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
		return nil, nil, false
	}

	cases = append(cases, reflect.SelectCase{
		Dir:  reflect.SelectRecv,
		Chan: reflect.ValueOf(signalChannel()),
	})

	if hasTimeout {
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(time.After(timeout)),
		})
	}

	idx, value, ok := reflect.Select(cases)

	if idx >= offset {
		return nil, nil, true
	}

	client := clients[idx]

	var pkt *mt.Pkt = nil
	if ok {
		pkt = value.Interface().(*mt.Pkt)
	} else {
		client.state = csDisconnected
	}

	return client, pkt, false
}
