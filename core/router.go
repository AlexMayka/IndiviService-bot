package core

type Handler func(ctx Context)
type Middleware func(h Handler) Handler

type Router interface {
	Use(mw ...Middleware)
	Command(cmd string, h Handler)
	OnMsg(h Handler)
	Match(cmd string) (Handler, bool)
}
