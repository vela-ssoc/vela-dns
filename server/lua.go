package server

import (
	"github.com/vela-ssoc/vela-kit/vela"
	"github.com/vela-ssoc/vela-kit/lua"
)

var xEnv vela.Environment

/*
	local s = rock.dns.server{
		name = "dnslog",
		bind = "tcp://127.0.0.1:53?timeout=1",
		region = region.sdk(),
	}

	s["=www.baidu.com."] = "127.0.0.1"
	s["www.*"]          = _(ctx) end

	s["*.abb.com.cn"] = _(ctx) end

	s["=www.x2.com."]    = _(ctx) ctx.say("1.1.1.1") end
	s["*.www"] = _(ctx) end
	s.to(kfk)

	s.on_request(_(ctx)

	end)

	s.on_reply(_(ctx)

	end)
	s.start()
*/

func constructor(L *lua.LState) int {
	cfg := newConfig(L)

	proc := L.NewVelaData(cfg.Name, dnsServerTypeOf)
	if proc.IsNil() {
		proc.Set(newDnsServer(cfg))
	} else {
		s := proc.Data.(*server)
		xEnv.Free(s.cfg.co)
		s.cfg = cfg
	}

	L.Push(proc)
	return 1
}

func LuaInjectApi(env vela.Environment, uv lua.UserKV) {
	uv.Set("server", lua.NewFunction(constructor))

	xEnv = env
}
