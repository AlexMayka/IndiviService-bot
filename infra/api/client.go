package api

import (
	"net/url"
	"telegram-sdk/core"
	"telegram-sdk/values"
)

// Client represents a low-level HTTP client for the Telegram Bot API.
// It holds the bot token, an HTTP transport, and the base URL for requests.
type Client struct {
	token     string
	transport core.Transport
	baseUrl   url.URL
}

// New creates and returns a new instance of Client using the given bot token and transport.
//
// The baseUrl is composed using predefined constants for scheme, host, and API path prefix.
func New(token string, transport core.Transport) *Client {
	return &Client{
		token:     token,
		transport: transport,
		baseUrl: url.URL{
			Scheme: values.SchemeTelegramBot,
			Host:   values.HostTelegramBot,
			Path:   values.PathBotPrefix + token,
		},
	}
}
