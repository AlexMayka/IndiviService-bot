package commands

// AnswerCallbackQueryRequest represents the request body for the answerCallbackQuery API call.
type AnswerCallbackQueryRequest struct {
	CallbackQueryID string `json:"callback_query_id"`           // Unique identifier for the query to be answered
	Text            string `json:"text,omitempty"`              // Text of the notification
	ShowAlert       bool   `json:"show_alert,omitempty"`        // If true, an alert will be shown by the client instead of notification
	URL             string `json:"url,omitempty"`               // URL that will be opened by the user's client
	CacheTime       int    `json:"cache_time,omitempty"`        // Maximum amount of time in seconds that result may be cached client-side
}