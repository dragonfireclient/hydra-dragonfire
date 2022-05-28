package main

import (
	"errors"
	"github.com/Shopify/go-lua"
	"github.com/anon55555/mt"
	"net"
)

type clientState uint8

const (
	csNew clientState = iota
	csConnected
	csDisconnected
)

type Handler interface {
	create(client *Client)
	push(l *lua.State)
	canConnect() (bool, string)
	connect()
	handle(pkt *mt.Pkt, l *lua.State, idx int)
}

type Client struct {
	address  string
	state    clientState
	handlers map[string]Handler
	conn     mt.Peer
	queue    chan *mt.Pkt
}

func getClient(l *lua.State) *Client {
	return lua.CheckUserData(l, 1, "hydra.client").(*Client)
}

func l_client(l *lua.State) int {
	client := &Client{
		address:  lua.CheckString(l, 1),
		state:    csNew,
		handlers: map[string]Handler{},
	}

	l.PushUserData(client)

	if lua.NewMetaTable(l, "hydra.client") {
		lua.NewLibrary(l, []lua.RegistryFunction{
			{Name: "address", Function: l_client_address},
			{Name: "state", Function: l_client_state},
			{Name: "handler", Function: l_client_handler},
			{Name: "connect", Function: l_client_connect},
			{Name: "disconnect", Function: l_client_disconnect},
		})
		l.SetField(-2, "__index")
	}
	l.SetMetaTable(-2)

	return 1
}

func l_client_address(l *lua.State) int {
	client := getClient(l)
	l.PushString(client.address)
	return 1
}

func l_client_state(l *lua.State) int {
	client := getClient(l)
	switch client.state {
	case csNew:
		l.PushString("new")
	case csConnected:
		l.PushString("connected")
	case csDisconnected:
		l.PushString("disconnected")
	}
	return 1
}

func l_client_handler(l *lua.State) int {
	client := getClient(l)
	name := lua.CheckString(l, 2)

	handler, exists := client.handlers[name]
	if !exists {
		switch name {
		case "callbacks":
			handler = &Callbacks{}

		case "auth":
			handler = &Auth{}

		default:
			return 0
		}

		client.handlers[name] = handler
		handler.create(client)
	}

	handler.push(l)
	return 1
}

func l_client_connect(l *lua.State) int {
	client := getClient(l)

	if client.state != csNew {
		l.PushBoolean(false)
		l.PushString("invalid state")
		return 2
	}

	for _, handler := range client.handlers {
		ok, err := handler.canConnect()

		if !ok {
			l.PushBoolean(false)
			l.PushString(err)
			return 2
		}
	}

	addr, err := net.ResolveUDPAddr("udp", client.address)
	if err != nil {
		l.PushBoolean(false)
		l.PushString(err.Error())
		return 2
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		l.PushBoolean(false)
		l.PushString(err.Error())
		return 2
	}

	client.state = csConnected
	client.conn = mt.Connect(conn)
	client.queue = make(chan *mt.Pkt, 1024)

	for _, handler := range client.handlers {
		handler.connect()
	}

	go func() {
		for {
			pkt, err := client.conn.Recv()

			if err == nil {
				client.queue <- &pkt
			} else if errors.Is(err, net.ErrClosed) {
				close(client.queue)
				return
			}
		}
	}()

	l.PushBoolean(true)
	return 1
}

func l_client_disconnect(l *lua.State) int {
	client := getClient(l)

	if client.state == csConnected {
		client.conn.Close()
	}

	return 0
}
