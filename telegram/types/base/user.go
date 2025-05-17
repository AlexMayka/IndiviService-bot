package base

import (
	"telegram-sdk/telegram/types/media/attachments"
)

type User struct {
	Id                      int    `json:"id"`
	IsBot                   bool   `json:"is_bot"`
	FirstName               string `json:"first_name"`
	LastName                string `json:"last_name,omitempty"`
	Username                string `json:"username,omitempty"`
	LanguageCode            string `json:"language_code,omitempty"`
	IsPremium               bool   `json:"is_premium,omitempty"`
	AddedToAttachmentMenu   bool   `json:"added_to_attachment_menu,omitempty"`
	CanJoinGroups           bool   `json:"can_join_groups,omitempty"`
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages,omitempty"`
	SupportsInlineQueries   bool   `json:"supports_inline_queries,omitempty"`
	CanConnectToBusiness    bool   `json:"can_connect_to_business,omitempty"`
	HasMainWebApp           bool   `json:"has_main_web_app,omitempty"`
}

type UserShared struct {
	RequestId int           `json:"request_id"`
	Users     []*SharedUser `json:"users"`
}

type SharedUser struct {
	UserId    int                      `json:"user_id"`
	FirstName string                   `json:"first_name,omitempty"`
	LastName  string                   `json:"last_name,omitempty"`
	Username  string                   `json:"username,omitempty"`
	Photo     []*attachments.PhotoSize `json:"photo,omitempty"`
}
