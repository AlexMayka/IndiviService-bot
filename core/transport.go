package core

import "net/http"

type Transport interface {
	Do(req *http.Request) (*http.Response, error)
}
