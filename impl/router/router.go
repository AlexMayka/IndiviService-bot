package router

import (
	"telegram-sdk/core"
)

type regex struct {
	pattern string
	method  core.Handler
}

type msgEquals struct {
	text    string
	handler core.Handler
}

type router struct {
	prefix string
	parent core.Router
	groups []core.Router
	mw     []core.Middleware

	cmd      map[string]core.Handler
	callback map[string]core.Handler
	regex    []regex
	msg      []msgEquals
	any      core.Handler

	state string
}

func New() core.Router {
	return &router{
		prefix: "",
		parent: nil,
		groups: make([]core.Router, 0),
		mw:     make([]core.Middleware, 0),

		cmd:      make(map[string]core.Handler),
		callback: make(map[string]core.Handler),
		regex:    make([]regex, 0),
		msg:      make([]msgEquals, 0),
		any:      nil,

		state: "",
	}
}

func (r *router) Group(prefix string) core.Router {
	childMv := make([]core.Router, len(r.groups))
	copy(childMv, r.groups)

	child := &router{
		prefix: prefix,
		parent: r,
		groups: childMv,
		mw:     r.mw,

		cmd:      make(map[string]core.Handler),
		callback: make(map[string]core.Handler),
		regex:    make([]regex, 0),
		msg:      make([]msgEquals, 0),
		any:      nil,

		state: r.state,
	}
	r.groups = append(r.groups, child)
	return child
}

func (r *router) Use(mw ...core.Middleware) {
	r.mw = append(r.mw, mw...)
}

func (r *router) SetState(state string) {
	r.state = state
}

func (r *router) Command(cmd string, handler core.Handler) {
	r.cmd[cmd] = handler
}

func (r *router) Callback(callback string, handler core.Handler) {
	r.callback[callback] = handler
}

func (r *router) Regex(pattern string, handler core.Handler) {
	r.regex = append(r.regex, regex{pattern: pattern, method: handler})
}

func (r *router) Msg(msg string, handler core.Handler) {
	r.msg = append(r.msg, msgEquals{text: msg, handler: handler})
}

func (r *router) Any(handler core.Handler) {
	r.any = handler
}
