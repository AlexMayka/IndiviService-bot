package giveaway

type UniqueGiftInfo struct {
	Gift              string `json:"gift"`
	Origin            string `json:"origin"`
	OwnedGiftId       string `json:"owned_gift_id,omitempty"`
	TransferStarCount string `json:"transfer_star_count,omitempty"`
}
