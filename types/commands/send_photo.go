package commands

import "telegram-sdk/types/updates"

// InputFile represents a file to be sent, can be either a file_id, URL, or file data
type InputFile interface{}

// SendPhotoRequest represents the request body for the sendPhoto API call.
type SendPhotoRequest struct {
	BusinessConnectionId string                        `json:"business_connection_id,omitempty"` // Optional business account ID
	ChatID               string                        `json:"chat_id"`                          // Target chat ID
	MessageThreadId      int                           `json:"message_thread_id,omitempty"`      // For topics in forums
	Photo                InputFile                     `json:"photo"`                            // Photo to send
	Caption              string                        `json:"caption,omitempty"`                // Photo caption
	ParseMode            Style                         `json:"parse_mode,omitempty"`             // Text formatting mode for caption
	CaptionEntities      []updates.MessageEntity       `json:"caption_entities,omitempty"`       // Explicit caption entity formatting
	ShowCaptionAboveMedia bool                         `json:"show_caption_above_media,omitempty"` // Show caption above the media
	HasSpoiler           bool                          `json:"has_spoiler,omitempty"`            // True if the photo needs to be covered with a spoiler animation
	DisableNotification  bool                          `json:"disable_notification,omitempty"`   // Send silently
	ProtectContent       bool                          `json:"protect_content,omitempty"`        // Prevent message forwarding
	AllowPaidBroadcast   bool                          `json:"allow_paid_broadcast,omitempty"`   // Paid promotion permission
	MessageEffectId      string                        `json:"message_effect_id,omitempty"`      // Optional visual effect
	ReplyParameters      *ReplyParameters              `json:"reply_parameters,omitempty"`       // Options for replying
	ReplyMarkup          *updates.InlineKeyboardMarkup `json:"reply_markup,omitempty"`           // Inline keyboard markup
}