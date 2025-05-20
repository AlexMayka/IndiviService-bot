package bot

import "telegram-sdk/core"

type Option func(*Bot)

func WithRouter(r core.Router) Option       { return func(bot *Bot) { bot.router = r } }
func WithPoller(p core.Poller) Option       { return func(bot *Bot) { bot.poller = p } }
func WithTransport(t core.Transport) Option { return func(bot *Bot) { bot.transport = t } }
