package commands

// GetChatRequest represents the request body for the getChat API call.
type GetChatRequest struct {
	ChatID string `json:"chat_id"` // Unique identifier for the target chat
}