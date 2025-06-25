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

// EditMessageText edits text of a message using the Telegram Bot API's editMessageText method.
func (c *Client) EditMessageText(ctx context.Context, req *commands.EditMessageTextRequest) error {
	return c.sendJSONRequest(ctx, values.EditMessageText, req)
}

// DeleteMessage deletes a message using the Telegram Bot API's deleteMessage method.
func (c *Client) DeleteMessage(ctx context.Context, req *commands.DeleteMessageRequest) error {
	return c.sendJSONRequest(ctx, values.DeleteMessage, req)
}

// SendPhoto sends a photo using the Telegram Bot API's sendPhoto method.
func (c *Client) SendPhoto(ctx context.Context, req *commands.SendPhotoRequest) error {
	return c.sendJSONRequest(ctx, values.SendPhoto, req)
}

// SendDocument sends a document using the Telegram Bot API's sendDocument method.
func (c *Client) SendDocument(ctx context.Context, req *commands.SendDocumentRequest) error {
	return c.sendJSONRequest(ctx, values.SendDocument, req)
}

// SendChatAction sends a chat action using the Telegram Bot API's sendChatAction method.
func (c *Client) SendChatAction(ctx context.Context, req *commands.SendChatActionRequest) error {
	return c.sendJSONRequest(ctx, values.SendChatAction, req)
}

// AnswerCallbackQuery answers a callback query using the Telegram Bot API's answerCallbackQuery method.
func (c *Client) AnswerCallbackQuery(ctx context.Context, req *commands.AnswerCallbackQueryRequest) error {
	return c.sendJSONRequest(ctx, values.AnswerCallbackQuery, req)
}

// GetChat gets information about a chat using the Telegram Bot API's getChat method.
func (c *Client) GetChat(ctx context.Context, req *commands.GetChatRequest) error {
	return c.sendJSONRequest(ctx, values.GetChat, req)
}

// SetMyCommands sets the list of the bot's commands using the Telegram Bot API's setMyCommands method.
func (c *Client) SetMyCommands(ctx context.Context, req *commands.SetMyCommandsRequest) error {
	return c.sendJSONRequest(ctx, values.SetMyCommands, req)
}

// sendJSONRequest is a helper method to send JSON requests to the Telegram API.
func (c *Client) sendJSONRequest(ctx context.Context, endpoint string, payload interface{}) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	u := utils.Endpoint(&c.baseUrl, endpoint)
	body, err := json.Marshal(payload)
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
