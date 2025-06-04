package core

// Handler defines a function that processes an incoming update using Context.
type Handler func(ctx Context)

// Middleware wraps a Handler with additional behavior, such as access control or logging.
type Middleware func(h Handler) Handler

// Router defines an interface for registering and organizing message handlers.
// It supports grouping, middleware chains, and message pattern matching.
type Router interface {
	// Group creates a sub-router with a specific prefix.
	Group(prefix string) Router

	// Use attaches one or more middleware functions to the router.
	Use(mw ...Middleware) Router

	// SetState binds the router to a specific FSM state.
	SetState(state string) Router

	// Command registers a handler for a specific bot command (e.g., "/start").
	Command(cmd string, h Handler)

	// Callback registers a handler for callback queries with specific data.
	Callback(callback string, h Handler)

	// Regex registers a handler triggered by matching a regex pattern.
	Regex(pattern string, h Handler)

	// Msg registers a handler for exact text messages.
	Msg(msg string, h Handler)

	// Any registers a fallback handler for unmatched messages.
	Any(h Handler)

	// GetGroups returns child routers created via Group().
	GetGroups() []Router

	// GetParent returns the parent router in the hierarchy.
	GetParent() Router

	// GetMw returns the middleware chain for this router.
	GetMw() []Middleware

	// GetCmd returns registered command handlers.
	GetCmd() map[string]Handler

	// GetCallback returns registered callback handlers.
	GetCallback() map[string]Handler

	// GetRegex returns registered regex matchers.
	GetRegex() []Regex

	// GetMsg returns exact message handlers.
	GetMsg() []MsgEquals

	// GetAny returns the fallback handler.
	GetAny() Handler

	// GetState returns the state associated with the router.
	GetState() string
}
