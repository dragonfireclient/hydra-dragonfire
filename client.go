package main

import (
	"errors"
	"github.com/dragonfireclient/hydra-dragonfire/convert"
	"github.com/dragonfireclient/mt"
	"github.com/yuin/gopher-lua"
	"net"
	"sync"
)

type clientState uint8

const (
	csNew clientState = iota
	csConnected
	csDisconnected
)

type Component interface {
	create(client *Client, l *lua.LState)
	push() lua.LValue
	connect()
	process(pkt *mt.Pkt)
}

type Client struct {
	address    string
	state      clientState
	conn       mt.Peer
	muConn     sync.Mutex
	queue      chan Event
	components map[string]Component
	muComp     sync.Mutex
	table      *lua.LTable
	userdata   *lua.LUserData
}

var clientFuncs = map[string]lua.LGFunction{
	"address": l_client_address,
	"state":   l_client_state,
	"connect": l_client_connect,
	"poll":    l_client_poll,
	"close":   l_client_close,
	"enable":  l_client_enable,
	"send":    l_client_send,
}

type EventError struct {
	err string
}

func (evt EventError) handle(l *lua.LState, val lua.LValue) {
	l.SetField(val, "type", lua.LString("error"))
	l.SetField(val, "error", lua.LString(evt.err))
}

type EventDisconnect struct {
	client *Client
}

func (evt EventDisconnect) handle(l *lua.LState, val lua.LValue) {
	l.SetField(val, "type", lua.LString("disconnect"))
	evt.client.state = csDisconnected
}

func getClient(l *lua.LState) *Client {
	return l.CheckUserData(1).Value.(*Client)
}

func getClients(l *lua.LState) []*Client {
	tbl := l.CheckTable(1)
	n := tbl.MaxN()

	clients := make([]*Client, 0, n)
	for i := 1; i <= n; i++ {
		clients = append(clients, l.RawGetInt(tbl, i).(*lua.LUserData).Value.(*Client))
	}

	return clients
}

func (client *Client) closeConn() {
	client.muConn.Lock()
	defer client.muConn.Unlock()

	if client.state == csConnected {
		client.conn.Close()
	}
}

func l_client(l *lua.LState) int {
	client := &Client{}

	client.address = l.CheckString(1)
	client.state = csNew
	client.components = map[string]Component{}
	client.table = l.NewTable()
	client.userdata = l.NewUserData()
	client.userdata.Value = client
	l.SetMetatable(client.userdata, l.GetTypeMetatable("hydra.client"))

	l.Push(client.userdata)
	return 1
}

func l_client_index(l *lua.LState) int {
	client := getClient(l)
	key := l.CheckString(2)

	if key == "data" {
		l.Push(client.table)
	} else if fun, exists := clientFuncs[key]; exists {
		l.Push(l.NewFunction(fun))
	} else if component, exists := client.components[key]; exists {
		l.Push(component.push())
	} else {
		l.Push(lua.LNil)
	}

	return 1
}

func l_client_address(l *lua.LState) int {
	client := getClient(l)
	l.Push(lua.LString(client.address))
	return 1
}

func l_client_state(l *lua.LState) int {
	client := getClient(l)
	switch client.state {
	case csNew:
		l.Push(lua.LString("new"))
	case csConnected:
		l.Push(lua.LString("connected"))
	case csDisconnected:
		l.Push(lua.LString("disconnected"))
	}
	return 1
}

func l_client_connect(l *lua.LState) int {
	client := getClient(l)

	if client.state != csNew {
		panic("can't reconnect")
	}

	addr, err := net.ResolveUDPAddr("udp", client.address)
	if err != nil {
		panic(err)
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		panic(err)
	}

	client.state = csConnected
	client.conn = mt.Connect(conn)
	client.queue = make(chan Event, 1024)

	go func() {
		for {
			pkt, err := client.conn.Recv()

			if err == nil {
				client.muComp.Lock()
				for _, comp := range client.components {
					comp.process(&pkt)
				}
				client.muComp.Unlock()
			} else if errors.Is(err, net.ErrClosed) {
				client.queue <- EventDisconnect{client: client}
				return
			} else {
				client.queue <- EventError{err: err.Error()}
			}
		}
	}()

	client.muComp.Lock()
	for _, comp := range client.components {
		comp.connect()
	}
	client.muComp.Unlock()

	return 0
}

func l_client_poll(l *lua.LState) int {
	client := getClient(l)
	return doPoll(l, []*Client{client})
}

func l_client_close(l *lua.LState) int {
	client := getClient(l)
	client.closeConn()
	return 0
}

func l_client_enable(l *lua.LState) int {
	client := getClient(l)
	n := l.GetTop()

	client.muComp.Lock()
	defer client.muComp.Unlock()

	for i := 2; i <= n; i++ {
		name := l.CheckString(i)

		if comp, exists := client.components[name]; !exists {
			switch name {
			case "auth":
				comp = &CompAuth{}
			case "map":
				comp = &CompMap{}
			case "pkts":
				comp = &CompPkts{}
			default:
				panic("invalid component: " + name)
			}

			client.components[name] = comp
			comp.create(client, l)
		}
	}

	return 0
}

func l_client_send(l *lua.LState) int {
	client := getClient(l)

	client.muConn.Lock()
	defer client.muConn.Unlock()

	if client.state != csConnected {
		panic("not connected")
	}

	cmd := convert.ReadCmd(l)
	doAck := l.ToBool(4)

	if client.state == csConnected {
		ack, err := client.conn.SendCmd(cmd)
		if err != nil && !errors.Is(err, net.ErrClosed) {
			panic(err)
		}

		if doAck && !cmd.DefaultPktInfo().Unrel {
			<-ack
		}
	}

	return 0
}
