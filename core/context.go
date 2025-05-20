package core

import (
	"context"
	"telegram-sdk/types/commands"
)

type Context interface {
	context.Context
	ChatId() int
	Text() string
	Send(msg *commands.SendMessageRequest) error
}
