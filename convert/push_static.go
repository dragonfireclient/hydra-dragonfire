package convert

import (
	"github.com/dragonfireclient/mt"
	"github.com/yuin/gopher-lua"
	"image/color"
)

//go:generate ./push_mkauto.lua

func CreateVec2(l *lua.LState, val [2]lua.LNumber) {
	l.Push(l.GetGlobal("vec2"))
	l.Push(val[0])
	l.Push(val[1])
	l.Call(2, 1)
}

func CreateVec3(l *lua.LState, val [3]lua.LNumber) {
	l.Push(l.GetGlobal("vec3"))
	l.Push(val[0])
	l.Push(val[1])
	l.Push(val[2])
	l.Call(3, 1)
}

func popValue(l *lua.LState) lua.LValue {
	ret := l.Get(-1)
	l.Pop(1)
	return ret
}

func PushVec2(l *lua.LState, val [2]lua.LNumber) lua.LValue {
	CreateVec2(l, val)
	return popValue(l)
}

func PushVec3(l *lua.LState, val [3]lua.LNumber) lua.LValue {
	CreateVec3(l, val)
	return popValue(l)
}

func PushBox1(l *lua.LState, val [2]lua.LNumber) lua.LValue {
	l.Push(l.GetGlobal("box"))
	l.Push(val[0])
	l.Push(val[1])
	l.Call(2, 1)
	return popValue(l)
}

func PushBox2(l *lua.LState, val [2][2]lua.LNumber) lua.LValue {
	l.Push(l.GetGlobal("box"))
	CreateVec2(l, val[0])
	CreateVec2(l, val[1])
	l.Call(2, 1)
	return popValue(l)
}

func PushBox3(l *lua.LState, val [2][3]lua.LNumber) lua.LValue {
	l.Push(l.GetGlobal("box"))
	CreateVec3(l, val[0])
	CreateVec3(l, val[1])
	l.Call(2, 1)
	return popValue(l)
}

func PushColor(l *lua.LState, val color.NRGBA) lua.LValue {
	tbl := l.NewTable()
	l.SetField(tbl, "r", lua.LNumber(val.R))
	l.SetField(tbl, "g", lua.LNumber(val.G))
	l.SetField(tbl, "b", lua.LNumber(val.B))
	l.SetField(tbl, "a", lua.LNumber(val.A))
	return tbl
}

func PushStringSet(l *lua.LState, val []string) lua.LValue {
	tbl := l.NewTable()
	for _, str := range val {
		l.SetField(tbl, str, lua.LTrue)
	}
	return tbl
}

func PushStringList[T ~string](l *lua.LState, val []T) lua.LValue {
	tbl := l.NewTable()
	for _, s := range val {
		tbl.Append(lua.LString(s))
	}
	return tbl
}

func Push4096[T uint8 | mt.Content](l *lua.LState, val [4096]T) lua.LValue {
	tbl := l.NewTable()
	for i, v := range val {
		l.RawSetInt(tbl, i, lua.LNumber(v))
	}
	return tbl
}

func PushFields(l *lua.LState, val []mt.Field) lua.LValue {
	tbl := l.NewTable()
	for _, pair := range val {
		l.SetField(tbl, pair.Name, lua.LString(pair.Value))
	}
	return tbl
}

func PushNodeMetaFields(l *lua.LState, val []mt.NodeMetaField) lua.LValue {
	tbl := l.NewTable()
	for _, pair := range val {
		l.SetField(tbl, pair.Name, lua.LString(pair.Value))
	}
	return tbl
}

func PushInv(l *lua.LState, val mt.Inv) lua.LValue {
	linv := l.NewTable()
	for _, list := range val {
		llist := l.NewTable()
		l.SetField(llist, "width", lua.LNumber(list.Width))
		for _, stack := range list.Stacks {
			lmeta := l.NewTable()
			l.SetField(lmeta, "fields", PushFields(l, stack.Fields()))
			if toolcaps, ok := stack.ToolCaps(); ok {
				l.SetField(lmeta, "tool_caps", PushToolCaps(l, toolcaps))
			}

			lstack := l.NewTable()
			l.SetField(lstack, "name", lua.LString(stack.Name))
			l.SetField(lstack, "count", lua.LNumber(stack.Count))
			l.SetField(lstack, "wear", lua.LNumber(stack.Wear))
			l.SetField(lstack, "meta", lmeta)

			llist.Append(lstack)
		}
		l.SetField(linv, list.Name, llist)
	}
	return linv
}

func PushNodeMetas(l *lua.LState, val map[uint16]*mt.NodeMeta) lua.LValue {
	tbl := l.NewTable()
	for i, meta := range val {
		l.RawSetInt(tbl, int(i), PushNodeMeta(l, *meta))
	}
	return tbl
}

func PushChangedNodeMetas(l *lua.LState, val map[[3]int16]*mt.NodeMeta) lua.LValue {
	lmetas := l.NewTable()
	for pos, meta := range val {
		lmeta := l.NewTable()
		l.SetField(lmeta, "pos", PushVec3(l, [3]lua.LNumber{lua.LNumber(pos[0]), lua.LNumber(pos[1]), lua.LNumber(pos[2])}))
		l.SetField(lmeta, "meta", PushNodeMeta(l, *meta))
		lmetas.Append(lmeta)
	}
	return lmetas
}

func PushGroups(l *lua.LState, val []mt.Group) lua.LValue {
	tbl := l.NewTable()
	for _, group := range val {
		l.SetField(tbl, group.Name, lua.LNumber(group.Rating))
	}
	return tbl
}

func PushGroupCaps(l *lua.LState, val []mt.ToolGroupCap) lua.LValue {
	lcaps := l.NewTable()
	for _, groupcap := range val {
		ltimes := l.NewTable()
		for _, digtime := range groupcap.Times {
			l.RawSetInt(ltimes, int(digtime.Rating), lua.LNumber(digtime.Time))
		}

		lcap := l.NewTable()
		l.SetField(lcap, "uses", lua.LNumber(groupcap.Uses))
		l.SetField(lcap, "max_lvl", lua.LNumber(groupcap.MaxLvl))
		l.SetField(lcap, "times", ltimes)

		l.SetField(lcaps, groupcap.Name, lcap)
	}
	return lcaps
}
