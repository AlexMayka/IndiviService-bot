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

	r.Command("/start", func(ctx core.Context) {
		err := ctx.Send(&commands.SendMessageRequest{Text: "AlexMayka"})
		if err != nil {
			fmt.Println(err)
		}
	})

	r.OnMsg(func(ctx core.Context) {
		err := ctx.Send(&commands.SendMessageRequest{Text: ctx.Text()})
		if err != nil {
			fmt.Println(err)
		}
	})

	b := bot.New(token, bot.WithRouter(r))
	b.Run(context.Background())
}
