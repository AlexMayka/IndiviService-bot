// Package dispatcher contains logic for compiling router trees into flattened structures
// and wrapping handlers with inherited middleware.
package dispatcher

import (
	"telegram-sdk/core"
)

// collectMiddleware recursively traverses the router tree upwards, collecting
// all middleware applied to the current router and its parent routers.
func collectMiddleware(all []core.Middleware, r core.Router) []core.Middleware {
	all = append(all, r.GetMw()...)
	if parent := r.GetParent(); parent != nil {
		all = collectMiddleware(all, parent)
	}
	return all
}

// wrap wraps the given handler with all middleware collected from the router hierarchy.
// Middleware is applied in reverse order (from parent to current).
func wrap(r core.Router, method core.Handler) core.Handler {
	mv := make([]core.Middleware, 0)
	mv = collectMiddleware(mv, r)

	tmp := method
	for i := len(mv) - 1; i >= 0; i-- {
		tmp = mv[i](tmp)
	}
	return tmp
}

// GetHandlers compiles the routes and handlers for a given router node.
// It wraps each handler with middleware and groups them by type (commands, callbacks, etc.)
// into a StateEntry associated with the router's FSM state.
func (c *Compiler) GetHandlers(r core.Router) {
	tmpCmd := map[string]core.Handler{}
	tmpCallback := map[string]core.Handler{}
	var tmpRegex []core.Regex
	var tmpMsg []core.MsgEquals
	var tmpAny core.Handler

	for key, value := range r.GetCmd() {
		tmpCmd[key] = wrap(r, value)
	}
	for key, value := range r.GetCallback() {
		tmpCallback[key] = wrap(r, value)
	}
	for _, value := range r.GetRegex() {
		value.Method = wrap(r, value.Method)
		tmpRegex = append(tmpRegex, value)
	}
	for _, value := range r.GetMsg() {
		value.Handler = wrap(r, value.Handler)
		tmpMsg = append(tmpMsg, value)
	}
	if r.GetAny() != nil {
		tmpAny = wrap(r, r.GetAny())
	}

	// Add or update the compiled handlers under the router's associated FSM state
	if state, ok := c.States[r.GetState()]; ok {
		state.associationCmd(tmpCmd)
		state.associationCallback(tmpCallback)
		state.associationRegex(tmpRegex)
		state.associationMsg(tmpMsg)
		state.Any = tmpAny
		return
	}

	newState := StateEntry{
		Cmd:      tmpCmd,
		Callback: tmpCallback,
		Regex:    tmpRegex,
		Msg:      tmpMsg,
		Any:      tmpAny,
	}
	c.States[r.GetState()] = &newState
}

// Compile recursively traverses the router tree, starting from the given router,
// compiling handlers for each stateful router.
func (c *Compiler) Compile(r core.Router) {
	c.GetHandlers(r)
	for _, value := range r.GetGroups() {
		c.Compile(value)
	}
}
