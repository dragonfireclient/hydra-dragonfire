package main

import (
	"github.com/Shopify/go-lua"
	"github.com/anon55555/mt"
	"reflect"
	"time"
)

func l_poll(l *lua.State) int {
	clients := make([]*Client, 0)

	lua.CheckType(l, 1, lua.TypeTable)
	i := 1
	for {
		l.RawGetInt(1, i)
		if l.IsNil(-1) {
			l.Pop(1)
			break
		}

		clients = append(clients, l.ToUserData(-1).(*Client))
		i++
	}

	var timeout time.Duration
	hasTimeout := false
	if l.IsNumber(3) {
		timeout = time.Duration(lua.CheckNumber(l, 3) * float64(time.Second))
		hasTimeout = true
	}

	for {
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
			l.PushBoolean(false)
			return 1
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
			l.PushBoolean(true)
			return 1
		}

		client := clients[idx]

		var pkt *mt.Pkt = nil
		if ok {
			pkt = value.Interface().(*mt.Pkt)
		} else {
			client.state = csDisconnected
		}

		for _, handler := range client.handlers {
			handler.handle(pkt, l, idx+1)
		}
	}

	panic("impossible")
	return 0
}
