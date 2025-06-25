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
	
	// SendPhoto sends a photo to the current chat.
	SendPhoto(req *commands.SendPhotoRequest) error
	
	// SendDocument sends a document to the current chat.
	SendDocument(req *commands.SendDocumentRequest) error
	
	// EditMessage edits an existing message.
	EditMessage(req *commands.EditMessageTextRequest) error
	
	// DeleteMessage deletes a message.
	DeleteMessage(req *commands.DeleteMessageRequest) error
	
	// SendChatAction sends a chat action (typing, uploading, etc).
	SendChatAction(req *commands.SendChatActionRequest) error
	
	// AnswerCallback answers a callback query.
	AnswerCallback(req *commands.AnswerCallbackQueryRequest) error
	
	// UserId returns the user ID of the current message sender.
	UserId() int
	
	// MessageId returns the message ID of the current update.
	MessageId() int
	
	// CallbackData returns callback data if the update is a callback query.
	CallbackData() string
	
	// IsCallback returns true if the current update is a callback query.
	IsCallback() bool
}
