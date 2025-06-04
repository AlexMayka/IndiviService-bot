package api

import (
	"context"
	jsoniter "github.com/json-iterator/go"
	"net/http"
	"strconv"
	"telegram-sdk/internal/utils"
	"telegram-sdk/types/updates"
	"telegram-sdk/values"
)

// GetMe checks the validity of the Telegram bot token by calling the getMe endpoint.
// It returns true if the API responds with a 200 OK status, false otherwise.
func (c *Client) GetMe(ctx context.Context) bool {
	u := utils.Endpoint(&c.baseUrl, values.GetMe)

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	resp, err := c.transport.Do(req)

	if err != nil || resp.StatusCode != http.StatusOK {
		return false
	}

	return true
}

// GetUpdates retrieves incoming updates for the bot using the long polling mechanism.
// The 'offset' parameter is used to avoid receiving duplicate updates.
//
// It returns a slice of updates.Update structs or an error if the HTTP request or decoding fails.
func (c *Client) GetUpdates(ctx context.Context, offset int) ([]updates.Update, error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	u := utils.Endpoint(&c.baseUrl, values.GetUpdates)
	params := map[string]string{
		"offset": strconv.Itoa(offset),
	}
	utils.AddParams(u, params)

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	resp, err := c.transport.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	res := &updates.Updates{}
	if err := json.NewDecoder(resp.Body).Decode(res); err != nil {
		return nil, err
	}
	return res.Result, nil
}
