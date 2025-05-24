package main

import (
	"context"
	"fmt"
	"os"
	"telegram-sdk/bot"
	"telegram-sdk/core"
	"telegram-sdk/impl/router"
	"telegram-sdk/types/commands"
	"time"
)

func LoggingMiddleware(next core.Handler) core.Handler {
	return func(ctx core.Context) {
		start := time.Now()
		next(ctx)
		fmt.Printf("[%s] chat=%d, text=%s, duration=%s\n", time.Now().Format(time.RFC3339), ctx.ChatId(), ctx.Text(), time.Since(start))
	}
}

func main() {
	token := os.Getenv("TELEGRAM_TOKEN")

	r := router.New()
	r.Use(LoggingMiddleware)

	admin := r.Group("/admin")

	user := admin.Group("/user")
	user.Command("/start", func(ctx core.Context) {
		ctx.Send(&commands.SendMessageRequest{Text: "Alex"})
	})

	b := bot.New(token, bot.WithRouter(admin))
	b.Run(context.Background())
}
