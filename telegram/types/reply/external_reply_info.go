package reply

import (
	"telegram-sdk/telegram/types/base"
	"telegram-sdk/telegram/types/interactive"
	"telegram-sdk/telegram/types/media/attachments"
	"telegram-sdk/telegram/types/media/giveaway"
	"telegram-sdk/telegram/types/media/others"
	"telegram-sdk/telegram/types/payment"
	"telegram-sdk/telegram/types/preview"
)

type ExternalReplyInfo struct {
	Origin             *base.MessageOrigin         `json:"origin"`
	Chat               *base.Chat                  `json:"chat,omitempty"`
	MessageID          int                         `json:"message_id,omitempty"`
	LinkPreviewOptions *preview.LinkPreviewOptions `json:"link_preview_options,omitempty"`
	Animation          *attachments.Animation      `json:"animation,omitempty"`
	Audio              *attachments.Audio          `json:"audio,omitempty"`
	Document           *attachments.Document       `json:"document,omitempty"`
	PaidMedia          *others.PaidMediaInfo       `json:"paid_media,omitempty"`
	Photo              *[]attachments.PhotoSize    `json:"photo,omitempty"`
	Sticker            *attachments.Sticker        `json:"sticker,omitempty"`
	Story              *others.Story               `json:"story,omitempty"`
	Video              *attachments.Video          `json:"video,omitempty"`
	VideoNote          *attachments.VideoNote      `json:"video_note,omitempty"`
	Voice              *attachments.Voice          `json:"voice,omitempty"`
	HasMediaSpoiler    bool                        `json:"has_media_spoiler,omitempty"`
	Contact            *interactive.Contact        `json:"contact,omitempty"`
	Dice               *interactive.Dice           `json:"dice,omitempty"`
	Game               *interactive.Game           `json:"game,omitempty"`
	Giveaway           *giveaway.GiveAway          `json:"giveaway,omitempty"`
	GiveawayWinners    *giveaway.GiveawayWinners   `json:"giveaway_winners,omitempty"`
	Invoice            *payment.Invoice            `json:"invoice,omitempty"`
	Location           *others.Location            `json:"location,omitempty"`
	Poll               *others.Poll                `json:"poll,omitempty"`
	Venue              *others.Venue               `json:"venue,omitempty"`
}
