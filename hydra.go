package main

import (
	_ "embed"
	"github.com/yuin/gopher-lua"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var lastTime = time.Now()
var signalChannel chan os.Signal

var serializeVer uint8 = 29
var protoVer uint16 = 40

//go:embed builtin/luax/init.lua
var builtinLuaX string

//go:embed builtin/vector.lua
var builtinVector string

//go:embed builtin/escapes.lua
var builtinEscapes string

//go:embed builtin/client.lua
var builtinClient string

//go:embed builtin/base64.lua
var builtinBase64 string

var builtinFiles = []string{
	builtinLuaX,
	builtinVector,
	builtinEscapes,
	builtinClient,
	builtinBase64,
}

var hydraFuncs = map[string]lua.LGFunction{
	"client": l_client,
	"map":    l_map,
	"dtime":  l_dtime,
	"poll":   l_poll,
	"close":  l_close,
}

func l_dtime(l *lua.LState) int {
	l.Push(lua.LNumber(time.Since(lastTime).Seconds()))
	lastTime = time.Now()
	return 1
}

func l_poll(l *lua.LState) int {
	return doPoll(l, getClients(l))
}

func l_close(l *lua.LState) int {
	for _, client := range getClients(l) {
		client.closeConn()
	}

	return 0
}

func main() {
	if len(os.Args) < 2 {
		panic("missing filename")
	}

	signalChannel = make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)

	l := lua.NewState()
	defer l.Close()

	arg := l.NewTable()
	for i, a := range os.Args {
		l.RawSetInt(arg, i-1, lua.LString(a))
	}
	l.SetGlobal("arg", arg)

	hydra := l.SetFuncs(l.NewTable(), hydraFuncs)
	l.SetField(hydra, "BS", lua.LNumber(10.0))
	l.SetField(hydra, "serialize_ver", lua.LNumber(serializeVer))
	l.SetField(hydra, "proto_ver", lua.LNumber(protoVer))
	l.SetGlobal("hydra", hydra)

	l.SetField(l.NewTypeMetatable("hydra.client"), "__index", l.NewFunction(l_client_index))
	l.SetField(l.NewTypeMetatable("hydra.map"), "__index", l.SetFuncs(l.NewTable(), mapFuncs))

	l.SetField(l.NewTypeMetatable("hydra.comp.auth"), "__index", l.SetFuncs(l.NewTable(), compAuthFuncs))
	l.SetField(l.NewTypeMetatable("hydra.comp.map"), "__index", l.SetFuncs(l.NewTable(), compMapFuncs))
	l.SetField(l.NewTypeMetatable("hydra.comp.pkts"), "__index", l.SetFuncs(l.NewTable(), compPktsFuncs))

	for _, str := range builtinFiles {
		if err := l.DoString(str); err != nil {
			panic(err)
		}
	}

	if err := l.DoFile(os.Args[1]); err != nil {
		panic(err)
	}
}
