package main

import (
	_ "embed"
	"github.com/Shopify/go-lua"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var lastTime = time.Now()
var canceled = false

//go:embed builtin/vector.lua
var vectorLibrary string

func l_dtime(l *lua.State) int {
	l.PushNumber(time.Since(lastTime).Seconds())
	lastTime = time.Now()
	return 1
}

func l_canceled(l *lua.State) int {
	l.PushBoolean(canceled)
	return 1
}

func signalChannel() chan os.Signal {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
	return sig
}

func main() {
	if len(os.Args) < 2 {
		panic("missing filename")
	}

	go func() {
		<-signalChannel()
		canceled = true
	}()

	l := lua.NewState()
	lua.OpenLibraries(l)

	lua.NewLibrary(l, []lua.RegistryFunction{
		{Name: "client", Function: l_client},
		{Name: "dtime", Function: l_dtime},
		{Name: "canceled", Function: l_canceled},
		{Name: "poll", Function: l_poll},
	})

	l.PushNumber(10.0)
	l.SetField(-2, "BS")

	l.SetGlobal("hydra")

	l.NewTable()
	for i, arg := range os.Args {
		l.PushString(arg)
		l.RawSetInt(-2, i - 1)
	}
	l.SetGlobal("arg")

	if err := lua.DoString(l, vectorLibrary); err != nil {
		panic(err)
	}

	if err := lua.DoFile(l, os.Args[1]); err != nil {
		panic(err)
	}
}
