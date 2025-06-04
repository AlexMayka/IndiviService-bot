// Package poller provides implementations of polling strategies
// to retrieve updates from the Telegram Bot API.
package poller

import (
	"context"
	"telegram-sdk/infra/api"
	"telegram-sdk/types/updates"
)

// LongPoll is a polling strategy that continuously requests updates from Telegram
// using the getUpdates method. It maintains an offset to avoid receiving duplicate updates.
type LongPoll struct {
	api    *api.Client // Client for sending requests to the Telegram Bot API
	offset int         // Last processed update offset
}

// NewLongPoll creates a new LongPoll poller using the given Telegram API client.
func NewLongPoll(api *api.Client) *LongPoll {
	return &LongPoll{api: api}
}

// Updates retrieves new updates from the Telegram Bot API.
//
// It returns only updates with UpdateID > offset and then increments the offset
// to continue from the last received update.
func (l *LongPoll) Updates(ctx context.Context) ([]updates.Update, error) {
	ups, err := l.api.GetUpdates(ctx, l.offset)
	if err != nil {
		return nil, err
	}

	if count := len(ups); count > 0 {
		l.offset = ups[count-1].UpdateId + 1
	}
	return ups, nil
}
