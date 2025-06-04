package dispatcher

import (
	"telegram-sdk/core"
)

// StateEntry stores all possible handlers for a specific FSM state.
//
// Each field corresponds to a different kind of Telegram update,
// allowing the Dispatcher to select the appropriate handler at runtime.
type StateEntry struct {
	// Cmd maps exact command strings (e.g. "/start") to handlers.
	Cmd map[string]core.Handler

	// Callback maps exact callback data strings to handlers.
	Callback map[string]core.Handler

	// Regex holds pattern-handler pairs for messages matched by regular expressions.
	Regex []core.Regex

	// Msg stores handlers for exact message text (non-command, non-callback).
	Msg []core.MsgEquals

	// Any is a fallback handler that runs if no other handler matches.
	Any core.Handler
}

// associationCmd merges the provided command handlers into the current state entry.
func (s *StateEntry) associationCmd(input map[string]core.Handler) {
	for k, v := range input {
		s.Cmd[k] = v
	}
}

// associationCallback merges the provided callback handlers into the current state entry.
func (s *StateEntry) associationCallback(input map[string]core.Handler) {
	for k, v := range input {
		s.Callback[k] = v
	}
}

// associationRegex appends the provided regex handlers to the current state entry.
func (s *StateEntry) associationRegex(input []core.Regex) {
	s.Regex = append(s.Regex, input...)
}

// associationMsg appends the provided exact message handlers to the current state entry.
func (s *StateEntry) associationMsg(input []core.MsgEquals) {
	s.Msg = append(s.Msg, input...)
}
