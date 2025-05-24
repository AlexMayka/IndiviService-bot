package bot

import (
	"context"
	"telegram-sdk/core"
	"telegram-sdk/impl/api"
	"telegram-sdk/impl/dispatcher"
	"telegram-sdk/impl/router"
	"telegram-sdk/impl/transport"
	"telegram-sdk/types/updates"
)

type Bot struct {
	token     string
	router    core.Router
	poller    core.Poller
	transport core.Transport
	api       *api.Client
}

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
		b.poller = dispatcher.NewLongPoll(b.api)
	}

	return b
}

func (b *Bot) Run(ctx context.Context) error {
	for {
		updates, err := b.poller.Updates(ctx)
		if err != nil {
			return err
		}
		for _, update := range updates {
			b.dispatch(ctx, update)
		}
	}
}

func (b *Bot) dispatch(ctx context.Context, u updates.Update) {

}
