package commands

import "telegram-sdk/types/updates"

// KeyboardBuilder provides a fluent interface for building inline keyboards
type KeyboardBuilder struct {
	rows [][]*updates.InlineKeyboardButton
}

// NewKeyboard creates a new keyboard builder
func NewKeyboard() *KeyboardBuilder {
	return &KeyboardBuilder{
		rows: make([][]*updates.InlineKeyboardButton, 0),
	}
}

// Row adds a new row of buttons to the keyboard
func (kb *KeyboardBuilder) Row(buttons ...*updates.InlineKeyboardButton) *KeyboardBuilder {
	kb.rows = append(kb.rows, buttons)
	return kb
}

// Button creates an inline keyboard button with text and callback data
func Button(text, callbackData string) *updates.InlineKeyboardButton {
	return &updates.InlineKeyboardButton{
		Text:         text,
		CallbackData: callbackData,
	}
}

// URLButton creates an inline keyboard button with text and URL
func URLButton(text, url string) *updates.InlineKeyboardButton {
	return &updates.InlineKeyboardButton{
		Text: text,
		Url:  url,
	}
}

// SwitchInlineButton creates an inline keyboard button that switches to inline mode
func SwitchInlineButton(text, query string) *updates.InlineKeyboardButton {
	return &updates.InlineKeyboardButton{
		Text:              text,
		SwitchInlineQuery: query,
	}
}

// Build creates the final InlineKeyboardMarkup
func (kb *KeyboardBuilder) Build() *updates.InlineKeyboardMarkup {
	return &updates.InlineKeyboardMarkup{
		InlineKeyboard: kb.rows,
	}
}