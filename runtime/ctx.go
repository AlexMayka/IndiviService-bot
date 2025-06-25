package runtime

import (
	"context"
	"strconv"
	"telegram-sdk/core"
	"telegram-sdk/infra/api"
	"telegram-sdk/types/commands"
	"telegram-sdk/types/updates"
)

// CtxBot implements the core.Context interface and wraps Telegram update data with shared bot services.
// It carries access to API, FSM, logger and the incoming update itself.
type CtxBot struct {
	context.Context
	Update *updates.Update
	API    *api.Client
	Log    core.Logger
	Fms    core.FMS
}

// ChatId returns the numeric ID of the chat the current update originated from.
func (ctx *CtxBot) ChatId() int {
	if ctx.Update.Message != nil {
		return ctx.Update.Message.Chat.Id
	}
	if ctx.Update.CallbackQuery != nil && ctx.Update.CallbackQuery.Message != nil {
		return ctx.Update.CallbackQuery.Message.Chat.Id
	}
	return 0
}

// Text returns the textual content of the incoming message.
func (ctx *CtxBot) Text() string {
	if ctx.Update.Message != nil {
		return ctx.Update.Message.Text
	}
	return ""
}

// Logger returns the logger instance for this context.
func (ctx *CtxBot) Logger() core.Logger { return ctx.Log }

// FMS returns the finite state machine instance.
func (ctx *CtxBot) FMS() core.FMS { return ctx.Fms }

// Send sends a message back to the current chat.
func (ctx *CtxBot) Send(msg *commands.SendMessageRequest) error {
	msg.ChatID = strconv.Itoa(ctx.ChatId())
	return ctx.API.SendMsg(ctx.Context, msg)
}

// SendPhoto sends a photo to the current chat.
func (ctx *CtxBot) SendPhoto(req *commands.SendPhotoRequest) error {
	req.ChatID = strconv.Itoa(ctx.ChatId())
	return ctx.API.SendPhoto(ctx.Context, req)
}

// SendDocument sends a document to the current chat.
func (ctx *CtxBot) SendDocument(req *commands.SendDocumentRequest) error {
	req.ChatID = strconv.Itoa(ctx.ChatId())
	return ctx.API.SendDocument(ctx.Context, req)
}

// EditMessage edits an existing message.
func (ctx *CtxBot) EditMessage(req *commands.EditMessageTextRequest) error {
	if req.ChatID == "" {
		req.ChatID = strconv.Itoa(ctx.ChatId())
	}
	return ctx.API.EditMessageText(ctx.Context, req)
}

// DeleteMessage deletes a message.
func (ctx *CtxBot) DeleteMessage(req *commands.DeleteMessageRequest) error {
	req.ChatID = strconv.Itoa(ctx.ChatId())
	return ctx.API.DeleteMessage(ctx.Context, req)
}

// SendChatAction sends a chat action (typing, uploading, etc).
func (ctx *CtxBot) SendChatAction(req *commands.SendChatActionRequest) error {
	req.ChatID = strconv.Itoa(ctx.ChatId())
	return ctx.API.SendChatAction(ctx.Context, req)
}

// AnswerCallback answers a callback query.
func (ctx *CtxBot) AnswerCallback(req *commands.AnswerCallbackQueryRequest) error {
	return ctx.API.AnswerCallbackQuery(ctx.Context, req)
}

// UserId returns the user ID of the current message sender.
func (ctx *CtxBot) UserId() int {
	if ctx.Update.Message != nil && ctx.Update.Message.From != nil {
		return ctx.Update.Message.From.Id
	}
	if ctx.Update.CallbackQuery != nil && ctx.Update.CallbackQuery.From != nil {
		return ctx.Update.CallbackQuery.From.Id
	}
	return 0
}

// MessageId returns the message ID of the current update.
func (ctx *CtxBot) MessageId() int {
	if ctx.Update.Message != nil {
		return ctx.Update.Message.MessageID
	}
	if ctx.Update.CallbackQuery != nil && ctx.Update.CallbackQuery.Message != nil {
		return ctx.Update.CallbackQuery.Message.MessageID
	}
	return 0
}

// CallbackData returns callback data if the update is a callback query.
func (ctx *CtxBot) CallbackData() string {
	if ctx.Update.CallbackQuery != nil {
		return ctx.Update.CallbackQuery.Data
	}
	return ""
}

// IsCallback returns true if the current update is a callback query.
func (ctx *CtxBot) IsCallback() bool {
	return ctx.Update.CallbackQuery != nil
}
