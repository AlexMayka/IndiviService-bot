package runtime

import "telegram-sdk/core"

// Option defines a function that modifies the configuration of a Bot.
// This pattern allows flexible dependency injection for router, poller, logger, etc.
type Option func(*Bot)

// WithRouter sets a custom Router implementation.
func WithRouter(r core.Router) Option { return func(bot *Bot) { bot.router = r } }

// WithPoller sets a custom Poller implementation.
func WithPoller(p core.Poller) Option { return func(bot *Bot) { bot.poller = p } }

// WithTransport sets a custom Transport implementation (e.g., HTTP client).
func WithTransport(t core.Transport) Option { return func(bot *Bot) { bot.transport = t } }

// WithFMS sets a custom FSM (finite state machine) implementation.
func WithFMS(f core.FMS) Option { return func(bot *Bot) { bot.fms = f } }

// WithLogger sets a custom Logger implementation.
func WithLogger(l core.Logger) Option { return func(bot *Bot) { bot.log = l } }
