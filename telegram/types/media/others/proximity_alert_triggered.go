package others

import "telegram-sdk/telegram/types/base"

type ProximityAlertTriggered struct {
	Traveler base.User `json:"traveler"`
	Watcher  base.User `json:"watcher"`
	Distance int       `json:"distance"`
}
