package updates

// Chat represents a Telegram chat: private, group, supergroup or channel.
type Chat struct {
	Id        int    `json:"id"`                   // Unique identifier for the chat
	Type      string `json:"type"`                 // Type of chat (e.g., "private", "group")
	Title     string `json:"title,omitempty"`      // Optional title for supergroups, channels, etc.
	Username  string `json:"username,omitempty"`   // Optional username of the chat
	FirstName string `json:"first_name,omitempty"` // Optional first name (for users)
	LastName  string `json:"last_name,omitempty"`  // Optional last name (for users)
	IsForum   bool   `json:"is_forum,omitempty"`   // True if this chat is a forum supergroup
}

// ChatBoostAdded indicates how many boosts were added to a chat.
type ChatBoostAdded struct {
	BoostCount int `json:"boost_count"` // Number of boosts added
}

// MessageAutoDeleteTimerChanged signals a change in the auto-delete timer.
type MessageAutoDeleteTimerChanged struct {
	MessageAutoDeleteTime int `json:"message_auto_delete_time"` // New auto-delete duration (in seconds)
}

// ChatBackground represents a chat background configuration.
type ChatBackground struct {
	Type *BackgroundType `json:"type"` // One of: fill, wallpaper, pattern, or theme
}

// BackgroundFill is the interface for various background fill styles.
type BackgroundFill interface {
	GetType() string
}

// BackgroundFillSolid represents a single-color background.
type BackgroundFillSolid struct {
	Type  string `json:"type"`  // Should be "fill"
	Color int    `json:"color"` // Hex RGB integer
}

func (b *BackgroundFillSolid) GetType() string {
	return b.Type
}

// BackgroundFillGradient represents a vertical gradient background.
type BackgroundFillGradient struct {
	Type          string `json:"type"`           // Should be "gradient"
	TopColor      int    `json:"top_color"`      // Top color (hex RGB)
	BottomColor   int    `json:"bottom_color"`   // Bottom color (hex RGB)
	RotationAngle int    `json:"rotation_angle"` // Degrees of rotation
}

func (b *BackgroundFillGradient) GetType() string {
	return b.Type
}

// BackgroundFillFreeformGradient represents a complex multi-point gradient.
type BackgroundFillFreeformGradient struct {
	Type  string `json:"type"`  // Should be "freeform_gradient"
	Color []int  `json:"color"` // Array of hex RGB color points
}

// BackgroundType is the interface for various background configurations.
type BackgroundType interface {
	getType() string
}

// BackgroundTypeFill uses a fill (solid or gradient) as the background.
type BackgroundTypeFill struct {
	Type             string         `json:"type"`               // Should be "fill"
	Fill             BackgroundFill `json:"fill"`               // Fill style
	DarkThemeDimming int            `json:"dark_theme_dimming"` // Darkness adjustment for dark theme
}

func (b *BackgroundTypeFill) getType() string {
	return b.Type
}

// BackgroundTypeWallpaper uses an image document as background.
type BackgroundTypeWallpaper struct {
	Type             string    `json:"type"`               // Should be "wallpaper"
	Document         *Document `json:"document"`           // Wallpaper file
	DarkThemeDimming int       `json:"dark_theme_dimming"` // Darkness adjustment
	IsBlurred        bool      `json:"is_blurred,omitempty"`
	IsMoving         bool      `json:"is_moving,omitempty"`
}

func (b *BackgroundTypeWallpaper) getType() string {
	return b.Type
}

// BackgroundTypePattern uses a pattern image + fill as background.
type BackgroundTypePattern struct {
	Type      string         `json:"type"`      // Should be "pattern"
	Document  *Document      `json:"document"`  // Pattern file
	Fill      BackgroundFill `json:"fill"`      // Fill under the pattern
	Intensity int            `json:"intensity"` // Pattern visibility
	IsBlurred bool           `json:"is_blurred,omitempty"`
	IsMoving  bool           `json:"is_moving,omitempty"`
}

func (b *BackgroundTypePattern) getType() string {
	return b.Type
}

// BackgroundTypeChatTheme refers to a predefined named theme.
type BackgroundTypeChatTheme struct {
	Type      string `json:"type"`       // Should be "chat_theme"
	ThemeName string `json:"theme_name"` // Identifier of the theme
}
