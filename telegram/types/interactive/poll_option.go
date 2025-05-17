package interactive

import "telegram-sdk/telegram/types/base"

type PollOption struct {
	Text          string               `json:"text"`
	TextParseMode string               `json:"text_parse_mode,omitempty"`
	TextEntities  []base.MessageEntity `json:"text_entities,omitempty"`
}
