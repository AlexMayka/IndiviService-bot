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
func (ctx *CtxBot) ChatId() int { return ctx.Update.Message.Chat.Id }

// Text returns the textual content of the incoming message.
func (ctx *CtxBot) Text() string { return ctx.Update.Message.Text }

// Logger returns the logger instance for this context.
func (ctx *CtxBot) Logger() core.Logger { return ctx.Log }

// FMS returns the finite state machine instance.
func (ctx *CtxBot) FMS() core.FMS { return ctx.Fms }

// Send sends a message back to the current chat.
func (ctx *CtxBot) Send(msg *commands.SendMessageRequest) error {
	msg.ChatID = strconv.Itoa(ctx.ChatId())
	return ctx.API.SendMsg(ctx.Context, msg)
}
