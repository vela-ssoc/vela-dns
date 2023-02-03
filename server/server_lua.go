package server

import (
	"github.com/vela-ssoc/vela-kit/lua"
	"github.com/vela-ssoc/vela-kit/pipe"
)

func (s *server) pipeL(L *lua.LState) int {
	s.cfg.pipe.CheckMany(L, pipe.Seek(0))
	return 0
}

func (s *server) toL(L *lua.LState) int {
	s.cfg.sdk.CheckMany(L, pipe.Seek(0))
	return 0
}

func (s *server) startL(L *lua.LState) int {
	xEnv.Start(L, s).From(s.CodeVM()).Do()
	return 0
}

func (s *server) Index(L *lua.LState, key string) lua.LValue {
	switch key {

	case "start":
		return L.NewFunction(s.startL)

	case "pipe":
		return L.NewFunction(s.pipeL)
	case "to":
		return L.NewFunction(s.toL)
	}
	return lua.LNil
}

func (s *server) NewIndex(L *lua.LState, key string, val lua.LValue) {
	s.addRoute(L, key, val)
}
