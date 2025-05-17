package base

import (
	"telegram-sdk/telegram/types/forum"
	"telegram-sdk/telegram/types/interactive"
	"telegram-sdk/telegram/types/media/attachments"
	"telegram-sdk/telegram/types/media/giveaway"
	"telegram-sdk/telegram/types/media/others"
	"telegram-sdk/telegram/types/media/passport"
	"telegram-sdk/telegram/types/media/video_chat"
	"telegram-sdk/telegram/types/media/webapp"
	"telegram-sdk/telegram/types/payment"
	"telegram-sdk/telegram/types/preview"
	"telegram-sdk/telegram/types/reply"
)

type Message struct {
	MessageID                     int                                      `json:"message_id"`
	MessageThreadId               int                                      `json:"message_thread_id,omitempty"`
	From                          *User                                    `json:"from,omitempty"`
	SenderChat                    *Chat                                    `json:"sender_chat,omitempty"`
	SenderBusinessBot             *User                                    `json:"sender_business_bot,omitempty"`
	Date                          int                                      `json:"date,omitempty"`
	BusinessConnectionId          string                                   `json:"business_connection_id,omitempty"`
	Chat                          *Chat                                    `json:"chat,omitempty"`
	ForwardOrigin                 *MessageOrigin                           `json:"forward_origin,omitempty"`
	IsTopicMessage                bool                                     `json:"is_topic_message,omitempty"`
	IsAutomaticForward            bool                                     `json:"is_automatic_forward,omitempty"`
	ReplyToMessage                *Message                                 `json:"reply_to_message,omitempty"`
	ExternalReply                 *reply.ExternalReplyInfo                 `json:"external_reply,omitempty"`
	Quote                         *others.TextQuote                        `json:"quote,omitempty"`
	ReplyToStory                  *others.Story                            `json:"reply_to_story,omitempty"`
	ViaBot                        *User                                    `json:"via_bot,omitempty"`
	EditDate                      int                                      `json:"edit_date,omitempty"`
	HasProtectedContent           bool                                     `json:"has_protected_content,omitempty"`
	IsFromOffline                 bool                                     `json:"is_from_offline,omitempty"`
	MediaGroupId                  string                                   `json:"media_group_id,omitempty"`
	AuthorSignature               string                                   `json:"author_signature,omitempty"`
	PaidStarCount                 int                                      `json:"paid_star_count,omitempty"`
	Text                          string                                   `json:"text,omitempty"`
	Entities                      *[]MessageEntity                         `json:"entities,omitempty"`
	LinkPreviewOptions            *preview.LinkPreviewOptions              `json:"link_preview_options,omitempty"`
	EffectId                      string                                   `json:"effect_id,omitempty"`
	Animation                     *attachments.Animation                   `json:"animation,omitempty"`
	Audio                         *attachments.Audio                       `json:"audio,omitempty"`
	Document                      *attachments.Document                    `json:"document,omitempty"`
	PaidMedia                     *others.PaidMedia                        `json:"paid_media,omitempty"`
	Photo                         []*attachments.PhotoSize                 `json:"photo,omitempty"`
	Sticker                       *attachments.Sticker                     `json:"sticker,omitempty"`
	Story                         *others.Story                            `json:"story,omitempty"`
	Video                         *attachments.Video                       `json:"video,omitempty"`
	VideoNote                     *attachments.VideoNote                   `json:"video_note,omitempty"`
	Voice                         *attachments.Voice                       `json:"voice,omitempty"`
	Caption                       string                                   `json:"caption,omitempty"`
	CaptionEntities               *[]MessageEntity                         `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia         bool                                     `json:"show_caption_above_media,omitempty"`
	HasMediaSpoiler               bool                                     `json:"has_media_spoiler,omitempty"`
	Contact                       *interactive.Contact                     `json:"contact,omitempty"`
	Dice                          *interactive.Dice                        `json:"dice,omitempty"`
	Game                          *interactive.Game                        `json:"game,omitempty"`
	Poll                          *others.Poll                             `json:"poll,omitempty"`
	Venue                         *others.Venue                            `json:"venue,omitempty"`
	Location                      *others.Location                         `json:"location,omitempty"`
	NewChatMembers                []*User                                  `json:"new_chat_members,omitempty"`
	LeftChatMember                *User                                    `json:"left_chat_member,omitempty"`
	NewChatTitle                  string                                   `json:"new_chat_title,omitempty"`
	NewChatPhoto                  []*attachments.PhotoSize                 `json:"new_chat_photo,omitempty"`
	DeleteChatPhoto               bool                                     `json:"delete_chat_photo,omitempty"`
	GroupChatCreated              bool                                     `json:"group_chat_created,omitempty"`
	SupergroupChatCreated         bool                                     `json:"supergroup_chat_created,omitempty"`
	ChannelChatCreated            bool                                     `json:"channel_chat_created,omitempty"`
	MessageAutoDeleteTimerChanged *MessageAutoDeleteTimerChanged           `json:"message_auto_delete_timer_changed,omitempty"`
	MigrateToChatId               int                                      `json:"migrate_to_chat_id,omitempty"`
	MigrateFromChatId             int                                      `json:"migrate_from_chat_id,omitempty"`
	PinnedMessage                 *MaybeInaccessibleMessage                `json:"pinned_message,omitempty"`
	Invoice                       *payment.Invoice                         `json:"invoice,omitempty"`
	SuccessfulPayment             *payment.SuccessfulPayment               `json:"successful_payment,omitempty"`
	RefundedPayment               *payment.RefundedPayment                 `json:"refunded_payment,omitempty"`
	UserShared                    *UserShared                              `json:"user_shared,omitempty"`
	Gift                          *giveaway.GiftInfo                       `json:"gift,omitempty"`
	UniqueGift                    *giveaway.UniqueGiftInfo                 `json:"unique_gift,omitempty"`
	ConnectedWebsite              string                                   `json:"connected_website,omitempty"`
	PassportData                  *passport.PassportData                   `json:"passport_data,omitempty"`
	ProximityAlertTriggered       *others.ProximityAlertTriggered          `json:"proximity_alert_triggered,omitempty"`
	BoostAdded                    *ChatBoostAdded                          `json:"boost_added,omitempty"`
	ChatBackgroundSet             *ChatBackground                          `json:"chat_background_set,omitempty"`
	ForumTopicCreated             *forum.ForumTopicCreated                 `json:"forum_topic_created,omitempty"`
	ForumTopicEdited              *forum.ForumTopicEdited                  `json:"forum_topic_edited,omitempty"`
	ForumTopicClosed              *forum.ForumTopicClosed                  `json:"forum_topic_closed,omitempty"`
	ForumTopicReopened            *forum.ForumTopicReopened                `json:"forum_topic_reopened,omitempty"`
	GeneralForumTopicHidden       *forum.GeneralForumTopicHidden           `json:"general_forum_topic_hidden,omitempty"`
	GeneralForumTopicUnhidden     *forum.GeneralForumTopicUnhidden         `json:"general_forum_topic_unhidden,omitempty"`
	GiveawayCreated               *giveaway.GiveawayCreated                `json:"giveaway_created,omitempty"`
	Giveaway                      *giveaway.GiveAway                       `json:"giveaway,omitempty"`
	GiveawayWinners               *giveaway.GiveawayWinners                `json:"giveaway_winners,omitempty"`
	GiveawayCompleted             *giveaway.GiveawayCompleted              `json:"giveaway_completed,omitempty"`
	PaidMessagePriceChanged       *PaidMessagePriceChanged                 `json:"paid_message_price_changed,omitempty"`
	VideoChatScheduled            *video_chat.VideoChatScheduled           `json:"video_chat_scheduled,omitempty"`
	VideoChatStarted              *video_chat.VideoChatStarted             `json:"video_chat_started,omitempty"`
	VideoChatEnded                *video_chat.VideoChatEnded               `json:"video_chat_ended,omitempty"`
	VideoChatParticipantsInvited  *video_chat.VideoChatParticipantsInvited `json:"video_chat_participants_invited,omitempty"`
	WebAppData                    *webapp.WebAppData                       `json:"web_app_data,omitempty"`
	ReplyMarkup                   *interactive.InlineKeyboardMarkup        `json:"reply_markup,omitempty"`
}

type MessageOrigin interface {
	GetType() string
}

type MessageOriginUser struct {
	Type       string `json:"type"`
	Date       int    `json:"date"`
	SenderUser *User  `json:"sender_user"`
}

func (mes MessageOriginUser) GetType() string {
	return mes.Type
}

type MessageOriginHiddenUser struct {
	Type           string `json:"type"`
	Date           int    `json:"date"`
	SenderUserName string `json:"sender_user_name"`
}

func (mes MessageOriginHiddenUser) GetType() string {
	return mes.Type
}

type MessageOriginChat struct {
	Type            string `json:"type"`
	Date            int    `json:"date"`
	SendChat        *Chat  `json:"send_chat"`
	AuthorSignature string `json:"author_signature,omitempty"`
}

func (mes MessageOriginChat) GetType() string {
	return mes.Type
}

type MessageOriginChannel struct {
	Type            string `json:"type"`
	Date            int    `json:"date"`
	ChatID          int    `json:"chat"`
	MessageID       int    `json:"message_id"`
	AuthorSignature string `json:"author_signature,omitempty"`
}

func (mes MessageOriginChannel) GetType() string {
	return mes.Type
}

type MessageEntity struct {
	Type          string `json:"type"`
	Offset        int    `json:"offset"`
	Length        int    `json:"length"`
	Url           string `json:"url,omitempty"`
	User          *User  `json:"user,omitempty"`
	Language      string `json:"language,omitempty"`
	CustomEmojiId string `json:"custom_emoji_id,omitempty"`
}

type MessageAutoDeleteTimerChanged struct {
	MessageAutoDeleteTime int `json:"message_auto_delete_time"`
}

type MaybeInaccessibleMessage struct {
	Type          string `json:"type"`
	Offset        int    `json:"offset"`
	Length        int    `json:"length"`
	Url           string `json:"url,omitempty"`
	User          *User  `json:"user,omitempty"`
	Language      string `json:"language,omitempty"`
	CustomEmojiId string `json:"custom_emoji_id,omitempty"`
}

type PaidMessagePriceChanged struct {
	PaidMessageStarCount int `json:"paid_message_star_count"`
}
