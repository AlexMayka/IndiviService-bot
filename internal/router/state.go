package router

import "telegram-sdk/core"

type regex struct {
	pattern string
	method  core.Handler
}

type msgEquals struct {
	text    string
	handler core.Handler
}

type StateEntry struct {
	cmd      map[string]core.Handler
	callback map[string]core.Handler
	regex    []regex
	msg      []msgEquals
	any      core.Handler
}
