package commands

// BotCommand represents a bot command with description
type BotCommand struct {
	Command     string `json:"command"`     // Text of the command; 1-32 characters
	Description string `json:"description"` // Description of the command; 1-256 characters
}

// BotCommandScope represents the scope of bot commands
type BotCommandScope interface{}

// SetMyCommandsRequest represents the request body for the setMyCommands API call.
type SetMyCommandsRequest struct {
	Commands     []BotCommand     `json:"commands"`               // List of bot commands to be set
	Scope        BotCommandScope  `json:"scope,omitempty"`        // JSON-serialized object, describing scope of users for which the commands are relevant
	LanguageCode string          `json:"language_code,omitempty"` // Two-letter ISO 639-1 language code
}