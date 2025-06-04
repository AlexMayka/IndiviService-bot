// Package values defines static configuration constants used across the SDK.
// It contains bot endpoint names, base API settings, and SDK metadata.
package values

import "time"

// Version defines the current SDK version.
const Version = "0.1.0"

// Bot API configuration constants.
const (
	SchemeTelegramBot = "https"
	HostTelegramBot   = "api.telegram.org"
	PathBotPrefix     = "/bot"
)

// Endpoints defines supported Telegram Bot API methods.
const (
	GetMe       = "getMe"
	GetUpdates  = "getUpdates"
	SendMessage = "sendMessage"
)

// Settings defines shared configuration options.
var (
	HTTPTimeout = 10 * time.Second
)
