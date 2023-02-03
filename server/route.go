package server

import (
	"github.com/vela-ssoc/vela-kit/audit"
	"github.com/vela-ssoc/vela-dns/internal"
	"github.com/vela-ssoc/vela-kit/grep"
	"github.com/vela-ssoc/vela-kit/lua"
	"github.com/vela-ssoc/vela-kit/pipe"
	"sort"
)

type route struct {
	weight int
	name   string
	match  func(string) bool
	handle func(*internal.Context)
}

func (r *route) compile() {
	ch := r.name[0]
	switch ch {
	case '=':
		r.weight = 0
		r.match = func(vv string) bool {
			return vv == r.name[1:]
		}

	default:
		r.weight = len(r.name)

		gx, e := grep.Compile(r.name, nil)
		if e != nil {
			r.match = func(vv string) bool {
				return vv == r.name
			}
			return
		}

		r.match = func(vv string) bool {
			return gx.Match(vv)
		}
	}

}

func (s *server) Len() int {
	return len(s.cfg.route)
}

func (s *server) Less(i, j int) bool {
	w1 := s.cfg.route[i].weight
	w2 := s.cfg.route[j].weight

	if w1 == 0 {
		return true
	}

	if w2 == 0 {
		return true
	}

	return w1 > w2
}

func (s *server) Swap(i, j int) {
	s.cfg.route[j], s.cfg.route[i] = s.cfg.route[i], s.cfg.route[j]
}

func (s *server) addRoute(_ *lua.LState, name string, val lua.LValue) {
	s.cfg.mutex.Lock()
	defer s.cfg.mutex.Unlock()

	if len(name) <= 0 {
		return
	}

	r := route{name: name}
	switch val.Type() {
	case lua.LTString:
		v := val.String()
		r.handle = func(ctx *internal.Context) {
			ctx.Say(v)
		}

	default:
		pp := pipe.New()
		pp.LValue(val)

		if pp.Len() == 0 {
			return
		}

		r.handle = func(ctx *internal.Context) {
			pp.Do(ctx, s.cfg.co, func(err error) {
				audit.Errorf("%s handle call fail %v", s.Name(), err).From(s.CodeVM()).High().Put()
			})
		}
	}

	r.compile()
	s.cfg.route = append(s.cfg.route, r)
	sort.Sort(s)
}
