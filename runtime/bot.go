// Package runtime provides the main entry point for running a Telegram bot.
// It connects all core components (router, transport, poller, FSM, logger) into a fully-functional bot instance.
package runtime

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"telegram-sdk/core"
	"telegram-sdk/infra/api"
	"telegram-sdk/infra/fms"
	"telegram-sdk/infra/logger"
	"telegram-sdk/infra/middleware"
	"telegram-sdk/infra/poller"
	"telegram-sdk/infra/router"
	"telegram-sdk/infra/transport"
	"telegram-sdk/internal/dispatcher"
	"telegram-sdk/values"
)

// Bot encapsulates a Telegram bot runtime instance.
// It holds all the necessary components to poll updates, route them, and handle user context.
type Bot struct {
	token     string
	router    core.Router
	poller    core.Poller
	transport core.Transport
	api       *api.Client
	fms       core.FMS
	log       core.Logger

	dispatcher *dispatcher.Dispatcher
}

// New creates a new bot instance using the provided token and options.
// Options allow customizing router, poller, logger, FMS, etc.
func New(token string, opts ...Option) *Bot {
	b := &Bot{token: token}
	for _, o := range opts {
		o(b)
	}

	if b.transport == nil {
		b.transport = transport.Default()
	}

	if b.api == nil {
		b.api = api.New(token, b.transport)
	}

	if b.router == nil {
		b.router = router.New()
	}

	if b.poller == nil {
		b.poller = poller.NewLongPoll(b.api)
	}

	if b.fms == nil {
		b.fms = fms.NewLocalFMS()
	}

	if b.log == nil {
		h := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: false})
		b.log = logger.NewLogger(h)
	}

	// Default logging middleware for incoming messages
	b.router.Use(middleware.Logging())

	return b
}

// Run starts the polling loop and dispatches incoming Telegram updates.
// It blocks indefinitely unless an error occurs or context is cancelled
func (b *Bot) Run(ctx context.Context) error {
	b.log.Info("bot started", "token_trunc", b.token[:6], "build", values.Version)

	if ok := b.api.GetMe(ctx); !ok {
		b.log.Error("Bad token", "token_trunc", b.token[:6], "build", values.Version)
		return errors.New("token_trunc")
	}

	b.log.Info("valid token", "token_trunc", b.token[:6])

	compiler := dispatcher.NewCompiler()
	compiler.Compile(b.router)
	b.log.Info("compile done")

	b.dispatcher = dispatcher.NewDispatcher(compiler)

	for {
		updates, err := b.poller.Updates(ctx)
		if err != nil {
			return err
		}
		for _, update := range updates {
			c := &CtxBot{Context: ctx, API: b.api, Update: &update, Fms: b.fms, Log: b.log}
			b.dispatcher.Dispatch(c)
		}
	}
}
