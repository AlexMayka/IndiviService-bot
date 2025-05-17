package video_chat

import "telegram-sdk/telegram/types/base"

type VideoChatParticipantsInvited struct {
	Users []*base.User `json:"users"`
}
