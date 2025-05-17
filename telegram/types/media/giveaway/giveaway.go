package giveaway

import "telegram-sdk/telegram/types/base"

type GiveAway struct {
	Chats                         []*base.Chat `json:"chats"`
	WinnersSelectionDate          int          `json:"winners_selection_date"`
	WinnerCount                   int          `json:"winner_count"`
	OnlyNewMembers                bool         `json:"only_new_members,omitempty"`
	HasPublicWinners              bool         `json:"has_public_winners,omitempty"`
	PrizeDescription              string       `json:"prize_description,omitempty"`
	CountryCodes                  []string     `json:"country_codes,omitempty"`
	PrizeStarCount                int          `json:"prize_star_count,omitempty"`
	PremiumSubscriptionMonthCount int          `json:"premium_subscription_month_count,omitempty"`
}
