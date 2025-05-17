package giveaway

import "telegram-sdk/telegram/types/base"

type GiftInfo struct {
	Gift                    Gift                  `json:"gift"`
	OwnedStarId             string                `json:"owned_star_id,omitempty"`
	ConvertStarCount        string                `json:"convert_star_count,omitempty"`
	PrepaidUpgradeStarCount int64                 `json:"prepaid_upgrade_star_count,omitempty"`
	CanBeUpgraded           bool                  `json:"can_be_upgraded,omitempty"`
	Text                    string                `json:"text,omitempty"`
	Entities                []*base.MessageEntity `json:"entities,omitempty"`
	IsPrivate               bool                  `json:"is_private,omitempty"`
}
