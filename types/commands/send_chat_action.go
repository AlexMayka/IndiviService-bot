package commands

// ChatAction represents the type of action to broadcast
type ChatAction = string

// Available chat actions for sendChatAction
const (
	ActionTyping          ChatAction = "typing"           // Typing text message
	ActionUploadPhoto     ChatAction = "upload_photo"     // Uploading photo
	ActionRecordVideo     ChatAction = "record_video"     // Recording video
	ActionUploadVideo     ChatAction = "upload_video"     // Uploading video
	ActionRecordVoice     ChatAction = "record_voice"     // Recording voice message
	ActionUploadVoice     ChatAction = "upload_voice"     // Uploading voice message
	ActionUploadDocument  ChatAction = "upload_document"  // Uploading document
	ActionChooseSticker   ChatAction = "choose_sticker"   // Choosing sticker
	ActionFindLocation    ChatAction = "find_location"    // Finding location
	ActionRecordVideoNote ChatAction = "record_video_note" // Recording video note
	ActionUploadVideoNote ChatAction = "upload_video_note" // Uploading video note
)

// SendChatActionRequest represents the request body for the sendChatAction API call.
type SendChatActionRequest struct {
	BusinessConnectionId string     `json:"business_connection_id,omitempty"` // Optional business account ID
	ChatID               string     `json:"chat_id"`                          // Target chat ID
	MessageThreadId      int        `json:"message_thread_id,omitempty"`      // For topics in forums
	Action               ChatAction `json:"action"`                           // Type of action to broadcast
}