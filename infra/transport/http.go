// Package transport provides an implementation of core.Transport using the standard net/http client.
package transport

import (
	"net/http"
	"telegram-sdk/core"
	"telegram-sdk/values"
)

// Default returns a default HTTP client that implements the core.Transport interface.
// It is configured with a timeout defined in values.HTTPTimeout.
// This function is typically used to provide the transport layer for API requests to Telegram.
func Default() core.Transport {
	return &http.Client{Timeout: values.HTTPTimeout}
}
