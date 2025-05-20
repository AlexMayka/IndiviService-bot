package utils

import "net/url"

func Endpoint(url *url.URL, elem ...string) *url.URL {
	return url.JoinPath(elem...)
}

func AddParams(url *url.URL, params map[string]string) {
	q := url.Query()
	for key, value := range params {
		q.Set(key, value)
	}
	url.RawQuery = q.Encode()
}
