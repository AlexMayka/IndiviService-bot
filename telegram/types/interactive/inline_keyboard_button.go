package interactive

import (
	"telegram-sdk/telegram/types/media/others"
	"telegram-sdk/telegram/types/media/webapp"
)

type InlineKeyboardButton struct {
	Text                         string                       `json:"text"`
	Url                          string                       `json:"url,omitempty"`
	CallbackData                 string                       `json:"callback_data,omitempty"`
	WebApp                       *webapp.WebAppInfo           `json:"web_app,omitempty"`
	LoginUrl                     *webapp.LoginUrl             `json:"login_url,omitempty"`
	SwitchInlineQuery            string                       `json:"switch_inline_query,omitempty"`
	SwitchInlineQueryCurrentChat string                       `json:"switch_inline_query_current_chat,omitempty"`
	SwitchInlineQueryChosenChat  *SwitchInlineQueryChosenChat `json:"switch_inline_query_chosen_chat,omitempty"`
	CopyText                     *others.CopyTextButton       `json:"copy_text,omitempty"`
	CallbackGame                 *CallbackGame                `json:"callback_game,omitempty"`
	Pay                          bool                         `json:"pay,omitempty"`
}
