package others

import "telegram-sdk/telegram/types/base"

type Story struct {
	Chat base.Chat `json:"chat"`
	Id   int       `json:"id"`
}
