package others

import "telegram-sdk/telegram/types/media/attachments"

type PaidMediaInfo struct {
	StarCount int         `json:"star_count"`
	PaidMedia []PaidMedia `json:"paid_media"`
}

type PaidMedia interface {
	GetType() string
}

type PaidMediaPreview struct {
	Type     string `json:"type"`
	Width    int    `json:"width,omitempty"`
	Height   int    `json:"height,omitempty"`
	Duration int    `json:"duration,omitempty"`
}

func (paid *PaidMediaPreview) GetType() string {
	return paid.Type
}

type PaidMediaPhoto struct {
	Type  string                  `json:"type"`
	Photo []attachments.PhotoSize `json:"photo"`
}

func (paid *PaidMediaPhoto) GetType() string {
	return paid.Type
}

type PaidMediaVideo struct {
	Type  string            `json:"type"`
	Video attachments.Video `json:"video"`
}

func (paid *PaidMediaVideo) GetType() string {
	return paid.Type
}
