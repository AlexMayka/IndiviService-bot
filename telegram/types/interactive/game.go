package interactive

import (
	"telegram-sdk/telegram/types/base"
	"telegram-sdk/telegram/types/media/attachments"
)

type Game struct {
	Title        string                   `json:"title"`
	Description  string                   `json:"description"`
	Photo        []*attachments.PhotoSize `json:"photo"`
	Text         string                   `json:"text"`
	TextEntities []*base.MessageEntity    `json:"text_entities,omitempty"`
	Animation    *attachments.Animation   `json:"animation,omitempty"`
}
