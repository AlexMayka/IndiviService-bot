package giveaway

import "telegram-sdk/telegram/types/media/attachments"

type Gift struct {
	Id               string              `json:"id"`
	Sticker          attachments.Sticker `json:"sticker"`
	StarCount        int                 `json:"star_count"`
	UpgradeStarCount int                 `json:"upgrade_star_count,omitempty"`
	TotalCount       int                 `json:"total_count,omitempty"`
	RemainCount      int                 `json:"remain_count,omitempty"`
}
