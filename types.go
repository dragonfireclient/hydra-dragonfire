package main

import (
	"github.com/Shopify/go-lua"
	"github.com/anon55555/mt"
	"image/color"
)

func luaPushVec2(l *lua.State, val [2]float64) {
	l.Global("vec2")
	l.PushNumber(val[0])
	l.PushNumber(val[1])
	l.Call(2, 1)
}

func luaPushVec3(l *lua.State, val [3]float64) {
	l.Global("vec3")
	l.PushNumber(val[0])
	l.PushNumber(val[1])
	l.PushNumber(val[2])
	l.Call(3, 1)
}

func luaPushBox1(l *lua.State, val [2]float64) {
	l.Global("box")
	l.PushNumber(val[0])
	l.PushNumber(val[1])
	l.Call(2, 1)
}

func luaPushBox2(l *lua.State, val [2][2]float64) {
	l.Global("box")
	luaPushVec2(l, val[0])
	luaPushVec2(l, val[1])
	l.Call(2, 1)
}

func luaPushBox3(l *lua.State, val [2][3]float64) {
	l.Global("box")
	luaPushVec3(l, val[0])
	luaPushVec3(l, val[1])
	l.Call(2, 1)
}

func luaPushColor(l *lua.State, val color.NRGBA) {
	l.NewTable()
	l.PushInteger(int(val.R))
	l.SetField(-2, "r")
	l.PushInteger(int(val.G))
	l.SetField(-2, "g")
	l.PushInteger(int(val.B))
	l.SetField(-2, "b")
	l.PushInteger(int(val.A))
	l.SetField(-2, "a")
}

func luaPushStringSet(l *lua.State, val []string) {
	l.NewTable()
	for _, str := range val {
		l.PushBoolean(true)
		l.SetField(-2, str)
	}
}

func luaPushStringList(l *lua.State, val []string) {
	l.NewTable()
	for i, str := range val {
		l.PushString(str)
		l.RawSetInt(-2, i+1)
	}
}

// i hate go for making me do this instead of just using luaPushStringList
// but i dont want to make an unsafe cast either
func luaPushTextureList(l *lua.State, val []mt.Texture) {
	l.NewTable()
	for i, str := range val {
		l.PushString(string(str))
		l.RawSetInt(-2, i+1)
	}
}
