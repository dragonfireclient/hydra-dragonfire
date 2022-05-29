package main

import (
	"github.com/HimbeerserverDE/srp"
	"github.com/anon55555/mt"
	"github.com/dragonfireclient/hydra-dragonfire/convert"
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

type Auth struct {
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

var authFuncs = map[string]lua.LGFunction{
	"username": l_auth_username,
	"password": l_auth_password,
	"language": l_auth_language,
	"version":  l_auth_version,
	"state":    l_auth_state,
}

func getAuth(l *lua.LState) *Auth {
	return l.CheckUserData(1).Value.(*Auth)
}

func (auth *Auth) create(client *Client, l *lua.LState) {
	if client.state != csNew {
		panic("can't add auth component after connect")
	}

	auth.client = client
	auth.language = "en_US"
	auth.version = "hydra-dragonfire"
	auth.state = asInit
	auth.userdata = l.NewUserData()
	auth.userdata.Value = auth
	l.SetMetatable(auth.userdata, l.GetTypeMetatable("hydra.auth"))
}

func (auth *Auth) push() lua.LValue {
	return auth.userdata
}

func (auth *Auth) connect() {
	if auth.username == "" {
		panic("missing username")
	}

	go func() {
		for auth.client.state == csConnected && auth.state == asInit {
			auth.client.conn.SendCmd(&mt.ToSrvInit{
				SerializeVer: serializeVer,
				MinProtoVer:  protoVer,
				MaxProtoVer:  protoVer,
				PlayerName:   auth.username,
			})
			time.Sleep(500 * time.Millisecond)
		}
	}()
}

func (auth *Auth) fail(err string) {
	auth.err = err
	auth.state = asError
	auth.client.disconnect()
}

func (auth *Auth) checkState(state authState, pkt *mt.Pkt) bool {
	if auth.state == state {
		return true
	}

	auth.fail("received " + string(convert.PushPktType(pkt)) + " in invalid state")
	return false
}

func (auth *Auth) process(pkt *mt.Pkt) {
	if auth.state == asError {
		return
	}

	switch cmd := pkt.Cmd.(type) {
	case *mt.ToCltHello:
		if !auth.checkState(asInit, pkt) {
			return
		}

		if cmd.SerializeVer != 28 {
			auth.fail("unsupported serialize version")
			return
		}

		if cmd.AuthMethods == mt.FirstSRP {
			salt, verifier, err := srp.NewClient([]byte(strings.ToLower(auth.username)), []byte(auth.password))
			if err != nil {
				auth.fail(err.Error())
				return
			}

			auth.client.conn.SendCmd(&mt.ToSrvFirstSRP{
				Salt:        salt,
				Verifier:    verifier,
				EmptyPasswd: auth.password == "",
			})
			auth.state = asVerified
		} else if cmd.AuthMethods == mt.SRP {
			var err error
			auth.srpBytesA, auth.bytesA, err = srp.InitiateHandshake()
			if err != nil {
				auth.fail(err.Error())
				return
			}

			auth.client.conn.SendCmd(&mt.ToSrvSRPBytesA{
				A:      auth.srpBytesA,
				NoSHA1: true,
			})
			auth.state = asRequested
		} else {
			auth.fail("invalid auth methods")
			return
		}

	case *mt.ToCltSRPBytesSaltB:
		if !auth.checkState(asRequested, pkt) {
			return
		}

		srpBytesK, err := srp.CompleteHandshake(auth.srpBytesA, auth.bytesA, []byte(strings.ToLower(auth.username)), []byte(auth.password), cmd.Salt, cmd.B)
		if err != nil {
			auth.fail(err.Error())
			return
		}

		M := srp.ClientProof([]byte(auth.username), cmd.Salt, auth.srpBytesA, cmd.B, srpBytesK)
		auth.srpBytesA = []byte{}
		auth.bytesA = []byte{}

		if M == nil {
			auth.fail("srp safety check fail")
			return
		}

		auth.client.conn.SendCmd(&mt.ToSrvSRPBytesM{
			M: M,
		})
		auth.state = asVerified

	case *mt.ToCltAcceptAuth:
		auth.client.conn.SendCmd(&mt.ToSrvInit2{Lang: auth.language})

	case *mt.ToCltTimeOfDay:
		if auth.state == asActive {
			return
		}

		if !auth.checkState(asVerified, pkt) {
			return
		}

		auth.client.conn.SendCmd(&mt.ToSrvCltReady{
			Major:    5,
			Minor:    6,
			Patch:    0,
			Reserved: 0,
			Formspec: 4,
			Version:  auth.version,
		})
		auth.state = asActive
	}
}

func (auth *Auth) accessProperty(l *lua.LState, key string, ptr *string) int {
	if str, ok := l.Get(2).(lua.LString); ok {
		if auth.client.state != csNew {
			panic("can't change " + key + " after connecting")
		}
		*ptr = string(str)
		return 0
	} else {
		l.Push(lua.LString(*ptr))
		return 1
	}
}

func l_auth_username(l *lua.LState) int {
	auth := getAuth(l)
	return auth.accessProperty(l, "username", &auth.username)
}

func l_auth_password(l *lua.LState) int {
	auth := getAuth(l)
	return auth.accessProperty(l, "password", &auth.password)
}

func l_auth_language(l *lua.LState) int {
	auth := getAuth(l)
	return auth.accessProperty(l, "language", &auth.language)
}

func l_auth_version(l *lua.LState) int {
	auth := getAuth(l)
	return auth.accessProperty(l, "version", &auth.version)
}

func l_auth_state(l *lua.LState) int {
	auth := getAuth(l)

	switch auth.state {
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
		l.Push(lua.LString(auth.err))
		return 2
	}

	return 1
}
