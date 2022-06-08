package main

import (
	"github.com/dragonfireclient/mt"
	"github.com/yuin/gopher-lua"
)

type CompMap struct {
	client   *Client
	mapdata  *Map
	userdata *lua.LUserData
}

var compMapFuncs = map[string]lua.LGFunction{
	"get": l_comp_map_get,
	"set": l_comp_map_set,
}

func getCompMap(l *lua.LState) *CompMap {
	return l.CheckUserData(1).Value.(*CompMap)
}

func (comp *CompMap) create(client *Client, l *lua.LState) {
	comp.client = client
	comp.mapdata = newMap(l)
	comp.userdata = l.NewUserData()
	comp.userdata.Value = comp
	l.SetMetatable(comp.userdata, l.GetTypeMetatable("hydra.comp.map"))
}

func (comp *CompMap) push() lua.LValue {
	return comp.userdata
}

func (comp *CompMap) connect() {
}

func (comp *CompMap) process(pkt *mt.Pkt) {
	comp.mapdata.process(comp.client, pkt)
}

func l_comp_map_set(l *lua.LState) int {
	comp := getCompMap(l)
	comp.mapdata = getMap(l, 2)
	return 0
}

func l_comp_map_get(l *lua.LState) int {
	comp := getCompMap(l)
	l.Push(comp.mapdata.userdata)
	return 1
}
