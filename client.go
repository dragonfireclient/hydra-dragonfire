package main

import (
	"errors"
	"github.com/anon55555/mt"
	"github.com/dragonfireclient/hydra-dragonfire/fromlua"
	"github.com/dragonfireclient/hydra-dragonfire/tolua"
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
	tolua() lua.LValue
	connect()
	process(pkt *mt.Pkt)
}

type Client struct {
	mu         sync.Mutex
	address    string
	state      clientState
	conn       mt.Peer
	queue      chan *mt.Pkt
	wildcard   bool
	subscribed map[string]struct{}
	components map[string]Component
	userdata   *lua.LUserData
}

var clientFuncs = map[string]lua.LGFunction{
	"address":     l_client_address,
	"state":       l_client_state,
	"connect":     l_client_connect,
	"poll":        l_client_poll,
	"disconnect":  l_client_disconnect,
	"enable":      l_client_enable,
	"subscribe":   l_client_subscribe,
	"unsubscribe": l_client_unsubscribe,
	"wildcard":    l_client_wildcard,
	"send":        l_client_send,
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

func getStrings(l *lua.LState) []string {
	n := l.GetTop()

	strs := make([]string, 0, n-1)
	for i := 2; i <= n; i++ {
		strs = append(strs, l.CheckString(i))
	}

	return strs
}

func (client *Client) disconnect() {
	client.mu.Lock()
	defer client.mu.Unlock()

	if client.state == csConnected {
		client.conn.Close()
	}
}

func l_client(l *lua.LState) int {
	client := &Client{}

	client.address = l.CheckString(1)
	client.state = csNew
	client.wildcard = false
	client.subscribed = map[string]struct{}{}
	client.components = map[string]Component{}
	client.userdata = l.NewUserData()
	client.userdata.Value = client
	l.SetMetatable(client.userdata, l.GetTypeMetatable("hydra.client"))

	l.Push(client.userdata)
	return 1
}

func l_client_index(l *lua.LState) int {
	client := getClient(l)
	key := l.CheckString(2)

	if fun, exists := clientFuncs[key]; exists {
		l.Push(l.NewFunction(fun))
	} else if component, exists := client.components[key]; exists {
		l.Push(component.tolua())
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
	client.queue = make(chan *mt.Pkt, 1024)

	go func() {
		for {
			pkt, err := client.conn.Recv()

			if err == nil {
				client.mu.Lock()

				for _, component := range client.components {
					component.process(&pkt)
				}

				if _, exists := client.subscribed[string(tolua.PktType(&pkt))]; exists || client.wildcard {
					client.queue <- &pkt
				}

				client.mu.Unlock()
			} else if errors.Is(err, net.ErrClosed) {
				close(client.queue)
				return
			}
		}
	}()

	client.mu.Lock()
	for _, component := range client.components {
		component.connect()
	}
	client.mu.Unlock()

	return 0
}

func l_client_poll(l *lua.LState) int {
	client := getClient(l)
	_, pkt, timeout := doPoll(l, []*Client{client})

	l.Push(tolua.Pkt(l, pkt))
	l.Push(lua.LBool(timeout))
	return 2
}

func l_client_disconnect(l *lua.LState) int {
	client := getClient(l)
	client.disconnect()
	return 0
}

func l_client_enable(l *lua.LState) int {
	client := getClient(l)
	client.mu.Lock()
	defer client.mu.Unlock()

	for _, compname := range getStrings(l) {
		if component, exists := client.components[compname]; !exists {
			switch compname {
			case "auth":
				component = &Auth{}
			default:
				panic("invalid component: " + compname)
			}

			client.components[compname] = component
			component.create(client, l)
		}
	}

	return 0
}

func l_client_subscribe(l *lua.LState) int {
	client := getClient(l)
	client.mu.Lock()
	defer client.mu.Unlock()

	for _, pkt := range getStrings(l) {
		client.subscribed[pkt] = struct{}{}
	}

	return 0
}

func l_client_unsubscribe(l *lua.LState) int {
	client := getClient(l)
	client.mu.Lock()
	defer client.mu.Unlock()

	for _, pkt := range getStrings(l) {
		delete(client.subscribed, pkt)
	}

	return 0
}

func l_client_wildcard(l *lua.LState) int {
	client := getClient(l)
	client.wildcard = l.ToBool(2)
	return 0
}

func l_client_send(l *lua.LState) int {
	client := getClient(l)
	cmd := fromlua.Cmd(l)
	doAck := l.ToBool(4)

	client.mu.Lock()
	defer client.mu.Unlock()

	if client.state == csConnected {
		ack, err := client.conn.SendCmd(cmd)
		if err != nil {
			panic(err)
		}

		if doAck && !cmd.DefaultPktInfo().Unrel {
			<-ack
		}
	}

	return 0
}
