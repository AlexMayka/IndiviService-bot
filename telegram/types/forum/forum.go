package forum

type ForumTopicEdited struct {
	Name              string `json:"name,omitempty"`
	IconCustomEmojiId string `json:"icon_custom_emoji_id,omitempty"`
}

type ForumTopicCreated struct {
	Name              string `json:"name"`
	IconColor         int    `json:"icon_color"`
	IconCustomEmojiId string `json:"icon_custom_emoji_id,omitempty"`
}

type ForumTopicClosed struct{}

type ForumTopicReopened struct{}

type GeneralForumTopicHidden struct{}

type GeneralForumTopicUnhidden struct{}
