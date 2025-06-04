package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"telegram-sdk/core"
	"telegram-sdk/infra/router"
	"telegram-sdk/runtime"
	"telegram-sdk/types/commands"
)

func main() {
	token := os.Getenv("TELEGRAM_TOKEN")
	if token == "" {
		panic("TELEGRAM_TOKEN is not set")
	}

	adminIDStr := os.Getenv("TELEGRAM_ADMIN_ID")
	if adminIDStr == "" {
		panic("TELEGRAM_ADMIN_ID is not set")
	}
	adminID, err := strconv.Atoi(adminIDStr)
	if err != nil {
		panic("TELEGRAM_ADMIN_ID must be a valid integer")
	}

	root := router.New()

	// User flow
	user := root.Group("/user")
	user.Command("/start", func(ctx core.Context) {
		ctx.Send(&commands.SendMessageRequest{Text: "Hi! Please enter your first name."})
		ctx.FMS().Set(ctx.ChatId(), "awaiting_name")
	})

	nameStep := root.Group("/name").SetState("awaiting_name")
	nameStep.Any(func(ctx core.Context) {
		ctx.FMS().SetParam(ctx.ChatId(), "name", ctx.Text())
		ctx.FMS().Set(ctx.ChatId(), "awaiting_lastname")
		ctx.Send(&commands.SendMessageRequest{Text: "Now enter your last name."})
	})

	lastnameStep := nameStep.Group("/lastname").SetState("awaiting_lastname")
	lastnameStep.Any(func(ctx core.Context) {
		name := ctx.FMS().GetParam(ctx.ChatId(), "name")
		lastName := ctx.Text()
		ctx.Send(&commands.SendMessageRequest{
			Text: fmt.Sprintf("Your full name is: %s %s", name, lastName),
		})
		ctx.FMS().Set(ctx.ChatId(), "")
	})

	// Admin-only route with middleware
	admin := root.Group("/admin").Use(AllowOnly(adminID))
	admin.Command("/ping", func(ctx core.Context) {
		ctx.Send(&commands.SendMessageRequest{Text: "pong"})
	})

	// Start the bot
	b := runtime.New(token, runtime.WithRouter(root))
	if err := b.Run(context.Background()); err != nil {
		panic(err)
	}
}

// AllowOnly is a middleware that restricts access to a specific chat ID
func AllowOnly(chatID int) core.Middleware {
	return func(next core.Handler) core.Handler {
		return func(ctx core.Context) {
			if ctx.ChatId() != chatID {
				ctx.Send(&commands.SendMessageRequest{Text: "Access denied"})
				return
			}
			next(ctx)
		}
	}
}
