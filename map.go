package main

import (
	"github.com/anon55555/mt"
	"github.com/dragonfireclient/hydra-dragonfire/convert"
	"github.com/yuin/gopher-lua"
	"sync"
)

type Map struct {
	client   *Client
	mu       sync.Mutex
	blocks   map[[3]int16]*mt.MapBlk
	userdata *lua.LUserData
}

var mapFuncs = map[string]lua.LGFunction{
	"clear": l_map_clear,
	"block": l_map_block,
	"node":  l_map_node,
}

func getMap(l *lua.LState) *Map {
	return l.CheckUserData(1).Value.(*Map)
}

func (mtmap *Map) create(client *Client, l *lua.LState) {
	mtmap.client = client
	mtmap.blocks = map[[3]int16]*mt.MapBlk{}
	mtmap.userdata = l.NewUserData()
	mtmap.userdata.Value = mtmap
	l.SetMetatable(mtmap.userdata, l.GetTypeMetatable("hydra.map"))
}

func (mtmap *Map) push() lua.LValue {
	return mtmap.userdata
}

func (mtmap *Map) connect() {
}

func (mtmap *Map) process(pkt *mt.Pkt) {
	switch cmd := pkt.Cmd.(type) {
	case *mt.ToCltBlkData:
		mtmap.mu.Lock()
		mtmap.blocks[cmd.Blkpos] = &cmd.Blk
		mtmap.client.conn.SendCmd(&mt.ToSrvGotBlks{Blks: [][3]int16{cmd.Blkpos}})
		mtmap.mu.Unlock()
	}
}

func l_map_clear(l *lua.LState) int {
	mtmap := getMap(l)

	mtmap.mu.Lock()
	defer mtmap.mu.Unlock()

	var cmd mt.ToSrvDeletedBlks
	for pos := range mtmap.blocks {
		cmd.Blks = append(cmd.Blks, pos)
	}

	mtmap.blocks = map[[3]int16]*mt.MapBlk{}

	mtmap.client.conn.SendCmd(&cmd)

	return 0
}

func l_map_block(l *lua.LState) int {
	mtmap := getMap(l)
	var blkpos [3]int16
	convert.ReadVec3Int16(l, l.Get(2), &blkpos)

	mtmap.mu.Lock()
	defer mtmap.mu.Unlock()

	block, ok := mtmap.blocks[blkpos]
	if ok {
		l.Push(convert.PushMapBlk(l, *block))
	} else {
		l.Push(lua.LNil)
	}

	return 1
}

func l_map_node(l *lua.LState) int {
	mtmap := getMap(l)

	var pos [3]int16
	convert.ReadVec3Int16(l, l.Get(2), &pos)
	blkpos, i := mt.Pos2Blkpos(pos)

	mtmap.mu.Lock()
	defer mtmap.mu.Unlock()

	block, block_exists := mtmap.blocks[blkpos]
	if block_exists {
		meta, meta_exists := block.NodeMetas[i]
		if !meta_exists {
			meta = &mt.NodeMeta{}
		}

		lnode := l.NewTable()
		l.SetField(lnode, "param0", lua.LNumber(block.Param0[i]))
		l.SetField(lnode, "param1", lua.LNumber(block.Param1[i]))
		l.SetField(lnode, "param2", lua.LNumber(block.Param2[i]))
		l.SetField(lnode, "meta", convert.PushNodeMeta(l, *meta))
		l.Push(lnode)
	} else {
		l.Push(lua.LNil)
	}

	return 1
}
