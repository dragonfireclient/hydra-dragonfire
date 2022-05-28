package main

import (
	"github.com/Shopify/go-lua"
	"github.com/anon55555/mt"
)

type Callbacks struct {
	wildcard   bool
	subscribed map[string]struct{}
}

func getCallbacks(l *lua.State) *Callbacks {
	return lua.CheckUserData(l, 1, "hydra.callbacks").(*Callbacks)
}

func (handler *Callbacks) create(client *Client) {
	handler.subscribed = map[string]struct{}{}
}

func (handler *Callbacks) push(l *lua.State) {
	l.PushUserData(handler)

	if lua.NewMetaTable(l, "hydra.callbacks") {
		lua.NewLibrary(l, []lua.RegistryFunction{
			{Name: "wildcard", Function: l_callbacks_wildcard},
			{Name: "subscribe", Function: l_callbacks_subscribe},
			{Name: "unsubscribe", Function: l_callbacks_unsubscribe},
		})
		l.SetField(-2, "__index")
	}
	l.SetMetaTable(-2)
}

func (handler *Callbacks) canConnect() (bool, string) {
	return true, ""
}

func (handler *Callbacks) connect() {
}

func (handler *Callbacks) handle(pkt *mt.Pkt, l *lua.State, idx int) {
	if !handler.wildcard && pkt != nil {
		if _, exists := handler.subscribed[pktToString(pkt)]; !exists {
			return
		}
	}

	if !l.IsFunction(2) {
		return
	}

	l.PushValue(2)      // callback
	l.RawGetInt(1, idx) // arg 1: client
	pktToLua(l, pkt)    // arg 2: pkt
	l.Call(2, 0)
}

func l_callbacks_wildcard(l *lua.State) int {
	handler := getCallbacks(l)
	handler.wildcard = l.ToBoolean(2)
	return 0
}

func l_callbacks_subscribe(l *lua.State) int {
	handler := getCallbacks(l)

	n := l.Top()
	for i := 2; i <= n; i++ {
		handler.subscribed[lua.CheckString(l, i)] = struct{}{}
	}

	return 0
}

func l_callbacks_unsubscribe(l *lua.State) int {
	handler := getCallbacks(l)

	n := l.Top()
	for i := 2; i <= n; i++ {
		delete(handler.subscribed, lua.CheckString(l, i))
	}

	return 0
}
