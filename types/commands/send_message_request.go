package commands

import "telegram-sdk/types/updates"

type Style = string

var (
	StyleHTML       Style = "HTML"
	StyleMarkdown   Style = "Markdown"
	StyleMarkdownV2 Style = "MarkdownV2"
)

type ReplyParameters struct {
	MessageID                int                     `json:"message_id"`
	ChatID                   any                     `json:"chat_id,omitempty"`
	AllowSendingWithoutReply bool                    `json:"allow_sending_without_reply,omitempty"`
	Quote                    string                  `json:"quote,omitempty"`
	QuoteParseMode           string                  `json:"quote_parse_mode,omitempty"`
	QuoteEntities            []updates.MessageEntity `json:"quote_entities,omitempty"`
	QuotePosition            int                     `json:"quote_position,omitempty"`
}

type SendMessageRequest struct {
	BusinessConnectionId string                        `json:"business_connection_id,omitempty"`
	ChatID               string                        `json:"chat_id"`
	MessageThreadId      int                           `json:"message_thread_id,omitempty"`
	Text                 string                        `json:"text"`
	ParseMode            Style                         `json:"parse_mode,omitempty"`
	Entities             []updates.MessageEntity       `json:"entities,omitempty"`
	LinkPreviewOptions   *updates.LinkPreviewOptions   `json:"link_preview_options,omitempty"`
	DisableNotification  bool                          `json:"disable_notification,omitempty"`
	ProtectContent       bool                          `json:"protect_content,omitempty"`
	AllowPaidBroadcast   bool                          `json:"allow_paid_broadcast,omitempty"`
	MessageEffectId      string                        `json:"message_effect_id,omitempty"`
	ReplyParameters      *ReplyParameters              `json:"reply_parameters,omitempty"`
	ReplyMarkup          *updates.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}
