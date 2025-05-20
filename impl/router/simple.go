package router

import "telegram-sdk/core"

type router struct {
	cmd map[string]core.Handler
	mw  []core.Middleware
	ms  core.Handler
}

func New() *router {
	return &router{
		cmd: make(map[string]core.Handler),
		mw:  []core.Middleware{},
		ms:  nil,
	}
}

func (s *router) Use(mw ...core.Middleware) {
	s.mw = append(s.mw, mw...)
}

func (s *router) Command(cmd string, h core.Handler) {
	s.cmd[cmd] = h
}

func (s *router) OnMsg(h core.Handler) {
	s.ms = h
}

func (s *router) Match(cmd string) (core.Handler, bool) {
	h, ok := s.cmd[cmd]
	if !ok && s.ms != nil {
		return s.ms, true
	}

	if !ok {
		return nil, false
	}

	for i := len(s.mw) - 1; i >= 0; i-- {
		h = s.mw[i](h)
	}

	return h, ok
}
