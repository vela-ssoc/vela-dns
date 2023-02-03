package vdns

import (
	"github.com/vela-ssoc/vela-kit/vela"
	"github.com/vela-ssoc/vela-dns/client"
	"github.com/vela-ssoc/vela-dns/server"
	"github.com/vela-ssoc/vela-kit/lua"
)

func WithEnv(env vela.Environment) {
	kv := lua.NewUserKV()
	client.WithEnv(env, kv)
	server.LuaInjectApi(env, kv)
	env.Set("dns", kv)
}
