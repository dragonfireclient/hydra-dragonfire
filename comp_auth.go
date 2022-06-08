package main

import (
	"github.com/HimbeerserverDE/srp"
	"github.com/dragonfireclient/hydra-dragonfire/convert"
	"github.com/dragonfireclient/mt"
	"github.com/yuin/gopher-lua"
	"strings"
	"time"
)

type authState uint8

const (
	asInit authState = iota
	asRequested
	asVerified
	asActive
	asError
)

type CompAuth struct {
	client            *Client
	username          string
	password          string
	language          string
	version           string
	state             authState
	err               string
	srpBytesA, bytesA []byte
	userdata          *lua.LUserData
}

var compAuthFuncs = map[string]lua.LGFunction{
	"username": l_comp_auth_username,
	"password": l_comp_auth_password,
	"language": l_comp_auth_language,
	"version":  l_comp_auth_version,
	"state":    l_comp_auth_state,
}

func getCompAuth(l *lua.LState) *CompAuth {
	return l.CheckUserData(1).Value.(*CompAuth)
}

func (comp *CompAuth) create(client *Client, l *lua.LState) {
	if client.state != csNew {
		panic("can't add auth component after connect")
	}

	comp.client = client
	comp.language = "en_US"
	comp.version = "hydra-dragonfire"
	comp.state = asInit
	comp.userdata = l.NewUserData()
	comp.userdata.Value = comp
	l.SetMetatable(comp.userdata, l.GetTypeMetatable("hydra.comp.auth"))
}

func (comp *CompAuth) push() lua.LValue {
	return comp.userdata
}

func (comp *CompAuth) connect() {
	if comp.username == "" {
		panic("missing username")
	}

	go func() {
		for comp.client.state == csConnected && comp.state == asInit {
			comp.client.conn.SendCmd(&mt.ToSrvInit{
				SerializeVer: serializeVer,
				MinProtoVer:  protoVer,
				MaxProtoVer:  protoVer,
				PlayerName:   comp.username,
			})
			time.Sleep(500 * time.Millisecond)
		}
	}()
}

func (comp *CompAuth) fail(err string) {
	comp.err = err
	comp.state = asError
	comp.client.closeConn()
}

func (comp *CompAuth) checkState(state authState, pkt *mt.Pkt) bool {
	if comp.state == state {
		return true
	}

	comp.fail("received " + string(convert.PushPktType(pkt)) + " in invalid state")
	return false
}

func (comp *CompAuth) process(pkt *mt.Pkt) {
	if comp.state == asError {
		return
	}

	switch cmd := pkt.Cmd.(type) {
	case *mt.ToCltHello:
		if !comp.checkState(asInit, pkt) {
			return
		}

		if cmd.SerializeVer != serializeVer {
			comp.fail("unsupported serialize version")
			return
		}

		if cmd.AuthMethods == mt.FirstSRP {
			salt, verifier, err := srp.NewClient([]byte(strings.ToLower(comp.username)), []byte(comp.password))
			if err != nil {
				comp.fail(err.Error())
				return
			}

			comp.client.conn.SendCmd(&mt.ToSrvFirstSRP{
				Salt:        salt,
				Verifier:    verifier,
				EmptyPasswd: comp.password == "",
			})
			comp.state = asVerified
		} else if cmd.AuthMethods == mt.SRP {
			var err error
			comp.srpBytesA, comp.bytesA, err = srp.InitiateHandshake()
			if err != nil {
				comp.fail(err.Error())
				return
			}

			comp.client.conn.SendCmd(&mt.ToSrvSRPBytesA{
				A:      comp.srpBytesA,
				NoSHA1: true,
			})
			comp.state = asRequested
		} else {
			comp.fail("invalid auth methods")
			return
		}

	case *mt.ToCltSRPBytesSaltB:
		if !comp.checkState(asRequested, pkt) {
			return
		}

		srpBytesK, err := srp.CompleteHandshake(comp.srpBytesA, comp.bytesA, []byte(strings.ToLower(comp.username)), []byte(comp.password), cmd.Salt, cmd.B)
		if err != nil {
			comp.fail(err.Error())
			return
		}

		M := srp.ClientProof([]byte(comp.username), cmd.Salt, comp.srpBytesA, cmd.B, srpBytesK)
		comp.srpBytesA = []byte{}
		comp.bytesA = []byte{}

		if M == nil {
			comp.fail("srp safety check fail")
			return
		}

		comp.client.conn.SendCmd(&mt.ToSrvSRPBytesM{
			M: M,
		})
		comp.state = asVerified

	case *mt.ToCltAcceptAuth:
		comp.client.conn.SendCmd(&mt.ToSrvInit2{Lang: comp.language})

	case *mt.ToCltTimeOfDay:
		if comp.state == asActive {
			return
		}

		if !comp.checkState(asVerified, pkt) {
			return
		}

		comp.client.conn.SendCmd(&mt.ToSrvCltReady{
			Major:    5,
			Minor:    6,
			Patch:    0,
			Reserved: 0,
			Formspec: 4,
			Version:  comp.version,
		})
		comp.state = asActive
	}
}

func (comp *CompAuth) accessProperty(l *lua.LState, key string, ptr *string) int {
	if str, ok := l.Get(2).(lua.LString); ok {
		if comp.client.state != csNew {
			panic("can't change " + key + " after connecting")
		}
		*ptr = string(str)
		return 0
	} else {
		l.Push(lua.LString(*ptr))
		return 1
	}
}

func l_comp_auth_username(l *lua.LState) int {
	comp := getCompAuth(l)
	return comp.accessProperty(l, "username", &comp.username)
}

func l_comp_auth_password(l *lua.LState) int {
	comp := getCompAuth(l)
	return comp.accessProperty(l, "password", &comp.password)
}

func l_comp_auth_language(l *lua.LState) int {
	comp := getCompAuth(l)
	return comp.accessProperty(l, "language", &comp.language)
}

func l_comp_auth_version(l *lua.LState) int {
	comp := getCompAuth(l)
	return comp.accessProperty(l, "version", &comp.version)
}

func l_comp_auth_state(l *lua.LState) int {
	comp := getCompAuth(l)

	switch comp.state {
	case asInit:
		l.Push(lua.LString("init"))
	case asRequested:
		l.Push(lua.LString("requested"))
	case asVerified:
		l.Push(lua.LString("verified"))
	case asActive:
		l.Push(lua.LString("active"))
	case asError:
		l.Push(lua.LString("error"))
		l.Push(lua.LString(comp.err))
		return 2
	}

	return 1
}
