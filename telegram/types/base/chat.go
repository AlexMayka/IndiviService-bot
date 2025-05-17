package base

import (
	"telegram-sdk/telegram/types/media/background"
)

type Chat struct {
	Id        int    `json:"id"`
	Type      string `json:"type"`
	Title     string `json:"title,omitempty"`
	Username  string `json:"username,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	IsForum   bool   `json:"is_forum,omitempty"`
}

type ChatBoostAdded struct {
	BoostCount int `json:"boost_count"`
}

type ChatBackground struct {
	Type *background.BackgroundType `json:"type"`
}
