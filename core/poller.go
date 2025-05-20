package core

import (
	"context"
	"telegram-sdk/types/updates"
)

type Poller interface {
	Updates(ctx context.Context) ([]updates.Update, error)
}
