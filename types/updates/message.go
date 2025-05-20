package updates

type Message struct {
	MessageID                     int                            `json:"message_id"`
	MessageThreadId               int                            `json:"message_thread_id,omitempty"`
	From                          *User                          `json:"from,omitempty"`
	SenderChat                    *Chat                          `json:"sender_chat,omitempty"`
	SenderBusinessBot             *User                          `json:"sender_business_bot,omitempty"`
	Date                          int                            `json:"date,omitempty"`
	BusinessConnectionId          string                         `json:"business_connection_id,omitempty"`
	Chat                          *Chat                          `json:"chat,omitempty"`
	ForwardOrigin                 *MessageOrigin                 `json:"forward_origin,omitempty"`
	IsTopicMessage                bool                           `json:"is_topic_message,omitempty"`
	IsAutomaticForward            bool                           `json:"is_automatic_forward,omitempty"`
	ReplyToMessage                *Message                       `json:"reply_to_message,omitempty"`
	ExternalReply                 *ExternalReplyInfo             `json:"external_reply,omitempty"`
	Quote                         *TextQuote                     `json:"quote,omitempty"`
	ReplyToStory                  *Story                         `json:"reply_to_story,omitempty"`
	ViaBot                        *User                          `json:"via_bot,omitempty"`
	EditDate                      int                            `json:"edit_date,omitempty"`
	HasProtectedContent           bool                           `json:"has_protected_content,omitempty"`
	IsFromOffline                 bool                           `json:"is_from_offline,omitempty"`
	MediaGroupId                  string                         `json:"media_group_id,omitempty"`
	AuthorSignature               string                         `json:"author_signature,omitempty"`
	PaidStarCount                 int                            `json:"paid_star_count,omitempty"`
	Text                          string                         `json:"text,omitempty"`
	Entities                      *[]MessageEntity               `json:"entities,omitempty"`
	LinkPreviewOptions            *LinkPreviewOptions            `json:"link_preview_options,omitempty"`
	EffectId                      string                         `json:"effect_id,omitempty"`
	Animation                     *Animation                     `json:"animation,omitempty"`
	Audio                         *Audio                         `json:"audio,omitempty"`
	Document                      *Document                      `json:"document,omitempty"`
	PaidMedia                     *PaidMedia                     `json:"paid_media,omitempty"`
	Photo                         []*PhotoSize                   `json:"photo,omitempty"`
	Sticker                       *Sticker                       `json:"sticker,omitempty"`
	Story                         *Story                         `json:"story,omitempty"`
	Video                         *Video                         `json:"video,omitempty"`
	VideoNote                     *VideoNote                     `json:"video_note,omitempty"`
	Voice                         *Voice                         `json:"voice,omitempty"`
	Caption                       string                         `json:"caption,omitempty"`
	CaptionEntities               *[]MessageEntity               `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia         bool                           `json:"show_caption_above_media,omitempty"`
	HasMediaSpoiler               bool                           `json:"has_media_spoiler,omitempty"`
	Contact                       *Contact                       `json:"contact,omitempty"`
	Dice                          *Dice                          `json:"dice,omitempty"`
	Game                          *Game                          `json:"game,omitempty"`
	Poll                          *Poll                          `json:"poll,omitempty"`
	Venue                         *Venue                         `json:"venue,omitempty"`
	Location                      *Location                      `json:"location,omitempty"`
	NewChatMembers                []*User                        `json:"new_chat_members,omitempty"`
	LeftChatMember                *User                          `json:"left_chat_member,omitempty"`
	NewChatTitle                  string                         `json:"new_chat_title,omitempty"`
	NewChatPhoto                  []*PhotoSize                   `json:"new_chat_photo,omitempty"`
	DeleteChatPhoto               bool                           `json:"delete_chat_photo,omitempty"`
	GroupChatCreated              bool                           `json:"group_chat_created,omitempty"`
	SupergroupChatCreated         bool                           `json:"supergroup_chat_created,omitempty"`
	ChannelChatCreated            bool                           `json:"channel_chat_created,omitempty"`
	MessageAutoDeleteTimerChanged *MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed,omitempty"`
	MigrateToChatId               int                            `json:"migrate_to_chat_id,omitempty"`
	MigrateFromChatId             int                            `json:"migrate_from_chat_id,omitempty"`
	PinnedMessage                 *MaybeInaccessibleMessage      `json:"pinned_message,omitempty"`
	Invoice                       *Invoice                       `json:"invoice,omitempty"`
	SuccessfulPayment             *SuccessfulPayment             `json:"successful_payment,omitempty"`
	RefundedPayment               *RefundedPayment               `json:"refunded_payment,omitempty"`
	UserShared                    *UserShared                    `json:"user_shared,omitempty"`
	Gift                          *GiftInfo                      `json:"gift,omitempty"`
	UniqueGift                    *UniqueGiftInfo                `json:"unique_gift,omitempty"`
	ConnectedWebsite              string                         `json:"connected_website,omitempty"`
	PassportData                  *PassportData                  `json:"passport_data,omitempty"`
	ProximityAlertTriggered       *ProximityAlertTriggered       `json:"proximity_alert_triggered,omitempty"`
	BoostAdded                    *ChatBoostAdded                `json:"boost_added,omitempty"`
	ChatBackgroundSet             *ChatBackground                `json:"chat_background_set,omitempty"`
	ForumTopicCreated             *ForumTopicCreated             `json:"forum_topic_created,omitempty"`
	ForumTopicEdited              *ForumTopicEdited              `json:"forum_topic_edited,omitempty"`
	ForumTopicClosed              *ForumTopicClosed              `json:"forum_topic_closed,omitempty"`
	ForumTopicReopened            *ForumTopicReopened            `json:"forum_topic_reopened,omitempty"`
	GeneralForumTopicHidden       *GeneralForumTopicHidden       `json:"general_forum_topic_hidden,omitempty"`
	GeneralForumTopicUnhidden     *GeneralForumTopicUnhidden     `json:"general_forum_topic_unhidden,omitempty"`
	GiveawayCreated               *GiveawayCreated               `json:"giveaway_created,omitempty"`
	Giveaway                      *GiveAway                      `json:"giveaway,omitempty"`
	GiveawayWinners               *GiveawayWinners               `json:"giveaway_winners,omitempty"`
	GiveawayCompleted             *GiveawayCompleted             `json:"giveaway_completed,omitempty"`
	PaidMessagePriceChanged       *PaidMessagePriceChanged       `json:"paid_message_price_changed,omitempty"`
	VideoChatScheduled            *VideoChatScheduled            `json:"video_chat_scheduled,omitempty"`
	VideoChatStarted              *VideoChatStarted              `json:"video_chat_started,omitempty"`
	VideoChatEnded                *VideoChatEnded                `json:"video_chat_ended,omitempty"`
	VideoChatParticipantsInvited  *VideoChatParticipantsInvited  `json:"video_chat_participants_invited,omitempty"`
	WebAppData                    *WebAppData                    `json:"web_app_data,omitempty"`
	ReplyMarkup                   *InlineKeyboardMarkup          `json:"reply_markup,omitempty"`
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

type TextQuote struct {
	Text     string           `json:"text"`
	Entities []*MessageEntity `json:"entities,omitempty"`
	Position int              `json:"position"`
	IsManual bool             `json:"is_manual,omitempty"`
}

type ExternalReplyInfo struct {
	Origin             *MessageOrigin      `json:"origin"`
	Chat               *Chat               `json:"chat,omitempty"`
	MessageID          int                 `json:"message_id,omitempty"`
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"`
	Animation          *Animation          `json:"animation,omitempty"`
	Audio              *Audio              `json:"audio,omitempty"`
	Document           *Document           `json:"document,omitempty"`
	PaidMedia          *PaidMediaInfo      `json:"paid_media,omitempty"`
	Photo              *[]PhotoSize        `json:"photo,omitempty"`
	Sticker            *Sticker            `json:"sticker,omitempty"`
	Story              *Story              `json:"story,omitempty"`
	Video              *Video              `json:"video,omitempty"`
	VideoNote          *VideoNote          `json:"video_note,omitempty"`
	Voice              *Voice              `json:"voice,omitempty"`
	HasMediaSpoiler    bool                `json:"has_media_spoiler,omitempty"`
	Contact            *Contact            `json:"contact,omitempty"`
	Dice               *Dice               `json:"dice,omitempty"`
	Game               *Game               `json:"game,omitempty"`
	Giveaway           *GiveAway           `json:"giveaway,omitempty"`
	GiveawayWinners    *GiveawayWinners    `json:"giveaway_winners,omitempty"`
	Invoice            *Invoice            `json:"invoice,omitempty"`
	Location           *Location           `json:"location,omitempty"`
	Poll               *Poll               `json:"poll,omitempty"`
	Venue              *Venue              `json:"venue,omitempty"`
}

type LinkPreviewOptions struct {
	IsDisabled       bool   `json:"is_disabled,omitempty"`
	Url              string `json:"url,omitempty"`
	PreferSmallMedia bool   `json:"prefer_small_media,omitempty"`
	PreferLargeMedia bool   `json:"prefer_large_media,omitempty"`
	ShowAboveText    bool   `json:"show_above_text,omitempty"`
}

type UniqueGiftInfo struct {
	Gift              string `json:"gift"`
	Origin            string `json:"origin"`
	OwnedGiftId       string `json:"owned_gift_id,omitempty"`
	TransferStarCount string `json:"transfer_star_count,omitempty"`
}

type Story struct {
	Chat Chat `json:"chat"`
	Id   int  `json:"id"`
}

type Game struct {
	Title        string           `json:"title"`
	Description  string           `json:"description"`
	Photo        []*PhotoSize     `json:"photo"`
	Text         string           `json:"text"`
	TextEntities []*MessageEntity `json:"text_entities,omitempty"`
	Animation    *Animation       `json:"animation,omitempty"`
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]*InlineKeyboardButton `json:"inline_keyboard"`
}

type ProximityAlertTriggered struct {
	Traveler User `json:"traveler"`
	Watcher  User `json:"watcher"`
	Distance int  `json:"distance"`
}
