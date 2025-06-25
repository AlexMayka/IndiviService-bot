package main

import (
	"context"
	"fmt"
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

	// Start command with inline keyboard
	r.Command("/start", func(ctx core.Context) {
		keyboard := commands.NewKeyboard().
			Row(
				commands.Button("📝 Text Features", "text_features"),
				commands.Button("🖼️ Media Features", "media_features"),
			).
			Row(
				commands.URLButton("🌐 GitHub", "https://github.com/AlexMayka/gogram"),
				commands.Button("ℹ️ Help", "help"),
			).
			Build()

		ctx.Send(&commands.SendMessageRequest{
			Text: "🤖 Welcome to Gogram Interactive Bot!\n\nChoose what you'd like to explore:",
			ReplyMarkup: keyboard,
		})
	})

	// Handle callback queries
	r.Callback("text_features", func(ctx core.Context) {
		ctx.AnswerCallback(&commands.AnswerCallbackQueryRequest{
			CallbackQueryID: ctx.CallbackData(),
			Text:           "Exploring text features...",
		})

		keyboard := commands.NewKeyboard().
			Row(
				commands.Button("Bold Text", "format_bold"),
				commands.Button("Italic Text", "format_italic"),
			).
			Row(
				commands.Button("Code Block", "format_code"),
				commands.Button("⬅️ Back", "back_main"),
			).
			Build()

		ctx.EditMessage(&commands.EditMessageTextRequest{
			ChatID:    fmt.Sprintf("%d", ctx.ChatId()),
			MessageID: ctx.MessageId(),
			Text:      "📝 Text Features\n\nSelect a formatting option:",
			ReplyMarkup: keyboard,
		})
	})

	r.Callback("media_features", func(ctx core.Context) {
		ctx.AnswerCallback(&commands.AnswerCallbackQueryRequest{
			CallbackQueryID: ctx.CallbackData(),
			Text:           "Media features coming up!",
		})

		keyboard := commands.NewKeyboard().
			Row(
				commands.Button("📷 Send Photo", "send_photo"),
				commands.Button("📄 Send Document", "send_document"),
			).
			Row(
				commands.Button("⬅️ Back", "back_main"),
			).
			Build()

		ctx.EditMessage(&commands.EditMessageTextRequest{
			ChatID:    fmt.Sprintf("%d", ctx.ChatId()),
			MessageID: ctx.MessageId(),
			Text:      "🖼️ Media Features\n\nSelect a media option:",
			ReplyMarkup: keyboard,
		})
	})

	// Formatting examples
	r.Callback("format_bold", func(ctx core.Context) {
		ctx.AnswerCallback(&commands.AnswerCallbackQueryRequest{
			CallbackQueryID: ctx.CallbackData(),
		})

		ctx.Send(&commands.SendMessageRequest{
			Text:      "*This text is bold!*",
			ParseMode: commands.StyleMarkdownV2,
		})
	})

	r.Callback("format_italic", func(ctx core.Context) {
		ctx.AnswerCallback(&commands.AnswerCallbackQueryRequest{
			CallbackQueryID: ctx.CallbackData(),
		})

		ctx.Send(&commands.SendMessageRequest{
			Text:      "_This text is italic!_",
			ParseMode: commands.StyleMarkdownV2,
		})
	})

	r.Callback("format_code", func(ctx core.Context) {
		ctx.AnswerCallback(&commands.AnswerCallbackQueryRequest{
			CallbackQueryID: ctx.CallbackData(),
		})

		ctx.Send(&commands.SendMessageRequest{
			Text: "```go\nfunc main() {\n    fmt.Println(\"Hello, Gogram!\")\n}\n```",
			ParseMode: commands.StyleMarkdownV2,
		})
	})

	// Media examples (would need actual file URLs or file_ids)
	r.Callback("send_photo", func(ctx core.Context) {
		ctx.AnswerCallback(&commands.AnswerCallbackQueryRequest{
			CallbackQueryID: ctx.CallbackData(),
			Text:           "This would send a photo if we had one!",
			ShowAlert:      true,
		})
	})

	r.Callback("send_document", func(ctx core.Context) {
		ctx.AnswerCallback(&commands.AnswerCallbackQueryRequest{
			CallbackQueryID: ctx.CallbackData(),
			Text:           "This would send a document if we had one!",
			ShowAlert:      true,
		})
	})

	// Help callback
	r.Callback("help", func(ctx core.Context) {
		ctx.AnswerCallback(&commands.AnswerCallbackQueryRequest{
			CallbackQueryID: ctx.CallbackData(),
		})

		helpText := `🤖 *Gogram Interactive Bot Help*

This bot demonstrates the capabilities of the Gogram SDK:

• Inline keyboards with callbacks
• Text formatting (bold, italic, code)
• Message editing
• Chat actions
• Callback query handling

Commands:
/start - Show main menu
/typing - Demonstrate typing action
/info - Show bot information

Built with ❤️ using Gogram SDK`

		ctx.Send(&commands.SendMessageRequest{
			Text:      helpText,
			ParseMode: commands.StyleMarkdownV2,
		})
	})

	// Back to main menu
	r.Callback("back_main", func(ctx core.Context) {
		ctx.AnswerCallback(&commands.AnswerCallbackQueryRequest{
			CallbackQueryID: ctx.CallbackData(),
		})

		keyboard := commands.NewKeyboard().
			Row(
				commands.Button("📝 Text Features", "text_features"),
				commands.Button("🖼️ Media Features", "media_features"),
			).
			Row(
				commands.URLButton("🌐 GitHub", "https://github.com/AlexMayka/gogram"),
				commands.Button("ℹ️ Help", "help"),
			).
			Build()

		ctx.EditMessage(&commands.EditMessageTextRequest{
			ChatID:    fmt.Sprintf("%d", ctx.ChatId()),
			MessageID: ctx.MessageId(),
			Text:      "🤖 Welcome to Gogram Interactive Bot!\n\nChoose what you'd like to explore:",
			ReplyMarkup: keyboard,
		})
	})

	// Typing action demonstration
	r.Command("/typing", func(ctx core.Context) {
		// Show typing action
		ctx.SendChatAction(&commands.SendChatActionRequest{
			ChatID: fmt.Sprintf("%d", ctx.ChatId()),
			Action: commands.ActionTyping,
		})

		// Simulate typing delay (in real implementation, you'd have a delay)
		ctx.Send(&commands.SendMessageRequest{
			Text: "I was just typing! 💬",
		})
	})

	// Bot info command
	r.Command("/info", func(ctx core.Context) {
		infoText := `🤖 *Bot Information*

SDK: Gogram v0\\.1
Features:
• Clean Architecture
• FSM Support
• Middleware Chain
• Regex Routing
• Inline Keyboards
• Rich Telegram API

This bot showcases the power and simplicity of the Gogram SDK\\!`

		ctx.Send(&commands.SendMessageRequest{
			Text:      infoText,
			ParseMode: commands.StyleMarkdownV2,
		})
	})

	// Default handler for unknown messages
	r.Any(func(ctx core.Context) {
		ctx.Send(&commands.SendMessageRequest{
			Text: "🤔 I don't understand that command. Try /start to see what I can do!",
		})
	})

	// Start the bot
	b := runtime.New(token, runtime.WithRouter(r))
	if err := b.Run(context.Background()); err != nil {
		panic(err)
	}
}