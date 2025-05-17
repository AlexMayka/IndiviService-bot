package giveaway

import "telegram-sdk/telegram/types/base"

type GiveawayCompleted struct {
	WinnerCount         int           `json:"winner_count"`
	UnclaimedPrizeCount int           `json:"unclaimed_prize_count,omitempty"`
	GiveawayMessage     *base.Message `json:"giveaway_message,omitempty"`
	IsStarGiveaway      bool          `json:"is_star_giveaway,omitempty"`
}
