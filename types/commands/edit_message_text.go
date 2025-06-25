package commands

import "telegram-sdk/types/updates"

// EditMessageTextRequest represents the request body for the editMessageText API call.
type EditMessageTextRequest struct {
	BusinessConnectionId string                        `json:"business_connection_id,omitempty"` // Optional business account ID
	ChatID               string                        `json:"chat_id,omitempty"`                // Required if inline_message_id is not specified
	MessageID            int                           `json:"message_id,omitempty"`             // Required if inline_message_id is not specified
	InlineMessageID      string                        `json:"inline_message_id,omitempty"`      // Required if chat_id and message_id are not specified
	Text                 string                        `json:"text"`                             // New text of the message
	ParseMode            Style                         `json:"parse_mode,omitempty"`             // Text formatting mode
	Entities             []updates.MessageEntity       `json:"entities,omitempty"`               // Explicit text entity formatting
	LinkPreviewOptions   *updates.LinkPreviewOptions   `json:"link_preview_options,omitempty"`   // Control preview behavior
	ReplyMarkup          *updates.InlineKeyboardMarkup `json:"reply_markup,omitempty"`           // Inline keyboard markup
}