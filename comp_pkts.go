package main

import (
	"github.com/dragonfireclient/hydra-dragonfire/convert"
	"github.com/dragonfireclient/mt"
	"github.com/yuin/gopher-lua"
	"sync"
)

type CompPkts struct {
	client     *Client
	mu         sync.Mutex
	wildcard   bool
	subscribed map[string]struct{}
	userdata   *lua.LUserData
}

var compPktsFuncs = map[string]lua.LGFunction{
	"subscribe":   l_comp_pkts_subscribe,
	"unsubscribe": l_comp_pkts_unsubscribe,
	"wildcard":    l_comp_pkts_wildcard,
}

type EventPkt struct {
	pktType string
	pktData *mt.Pkt
}

func (evt EventPkt) handle(l *lua.LState, val lua.LValue) {
	l.SetField(val, "type", lua.LString("pkt"))
	l.SetField(val, "pkt_type", lua.LString(evt.pktType))
	l.SetField(val, "pkt_data", convert.PushPkt(l, evt.pktData))
}

func getCompPkts(l *lua.LState) *CompPkts {
	return l.CheckUserData(1).Value.(*CompPkts)
}

func (comp *CompPkts) create(client *Client, l *lua.LState) {
	comp.client = client
	comp.wildcard = false
	comp.subscribed = map[string]struct{}{}
	comp.userdata = l.NewUserData()
	comp.userdata.Value = comp
	l.SetMetatable(comp.userdata, l.GetTypeMetatable("hydra.comp.pkts"))
}

func (comp *CompPkts) push() lua.LValue {
	return comp.userdata
}

func (comp *CompPkts) connect() {
}

func (comp *CompPkts) process(pkt *mt.Pkt) {
	pktType := string(convert.PushPktType(pkt))

	comp.mu.Lock()
	_, subscribed := comp.subscribed[pktType]
	comp.mu.Unlock()

	if subscribed || comp.wildcard {
		comp.client.queue <- EventPkt{pktType: pktType, pktData: pkt}
	}
}

func l_comp_pkts_subscribe(l *lua.LState) int {
	comp := getCompPkts(l)
	n := l.GetTop()

	comp.mu.Lock()
	defer comp.mu.Unlock()

	for i := 2; i <= n; i++ {
		comp.subscribed[l.CheckString(i)] = struct{}{}
	}

	return 0
}

func l_comp_pkts_unsubscribe(l *lua.LState) int {
	comp := getCompPkts(l)
	n := l.GetTop()

	comp.mu.Lock()
	defer comp.mu.Unlock()

	for i := 2; i <= n; i++ {
		delete(comp.subscribed, l.CheckString(i))
	}

	return 0
}

func l_comp_pkts_wildcard(l *lua.LState) int {
	comp := getCompPkts(l)
	comp.wildcard = l.ToBool(2)
	return 0
}
