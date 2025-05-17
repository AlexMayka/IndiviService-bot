package others

import "telegram-sdk/telegram/types/base"

type TextQuote struct {
	Text     string                `json:"text"`
	Entities []*base.MessageEntity `json:"entities,omitempty"`
	Position int                   `json:"position"`
	IsManual bool                  `json:"is_manual,omitempty"`
}
