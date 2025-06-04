// Package utils provides helper functions for working with URLs,
// including endpoint construction and query parameter manipulation.
package utils

import "net/url"

// Endpoint returns a new URL by joining the base URL with additional path segments.
//
// Example:
//
//	base := &url.URL{Scheme: "https", Host: "api.telegram.org", Path: "/bot<token>"}
//	full := Endpoint(base, "sendMessage") // → https://api.telegram.org/bot<token>/sendMessage
func Endpoint(base *url.URL, elem ...string) *url.URL {
	return base.JoinPath(elem...)
}

// AddParams adds query parameters to the given URL in-place.
//
// Example:
//
//	u := &url.URL{Scheme: "https", Host: "api.telegram.org", Path: "/bot<token>/getUpdates"}
//	AddParams(u, map[string]string{"offset": "123"})
//	// Result: https://api.telegram.org/bot<token>/getUpdates?offset=123
func AddParams(u *url.URL, params map[string]string) {
	q := u.Query()
	for key, value := range params {
		q.Set(key, value)
	}
	u.RawQuery = q.Encode()
}
