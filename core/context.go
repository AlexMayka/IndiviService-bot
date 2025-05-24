package core

import (
	"context"
	"telegram-sdk/types/commands"
)

type Context interface {
	context.Context
	FMS() FMS

	ChatId() int
	Text() string
	Send(msg *commands.SendMessageRequest) error
}
