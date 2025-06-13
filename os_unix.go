//go:build linux || openbsd || freebsd || netbsd || dragonfly || solaris || darwin || aix

package main

import (
	"os"
	"os/signal"
	"runtime/pprof"
	"syscall"
)

func backtrace_listen() {
	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGUSR1)
		for {
			<-ch
			pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)
		}
	}()
}
