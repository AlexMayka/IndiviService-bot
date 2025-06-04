// Package router provides a hierarchical and middleware-enabled routing system
// for processing incoming Telegram updates based on commands, callbacks, patterns, and text messages.
package router

import (
	"telegram-sdk/core"
)

// router implements the core.Router interface.
// It allows grouping handlers by path prefix and state, and supports middleware chaining.
type router struct {
	prefix string            // Optional prefix used in this routing group
	parent core.Router       // Parent router for hierarchical organization
	groups []core.Router     // Child groups under this router
	mw     []core.Middleware // Middlewares applied to this group

	cmd      map[string]core.Handler // Handlers for exact command matches
	callback map[string]core.Handler // Handlers for callback data matches
	regex    []core.Regex            // Handlers for regex-matching incoming messages
	msg      []core.MsgEquals        // Handlers for messages with exact matching text
	any      core.Handler            // Handler for any unmatched message

	state string // Optional FSM state associated with this router group
}

// New creates and returns a new root router instance.
func New() core.Router {
	return &router{
		prefix: "",
		parent: nil,
		groups: make([]core.Router, 0),
		mw:     make([]core.Middleware, 0),

		cmd:      make(map[string]core.Handler),
		callback: make(map[string]core.Handler),
		regex:    make([]core.Regex, 0),
		msg:      make([]core.MsgEquals, 0),
		any:      nil,

		state: "",
	}
}

// Group creates a new sub-router with a specific prefix under the current router.
// Middlewares and FSM state from the parent are inherited by default.
func (r *router) Group(prefix string) core.Router {
	childMv := make([]core.Router, len(r.groups))
	copy(childMv, r.groups)

	child := &router{
		prefix: prefix,
		parent: r,
		groups: childMv,
		mw:     make([]core.Middleware, 0),

		cmd:      make(map[string]core.Handler),
		callback: make(map[string]core.Handler),
		regex:    make([]core.Regex, 0),
		msg:      make([]core.MsgEquals, 0),
		any:      nil,

		state: r.state,
	}
	r.groups = append(r.groups, child)
	return child
}

// Use registers one or more middlewares for the current router group.
func (r *router) Use(mw ...core.Middleware) core.Router {
	r.mw = append(r.mw, mw...)
	return r
}

// SetState sets a required FSM state for all handlers under this router group.
func (r *router) SetState(state string) core.Router {
	r.state = state
	return r
}

// Command registers a handler for a specific Telegram command (e.g., "/start").
func (r *router) Command(cmd string, handler core.Handler) {
	r.cmd[cmd] = handler
}

// Callback registers a handler for a specific callback data string.
func (r *router) Callback(callback string, handler core.Handler) {
	r.callback[callback] = handler
}

// Regex registers a handler triggered when the message matches a regex pattern.
func (r *router) Regex(pattern string, handler core.Handler) {
	r.regex = append(r.regex, core.Regex{Pattern: pattern, Method: handler})
}

// Msg registers a handler for an exact message text.
func (r *router) Msg(msg string, handler core.Handler) {
	r.msg = append(r.msg, core.MsgEquals{Text: msg, Handler: handler})
}

// Any registers a fallback handler triggered if no other matchers are hit.
func (r *router) Any(handler core.Handler) {
	r.any = handler
}

// GetGroups returns the child groups of this router.
func (r *router) GetGroups() []core.Router {
	return r.groups
}

// GetParent returns the parent of this router group.
func (r *router) GetParent() core.Router {
	return r.parent
}

// GetMw returns the middleware chain associated with this router.
func (r *router) GetMw() []core.Middleware {
	return r.mw
}

// GetCmd returns the map of command handlers.
func (r *router) GetCmd() map[string]core.Handler {
	return r.cmd
}

// GetCallback returns the map of callback query handlers.
func (r *router) GetCallback() map[string]core.Handler {
	return r.callback
}

// GetRegex returns the slice of regex-based message handlers.
func (r *router) GetRegex() []core.Regex {
	return r.regex
}

// GetMsg returns the slice of exact message match handlers.
func (r *router) GetMsg() []core.MsgEquals {
	return r.msg
}

// GetAny returns the fallback handler.
func (r *router) GetAny() core.Handler {
	return r.any
}

// GetState returns the required FSM state associated with this router group.
func (r *router) GetState() string {
	return r.state
}
