package background

import "telegram-sdk/telegram/types/media/attachments"

type BackgroundType interface {
	getType() string
}

type BackgroundTypeFill struct {
	Type             string         `json:"type"`
	Fill             BackgroundFill `json:"fill"`
	DarkThemeDimming int            `json:"dark_theme_dimming"`
}

func (back *BackgroundTypeFill) getType() string {
	return back.Type
}

type BackgroundTypeWallpaper struct {
	Type             string                `json:"type"`
	Document         *attachments.Document `json:"document"`
	DarkThemeDimming int                   `json:"dark_theme_dimming"`
	IsBlurred        bool                  `json:"is_blurred,omitempty"`
	IsMoving         bool                  `json:"is_moving,omitempty"`
}

func (back *BackgroundTypeWallpaper) getType() string {
	return back.Type
}

type BackgroundTypePattern struct {
	Type      string                `json:"type"`
	Document  *attachments.Document `json:"document"`
	Fill      BackgroundFill        `json:"fill"`
	Intensity int                   `json:"intensity"`
	IsBlurred bool                  `json:"is_blurred,omitempty"`
	IsMoving  bool                  `json:"is_moving,omitempty"`
}

func (back *BackgroundTypePattern) getType() string {
	return back.Type
}

type BackgroundTypeChatTheme struct {
	Type      string `json:"type"`
	ThemeName string `json:"theme_name"`
}
