package dispatcher

import (
	"context"
	"telegram-sdk/impl/api"
	"telegram-sdk/types/updates"
)

type LongPoll struct {
	api    *api.Client
	offset int
}

func NewLongPoll(api *api.Client) *LongPoll { return &LongPoll{api: api} }

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
