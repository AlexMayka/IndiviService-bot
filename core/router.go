package core

type Handler func(ctx Context)
type Middleware func(h Handler) Handler

type Router interface {
	Group(prefix string) Router

	Use(mw ...Middleware)
	SetState(state string)

	Command(cmd string, h Handler)
	Callback(callback string, h Handler)
	Regex(pattern string, h Handler)
	Msg(msg string, h Handler)
	Any(h Handler)
}
