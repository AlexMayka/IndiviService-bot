package attachments

import (
	"telegram-sdk/telegram/types/media/others"
)

type Sticker struct {
	FileID           string               `json:"file_id"`
	FileUniqueID     string               `json:"file_unique_id"`
	Type             string               `json:"type"`
	Width            int                  `json:"width"`
	Height           int                  `json:"height"`
	IsAnimated       bool                 `json:"is_animated"`
	IsVideo          bool                 `json:"is_video"`
	Thumbnail        *PhotoSize           `json:"thumbnail,omitempty"`
	Emoji            string               `json:"emoji,omitempty"`
	SetName          string               `json:"set_name,omitempty"`
	PremiumAnimation *File                `json:"premium_animation,omitempty"`
	MaskPosition     *others.MaskPosition `json:"mask_position,omitempty"`
	CustomEmojiID    string               `json:"custom_emoji_id,omitempty"`
	NeedsRepainting  bool                 `json:"needs_repainting,omitempty"`
	FileSize         int64                `json:"file_size,omitempty"`
}
