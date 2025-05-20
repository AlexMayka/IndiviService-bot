package transport

import (
	"net/http"
	"telegram-sdk/core"
	"time"
)

func Default() core.Transport {
	return &http.Client{Timeout: 10 * time.Second}
}
