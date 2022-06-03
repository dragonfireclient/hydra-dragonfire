package main

import (
	"github.com/anon55555/mt"
	"github.com/dragonfireclient/hydra-dragonfire/convert"
	"github.com/yuin/gopher-lua"
	"sync"
)

type MapBlk struct {
	data  *mt.MapBlk
	edges [6]*PathfindEdge
}

type Map struct {
	mu       sync.Mutex
	pathfind bool
	blocks   map[[3]int16]*MapBlk
	userdata *lua.LUserData
}

var mapFuncs = map[string]lua.LGFunction{
	"block": l_map_block,
	"node":  l_map_node,
}

func getMap(l *lua.LState, idx int) *Map {
	return l.CheckUserData(idx).Value.(*Map)
}

func newMap(l *lua.LState) *Map {
	mp := &Map{}
	mp.blocks = map[[3]int16]*MapBlk{}
	mp.userdata = l.NewUserData()
	mp.userdata.Value = mp
	l.SetMetatable(mp.userdata, l.GetTypeMetatable("hydra.map"))
	return mp
}

func (mp *Map) process(client *Client, pkt *mt.Pkt) {
	switch cmd := pkt.Cmd.(type) {
	case *mt.ToCltBlkData:
		go func() {
			mp.mu.Lock()
			defer mp.mu.Unlock()

			blk := &MapBlk{}
			blk.data = &cmd.Blk
			if mp.pathfind {
				pfPreprocess(mp, cmd.Blkpos, blk)
			}

			mp.blocks[cmd.Blkpos] = blk
		}()

		client.conn.SendCmd(&mt.ToSrvGotBlks{Blks: [][3]int16{cmd.Blkpos}})
	}
}

func l_map(l *lua.LState) int {
	mp := newMap(l)
	mp.pathfind = l.ToBool(1)
	l.Push(mp.userdata)
	return 1
}

func l_map_block(l *lua.LState) int {
	mp := getMap(l, 1)
	var blkpos [3]int16
	convert.ReadVec3Int16(l, l.Get(2), &blkpos)

	mp.mu.Lock()
	defer mp.mu.Unlock()

	blk, ok := mp.blocks[blkpos]
	if ok {
		l.Push(convert.PushMapBlk(l, *blk.data))
	} else {
		l.Push(lua.LNil)
	}

	return 1
}

func l_map_node(l *lua.LState) int {
	mp := getMap(l, 1)

	var pos [3]int16
	convert.ReadVec3Int16(l, l.Get(2), &pos)
	blkpos, i := mt.Pos2Blkpos(pos)

	mp.mu.Lock()
	defer mp.mu.Unlock()

	blk, blk_ok := mp.blocks[blkpos]
	if blk_ok {
		meta, meta_ok := blk.data.NodeMetas[i]
		if !meta_ok {
			meta = &mt.NodeMeta{}
		}

		lnode := l.NewTable()
		l.SetField(lnode, "param0", lua.LNumber(blk.data.Param0[i]))
		l.SetField(lnode, "param1", lua.LNumber(blk.data.Param1[i]))
		l.SetField(lnode, "param2", lua.LNumber(blk.data.Param2[i]))
		l.SetField(lnode, "meta", convert.PushNodeMeta(l, *meta))
		l.Push(lnode)
	} else {
		l.Push(lua.LNil)
	}

	return 1
}
