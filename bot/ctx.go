package bot

import (
	"context"
	"strconv"
	"telegram-sdk/core"
	"telegram-sdk/impl/api"
	"telegram-sdk/types/commands"
	"telegram-sdk/types/updates"
)

type CtxBot struct {
	context.Context
	fms core.FMS
	api *api.Client
	m   updates.Message
}

func (ctx *CtxBot) FMS() core.FMS { return ctx.fms }
func (ctx *CtxBot) ChatId() int   { return ctx.m.Chat.Id }
func (ctx *CtxBot) Text() string  { return ctx.m.Text }

func (ctx *CtxBot) Send(msg *commands.SendMessageRequest) error {
	msg.ChatID = strconv.Itoa(ctx.ChatId())
	return ctx.api.SendMsg(ctx.Context, msg)
}
