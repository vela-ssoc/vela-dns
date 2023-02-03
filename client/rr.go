package client

import (
	"fmt"
	"github.com/miekg/dns"
	"github.com/vela-ssoc/vela-kit/lua"
)

type RR struct {
	data dns.RR
}

func (rr *RR) String() string                         { return fmt.Sprintf("dns.RR %p", rr) }
func (rr *RR) Type() lua.LValueType                   { return lua.LTObject }
func (rr *RR) AssertFloat64() (float64, bool)         { return 0, false }
func (rr *RR) AssertString() (string, bool)           { return "", false }
func (rr *RR) AssertFunction() (*lua.LFunction, bool) { return nil, false }
func (rr *RR) Peek() lua.LValue                       { return rr }

func (rr *RR) Index(L *lua.LState, key string) lua.LValue {
	switch key {
	case "name":
		return lua.S2L(rr.data.Header().Name)
	case "type":
		return lua.S2L(dns.Type(rr.data.Header().Rrtype).String())
	case "ttl":
		return lua.LNumber(rr.data.Header().Ttl)
	case "raw":
		return lua.S2L(rr.data.String())
	case "header":
		return lua.S2L(rr.data.Header().String())

	case "value":
		switch record := rr.data.(type) {
		case *dns.A:
			return lua.S2L(record.A.String())
		case *dns.CNAME:
			return lua.S2L(record.Target)
		case *dns.AAAA:
			return lua.S2L(record.AAAA.String())
		case *dns.DNAME:
			return lua.S2L(record.Target)
		default:
			return lua.S2L(rr.data.String())
		}
	}

	return lua.LNil
}
