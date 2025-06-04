package core

import "net/http"

// Transport defines a minimal interface for executing HTTP requests.
// It allows SDK consumers to inject custom HTTP clients.
type Transport interface {
	// Do executes the given HTTP request and returns the HTTP response.
	Do(req *http.Request) (*http.Response, error)
}
