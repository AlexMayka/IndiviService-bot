package core

import (
	"context"
	"telegram-sdk/types/updates"
)

// Poller defines a mechanism for fetching updates from Telegram servers.
// Implementations may use long polling, webhooks, or other mechanisms.
type Poller interface {
	// Updates retrieves a batch of updates from the Telegram API.
	Updates(ctx context.Context) ([]updates.Update, error)
}
