package updates

type InlineKeyboardButton struct {
	Text                         string                       `json:"text"`
	Url                          string                       `json:"url,omitempty"`
	CallbackData                 string                       `json:"callback_data,omitempty"`
	WebApp                       *WebAppInfo                  `json:"web_app,omitempty"`
	LoginUrl                     *LoginUrl                    `json:"login_url,omitempty"`
	SwitchInlineQuery            string                       `json:"switch_inline_query,omitempty"`
	SwitchInlineQueryCurrentChat string                       `json:"switch_inline_query_current_chat,omitempty"`
	SwitchInlineQueryChosenChat  *SwitchInlineQueryChosenChat `json:"switch_inline_query_chosen_chat,omitempty"`
	CopyText                     *CopyTextButton              `json:"copy_text,omitempty"`
	CallbackGame                 *CallbackGame                `json:"callback_game,omitempty"`
	Pay                          bool                         `json:"pay,omitempty"`
}

type LoginUrl struct {
	Url                string `json:"url"`
	ForwardText        string `json:"forward_text,omitempty"`
	BotUsername        string `json:"bot_username,omitempty"`
	RequestWriteAccess bool   `json:"request_write_access,omitempty"`
}

type SwitchInlineQueryChosenChat struct {
	Query             string `json:"query,omitempty"`
	AllowUserChats    bool   `json:"allow_user_chats,omitempty"`
	AllowBotChats     bool   `json:"allow_bot_chats,omitempty"`
	AllowGroupChats   bool   `json:"allow_group_chats,omitempty"`
	AllowChannelChats bool   `json:"allow_channel_chats,omitempty"`
}

type CopyTextButton struct {
	Text string `json:"text"`
}

type CallbackGame struct{}
