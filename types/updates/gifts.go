package updates

type Gift struct {
	Id               string  `json:"id"`
	Sticker          Sticker `json:"sticker"`
	StarCount        int     `json:"star_count"`
	UpgradeStarCount int     `json:"upgrade_star_count,omitempty"`
	TotalCount       int     `json:"total_count,omitempty"`
	RemainCount      int     `json:"remain_count,omitempty"`
}

type GiveAway struct {
	Chats                         []*Chat  `json:"chats"`
	WinnersSelectionDate          int      `json:"winners_selection_date"`
	WinnerCount                   int      `json:"winner_count"`
	OnlyNewMembers                bool     `json:"only_new_members,omitempty"`
	HasPublicWinners              bool     `json:"has_public_winners,omitempty"`
	PrizeDescription              string   `json:"prize_description,omitempty"`
	CountryCodes                  []string `json:"country_codes,omitempty"`
	PrizeStarCount                int      `json:"prize_star_count,omitempty"`
	PremiumSubscriptionMonthCount int      `json:"premium_subscription_month_count,omitempty"`
}

type GiveawayWinners struct {
	Chat                          *Chat  `json:"chat"`
	GiveawayMessageID             int    `json:"giveaway_message_id"`
	WinnersSelectionDate          int    `json:"winners_selection_date"`
	WinnerCount                   int    `json:"winner_count"`
	Winners                       *User  `json:"winners"`
	AdditionalChatCount           int    `json:"additional_chat_count,omitempty"`
	PrizeStarCount                int    `json:"prize_star_count,omitempty"`
	PremiumSubscriptionMonthCount int    `json:"premium_subscription_month_count,omitempty"`
	UnclaimedPrizeCount           int    `json:"unclaimed_prize_count,omitempty"`
	OnlyNewMembers                bool   `json:"only_new_members,omitempty"`
	WasRefunded                   bool   `json:"was_refunded,omitempty"`
	PrizeDescription              string `json:"prize_description,omitempty"`
}

type GiveawayCreated struct {
	PrizeStarCount int `json:"prize_star_count,omitempty"`
}

type GiveawayCompleted struct {
	WinnerCount         int      `json:"winner_count"`
	UnclaimedPrizeCount int      `json:"unclaimed_prize_count,omitempty"`
	GiveawayMessage     *Message `json:"giveaway_message,omitempty"`
	IsStarGiveaway      bool     `json:"is_star_giveaway,omitempty"`
}

type GiftInfo struct {
	Gift                    Gift             `json:"gift"`
	OwnedStarId             string           `json:"owned_star_id,omitempty"`
	ConvertStarCount        string           `json:"convert_star_count,omitempty"`
	PrepaidUpgradeStarCount int64            `json:"prepaid_upgrade_star_count,omitempty"`
	CanBeUpgraded           bool             `json:"can_be_upgraded,omitempty"`
	Text                    string           `json:"text,omitempty"`
	Entities                []*MessageEntity `json:"entities,omitempty"`
	IsPrivate               bool             `json:"is_private,omitempty"`
}
