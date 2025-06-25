package commands

// DeleteMessageRequest represents the request body for the deleteMessage API call.
type DeleteMessageRequest struct {
	ChatID    string `json:"chat_id"`    // Unique identifier for the target chat
	MessageID int    `json:"message_id"` // Identifier of the message to delete
}