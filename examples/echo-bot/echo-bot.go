package main

import (
	"context"
	"os"
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

	r := router.New()

	// Start command
	r.Command("/start", func(ctx core.Context) {
		ctx.Send(&commands.SendMessageRequest{
			Text: "👋 Hello! Send me any message and I'll echo it back.",
		})
	})

	// Echo handler
	r.Any(func(ctx core.Context) {
		text := ctx.Text()
		ctx.Send(&commands.SendMessageRequest{
			Text: "🔁 " + text,
		})
	})

	b := runtime.New(token, runtime.WithRouter(r))
	if err := b.Run(context.Background()); err != nil {
		panic(err)
	}
}
