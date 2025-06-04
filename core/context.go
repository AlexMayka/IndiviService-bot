package core

import (
	"context"
	"telegram-sdk/types/commands"
)

// Context provides request-scoped data and behavior for a single update.
// It wraps the base context.Context and enriches it with bot-specific methods
// such as logging, message sending, chat state, and FSM access.
type Context interface {
	context.Context

	// FMS returns the finite state machine associated with the current user.
	FMS() FMS

	// Logger returns the logger instance used within this context.
	Logger() Logger

	// ChatId returns the chat ID of the current message sender.
	ChatId() int

	// Text returns the text content of the incoming message.
	Text() string

	// Send sends a message to the current chat.
	Send(msg *commands.SendMessageRequest) error
}
