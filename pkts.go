package main

import (
	"github.com/anon55555/mt"
	"github.com/dragonfireclient/hydra-dragonfire/convert"
	"github.com/yuin/gopher-lua"
	"sync"
)

type Pkts struct {
	client     *Client
	mu         sync.Mutex
	wildcard   bool
	subscribed map[string]struct{}
	userdata   *lua.LUserData
}

var pktsFuncs = map[string]lua.LGFunction{
	"subscribe":   l_pkts_subscribe,
	"unsubscribe": l_pkts_unsubscribe,
	"wildcard":    l_pkts_wildcard,
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

func getPkts(l *lua.LState) *Pkts {
	return l.CheckUserData(1).Value.(*Pkts)
}

func (pkts *Pkts) create(client *Client, l *lua.LState) {
	pkts.client = client
	pkts.wildcard = false
	pkts.subscribed = map[string]struct{}{}
	pkts.userdata = l.NewUserData()
	pkts.userdata.Value = pkts
	l.SetMetatable(pkts.userdata, l.GetTypeMetatable("hydra.pkts"))
}

func (pkts *Pkts) push() lua.LValue {
	return pkts.userdata
}

func (pkts *Pkts) connect() {
}

func (pkts *Pkts) process(pkt *mt.Pkt) {
	pktType := string(convert.PushPktType(pkt))

	pkts.mu.Lock()
	_, subscribed := pkts.subscribed[pktType]
	pkts.mu.Unlock()

	if subscribed || pkts.wildcard {
		pkts.client.queue <- EventPkt{pktType: pktType, pktData: pkt}
	}
}

func l_pkts_subscribe(l *lua.LState) int {
	pkts := getPkts(l)
	n := l.GetTop()

	pkts.mu.Lock()
	defer pkts.mu.Unlock()

	for i := 2; i <= n; i++ {
		pkts.subscribed[l.CheckString(i)] = struct{}{}
	}

	return 0
}

func l_pkts_unsubscribe(l *lua.LState) int {
	pkts := getPkts(l)
	n := l.GetTop()

	pkts.mu.Lock()
	defer pkts.mu.Unlock()

	for i := 2; i <= n; i++ {
		delete(pkts.subscribed, l.CheckString(i))
	}

	return 0
}

func l_pkts_wildcard(l *lua.LState) int {
	pkts := getPkts(l)
	pkts.wildcard = l.ToBool(2)
	return 0
}
