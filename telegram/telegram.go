package telegram

import (
	"net/url"
	"telegram-sdk/telegram/values"
)

type Bot struct {
	url url.URL
}

func NewBot(token string) *Bot {
	return &Bot{
		url: url.URL{
			Scheme: values.SchemeTelegramBot,
			Host:   values.HostTelegramBot,
			Path:   values.PathBotPrefix + token,
		},
	}
}

// ToDO: GetMe
// ToDo: Add EndPoint for command
