package commands

import "telegram-sdk/types/updates"

// SendDocumentRequest represents the request body for the sendDocument API call.
type SendDocumentRequest struct {
	BusinessConnectionId    string                        `json:"business_connection_id,omitempty"`     // Optional business account ID
	ChatID                  string                        `json:"chat_id"`                              // Target chat ID
	MessageThreadId         int                           `json:"message_thread_id,omitempty"`          // For topics in forums
	Document                InputFile                     `json:"document"`                             // File to send
	Thumbnail               InputFile                     `json:"thumbnail,omitempty"`                  // Thumbnail of the file
	Caption                 string                        `json:"caption,omitempty"`                    // Document caption
	ParseMode               Style                         `json:"parse_mode,omitempty"`                 // Text formatting mode for caption
	CaptionEntities         []updates.MessageEntity       `json:"caption_entities,omitempty"`           // Explicit caption entity formatting
	DisableContentTypeDetection bool                      `json:"disable_content_type_detection,omitempty"` // Disables automatic server-side content type detection
	DisableNotification     bool                          `json:"disable_notification,omitempty"`       // Send silently
	ProtectContent          bool                          `json:"protect_content,omitempty"`            // Prevent message forwarding
	AllowPaidBroadcast      bool                          `json:"allow_paid_broadcast,omitempty"`       // Paid promotion permission
	MessageEffectId         string                        `json:"message_effect_id,omitempty"`          // Optional visual effect
	ReplyParameters         *ReplyParameters              `json:"reply_parameters,omitempty"`           // Options for replying
	ReplyMarkup             *updates.InlineKeyboardMarkup `json:"reply_markup,omitempty"`               // Inline keyboard markup
}