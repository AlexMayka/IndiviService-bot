package api

import (
	"bytes"
	"context"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"io"
	"net/http"
	"telegram-sdk/internal/utils"
	"telegram-sdk/types/commands"
	"telegram-sdk/values"
)

// SendMsg sends a message using the Telegram Bot API's sendMessage method.
//
// It marshals the given SendMessageRequest into JSON, constructs an HTTP POST request,
// sets appropriate headers, and dispatches it via the configured Transport.
// If the Telegram API returns an error (HTTP status code >= 400), the error message is
// read from the response body and returned.
//
// Parameters:
//   - ctx: context for cancelation and timeout control.
//   - msg: a pointer to SendMessageRequest containing chat ID, text, and optional parameters.
//
// Returns:
//   - error: if marshaling, request creation, sending, or decoding fails.
func (c *Client) SendMsg(ctx context.Context, msg *commands.SendMessageRequest) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	u := utils.Endpoint(&c.baseUrl, values.SendMessage)
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u.String(), bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.transport.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		respData, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("telegram error: %s\n%s", resp.Status, string(respData))
	}

	return nil
}
