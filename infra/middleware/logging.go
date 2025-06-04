// Package middleware provides reusable middleware functions
// that can be applied to handlers during message dispatch in the Telegram SDK.
package middleware

import (
	"telegram-sdk/core"
	"time"
)

// Logging returns a middleware that logs basic information about
// each incoming update processed by a handler.
//
// It logs the chat ID, incoming text, current FSM state, and handler execution duration.
//
// This middleware is useful for simple diagnostics and development-time visibility.
func Logging() core.Middleware {
	return func(next core.Handler) core.Handler {
		return func(ctx core.Context) {
			start := time.Now()

			chatID := ctx.ChatId()
			text := ctx.Text()
			state := ctx.FMS().Get(chatID)

			next(ctx)

			duration := time.Since(start)
			ctx.Logger().Info("handler completed",
				"chat_id", chatID,
				"text", text,
				"state", state,
				"duration_ms", duration.Milliseconds(),
			)
		}
	}
}
