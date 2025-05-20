package api

import (
	"bytes"
	"context"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"telegram-sdk/core"
	"telegram-sdk/internal/utils"
	"telegram-sdk/types/commands"
	"telegram-sdk/types/updates"
	"telegram-sdk/values"
)

type Client struct {
	token     string
	transport core.Transport
	baseUrl   url.URL
}

func New(token string, transport core.Transport) *Client {
	return &Client{
		token:     token,
		transport: transport,
		baseUrl: url.URL{
			Scheme: values.SchemeTelegramBot,
			Host:   values.HostTelegramBot,
			Path:   values.PathBotPrefix + token,
		},
	}
}

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

func (c *Client) SendMsg(ctx context.Context, msg *commands.SendMessageRequest) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	u := utils.Endpoint(&c.baseUrl, values.SendMessage)
	body, err := json.Marshal(msg) // заменил Marshal
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
