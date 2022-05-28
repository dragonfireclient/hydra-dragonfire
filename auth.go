package main

import (
	"github.com/HimbeerserverDE/srp"
	"github.com/Shopify/go-lua"
	"github.com/anon55555/mt"
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
	state             authState
	err               string
	srpBytesA, bytesA []byte
}

func getAuth(l *lua.State) *Auth {
	return lua.CheckUserData(l, 1, "hydra.auth").(*Auth)
}

func (auth *Auth) create(client *Client) {
	auth.client = client
	auth.language = "en_US"
	auth.state = asInit
}

func (auth *Auth) push(l *lua.State) {
	l.PushUserData(auth)

	if lua.NewMetaTable(l, "hydra.auth") {
		lua.NewLibrary(l, []lua.RegistryFunction{
			{Name: "username", Function: l_auth_username},
			{Name: "password", Function: l_auth_password},
			{Name: "language", Function: l_auth_language},
			{Name: "state", Function: l_auth_state},
		})
		l.SetField(-2, "__index")
	}
	l.SetMetaTable(-2)
}

func (auth *Auth) canConnect() (bool, string) {
	if auth.username == "" {
		return false, "missing username"
	}

	return true, ""
}

func (auth *Auth) connect() {
	go func() {
		for auth.state == asInit && auth.client.state == csConnected {
			auth.client.conn.SendCmd(&mt.ToSrvInit{
				SerializeVer: 28,
				MinProtoVer:  39,
				MaxProtoVer:  39,
				PlayerName:   auth.username,
			})
			time.Sleep(500 * time.Millisecond)
		}
	}()
}

func (auth *Auth) setError(err string) {
	auth.state = asError
	auth.err = err
	auth.client.conn.Close()
}

func (auth *Auth) checkState(state authState, pkt *mt.Pkt) bool {
	if auth.state == state {
		return true
	}

	auth.setError("received " + pktToString(pkt) + " in invalid state")
	return false
}

func (auth *Auth) handle(pkt *mt.Pkt, l *lua.State, idx int) {
	if pkt == nil {
		return
	}

	switch cmd := pkt.Cmd.(type) {
	case *mt.ToCltHello:
		if !auth.checkState(asInit, pkt) {
			return
		}

		if cmd.SerializeVer != 28 {
			auth.setError("unsupported serialize_ver")
			return
		}

		if cmd.AuthMethods == mt.FirstSRP {
			salt, verifier, err := srp.NewClient([]byte(strings.ToLower(auth.username)), []byte(auth.password))
			if err != nil {
				auth.setError(err.Error())
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
				auth.setError(err.Error())
				return
			}

			auth.client.conn.SendCmd(&mt.ToSrvSRPBytesA{
				A:      auth.srpBytesA,
				NoSHA1: true,
			})
			auth.state = asRequested
		} else {
			auth.setError("invalid auth methods")
			return			
		}

	case *mt.ToCltSRPBytesSaltB:
		if !auth.checkState(asRequested, pkt) {
			return
		}

		srpBytesK, err := srp.CompleteHandshake(auth.srpBytesA, auth.bytesA, []byte(strings.ToLower(auth.username)), []byte(auth.password), cmd.Salt, cmd.B)
		if err != nil {
			auth.setError(err.Error())
			return
		}

		M := srp.ClientProof([]byte(auth.username), cmd.Salt, auth.srpBytesA, cmd.B, srpBytesK)
		auth.srpBytesA = []byte{}
		auth.bytesA = []byte{}

		if M == nil {
			auth.setError("srp safety check fail")
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
			// Version:  "hydra-dragonfire",
			Version: "astolfo",
		})
		auth.state = asActive
	}
}

func l_auth_username(l *lua.State) int {
	auth := getAuth(l)

	if l.IsString(2) {
		if auth.client.state > csNew {
			panic("can't change username after connecting")
		}
		auth.username = lua.CheckString(l, 2)
		return 0
	} else {
		l.PushString(auth.username)
		return 1
	}
}

func l_auth_password(l *lua.State) int {
	auth := getAuth(l)

	if l.IsString(2) {
		if auth.client.state > csNew {
			panic("can't change password after connecting")
		}
		auth.password = lua.CheckString(l, 2)
		return 0
	} else {
		l.PushString(auth.password)
		return 1
	}
}

func l_auth_language(l *lua.State) int {
	auth := getAuth(l)

	if l.IsString(2) {
		if auth.client.state > csNew {
			panic("can't change language after connecting")
		}
		auth.language = lua.CheckString(l, 2)
		return 0
	} else {
		l.PushString(auth.language)
		return 1
	}
}

func l_auth_state(l *lua.State) int {
	auth := getAuth(l)

	switch auth.state {
	case asInit:
		l.PushString("init")
	case asRequested:
		l.PushString("requested")
	case asVerified:
		l.PushString("verified")
	case asActive:
		l.PushString("active")
	case asError:
		l.PushString("error")
		l.PushString(auth.err)
		return 2
	}

	return 1
}
