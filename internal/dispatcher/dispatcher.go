package dispatcher

import (
	"regexp"
	"strings"
	"telegram-sdk/core"
)

// Dispatcher is responsible for routing updates (wrapped in Context)
// to appropriate handlers according to the current FSM state.
type Dispatcher struct {
	compiler *Compiler
}

// NewDispatcher creates a new Dispatcher with the provided Compiler.
// The Compiler holds precompiled router definitions grouped by FSM states.
func NewDispatcher(compiler *Compiler) *Dispatcher {
	return &Dispatcher{compiler: compiler}
}

// Dispatch routes the incoming message to the correct handler.
//
// It performs routing in the following order within the current FSM state:
//  1. Exact match for command text.
//  2. Exact match for callback data.
//  3. Match against regular expressions.
//  4. Exact match for plain message text.
//  5. Fallback handler (if Any is defined).
//
// If no handler matches, the dispatcher silently exits.
func (b *Dispatcher) Dispatch(ctx core.Context) {
	st := ctx.FMS().Get(ctx.ChatId())

	entry, ok := b.compiler.GetState(st)
	if !ok {
		return // no state entry found
	}

	text := ctx.Text()

	// Handle commands for regular messages
	if handler, ok := entry.Cmd[text]; ok {
		handler(ctx)
		return
	}

	// Handle callback queries
	if ctx.IsCallback() {
		callbackData := ctx.CallbackData()
		if handler, ok := entry.Callback[callbackData]; ok {
			handler(ctx)
			return
		}
	}

	for _, r := range entry.Regex {
		matched, err := regexp.MatchString(r.Pattern, text)
		if err != nil {
			return // regex error (invalid pattern)
		}

		if matched {
			r.Method(ctx)
			return
		}
	}

	for _, m := range entry.Msg {
		if strings.Compare(m.Text, text) == 0 {
			m.Handler(ctx)
			return
		}
	}

	if entry.Any != nil {
		entry.Any(ctx)
	}
}
