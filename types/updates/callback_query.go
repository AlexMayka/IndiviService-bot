package updates

// CallbackQuery represents an incoming callback query from a callback button in an inline keyboard
type CallbackQuery struct {
	ID              string   `json:"id"`                          // Unique identifier for this query
	From            *User    `json:"from"`                        // Sender of the query
	Message         *Message `json:"message,omitempty"`           // Message with the callback button that originated the query
	InlineMessageID string   `json:"inline_message_id,omitempty"` // Identifier of the message sent via the bot in inline mode
	ChatInstance    string   `json:"chat_instance"`               // Global identifier, uniquely corresponding to the chat to which the message was sent
	Data            string   `json:"data,omitempty"`              // Data associated with the callback button
	GameShortName   string   `json:"game_short_name,omitempty"`   // Short name of a Game to be returned
}