package giveaway

import "telegram-sdk/telegram/types/base"

type GiveawayWinners struct {
	Chat                          *base.Chat `json:"chat"`
	GiveawayMessageID             int        `json:"giveaway_message_id"`
	WinnersSelectionDate          int        `json:"winners_selection_date"`
	WinnerCount                   int        `json:"winner_count"`
	Winners                       *base.User `json:"winners"`
	AdditionalChatCount           int        `json:"additional_chat_count,omitempty"`
	PrizeStarCount                int        `json:"prize_star_count,omitempty"`
	PremiumSubscriptionMonthCount int        `json:"premium_subscription_month_count,omitempty"`
	UnclaimedPrizeCount           int        `json:"unclaimed_prize_count,omitempty"`
	OnlyNewMembers                bool       `json:"only_new_members,omitempty"`
	WasRefunded                   bool       `json:"was_refunded,omitempty"`
	PrizeDescription              string     `json:"prize_description,omitempty"`
}
