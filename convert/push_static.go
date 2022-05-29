package convert

import (
	"github.com/anon55555/mt"
	"github.com/yuin/gopher-lua"
	"image/color"
)

//go:generate ./push_mkauto.lua

func vec2(l *lua.LState, val [2]lua.LNumber) {
	l.Push(l.GetGlobal("vec2"))
	l.Push(val[0])
	l.Push(val[1])
	l.Call(2, 1)
}

func vec3(l *lua.LState, val [3]lua.LNumber) {
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

func pushVec2(l *lua.LState, val [2]lua.LNumber) lua.LValue {
	vec2(l, val)
	return popValue(l)
}

func pushVec3(l *lua.LState, val [3]lua.LNumber) lua.LValue {
	vec3(l, val)
	return popValue(l)
}

func pushBox1(l *lua.LState, val [2]lua.LNumber) lua.LValue {
	l.Push(l.GetGlobal("box"))
	l.Push(val[0])
	l.Push(val[1])
	l.Call(2, 1)
	return popValue(l)
}

func pushBox2(l *lua.LState, val [2][2]lua.LNumber) lua.LValue {
	l.Push(l.GetGlobal("box"))
	vec2(l, val[0])
	vec2(l, val[1])
	l.Call(2, 1)
	return popValue(l)
}

func pushBox3(l *lua.LState, val [2][3]lua.LNumber) lua.LValue {
	l.Push(l.GetGlobal("box"))
	vec3(l, val[0])
	vec3(l, val[1])
	l.Call(2, 1)
	return popValue(l)
}

func pushColor(l *lua.LState, val color.NRGBA) lua.LValue {
	tbl := l.NewTable()
	l.SetField(tbl, "r", lua.LNumber(val.R))
	l.SetField(tbl, "g", lua.LNumber(val.G))
	l.SetField(tbl, "b", lua.LNumber(val.B))
	l.SetField(tbl, "a", lua.LNumber(val.A))
	return tbl
}

func pushStringSet(l *lua.LState, val []string) lua.LValue {
	tbl := l.NewTable()
	for _, str := range val {
		l.SetField(tbl, str, lua.LTrue)
	}
	return tbl
}

func stringList[T ~string](l *lua.LState, val []T) lua.LValue {
	tbl := l.NewTable()
	for _, s := range val {
		tbl.Append(lua.LString(s))
	}
	return tbl
}

func pushStringList(l *lua.LState, val []string) lua.LValue {
	return stringList[string](l, val)
}

func pushTextureList(l *lua.LState, val []mt.Texture) lua.LValue {
	return stringList[mt.Texture](l, val)
}
