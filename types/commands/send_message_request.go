// Package commands defines the structure of outgoing Telegram Bot API requests,
// such as sending messages, and associated styling and reply parameters.
package commands

import "telegram-sdk/types/updates"

// Style represents the format in which the message text should be parsed.
type Style = string

// Available text formatting styles for messages.
const (
	StyleHTML       Style = "HTML"       // HTML-style formatting
	StyleMarkdown   Style = "Markdown"   // Markdown-style formatting (legacy)
	StyleMarkdownV2 Style = "MarkdownV2" // Modern Markdown V2 formatting
)

// ReplyParameters specifies how a sent message should reply to an existing message.
type ReplyParameters struct {
	MessageID                int                     `json:"message_id"`                            // ID of the original message being replied to
	ChatID                   any                     `json:"chat_id,omitempty"`                     // Optional: target chat for reply
	AllowSendingWithoutReply bool                    `json:"allow_sending_without_reply,omitempty"` // Whether the reply should be sent if the original message is missing
	Quote                    string                  `json:"quote,omitempty"`                       // Optional custom quote text
	QuoteParseMode           string                  `json:"quote_parse_mode,omitempty"`            // Format of quote text (HTML, Markdown, etc.)
	QuoteEntities            []updates.MessageEntity `json:"quote_entities,omitempty"`              // Parsed formatting of quote
	QuotePosition            int                     `json:"quote_position,omitempty"`              // Position of the quoted message text
}

// SendMessageRequest represents the request body for the sendMessage API call.
type SendMessageRequest struct {
	BusinessConnectionId string                        `json:"business_connection_id,omitempty"` // Optional business account ID
	ChatID               string                        `json:"chat_id"`                          // Target chat ID
	MessageThreadId      int                           `json:"message_thread_id,omitempty"`      // For topics in forums
	Text                 string                        `json:"text"`                             // Message text
	ParseMode            Style                         `json:"parse_mode,omitempty"`             // Text formatting mode
	Entities             []updates.MessageEntity       `json:"entities,omitempty"`               // Explicit text entity formatting
	LinkPreviewOptions   *updates.LinkPreviewOptions   `json:"link_preview_options,omitempty"`   // Control preview behavior
	DisableNotification  bool                          `json:"disable_notification,omitempty"`   // Send silently
	ProtectContent       bool                          `json:"protect_content,omitempty"`        // Prevent message forwarding
	AllowPaidBroadcast   bool                          `json:"allow_paid_broadcast,omitempty"`   // Paid promotion permission
	MessageEffectId      string                        `json:"message_effect_id,omitempty"`      // Optional visual effect
	ReplyParameters      *ReplyParameters              `json:"reply_parameters,omitempty"`       // Options for replying
	ReplyMarkup          *updates.InlineKeyboardMarkup `json:"reply_markup,omitempty"`           // Inline keyboard or custom reply markup
}
