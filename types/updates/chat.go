package updates

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

type MessageAutoDeleteTimerChanged struct {
	MessageAutoDeleteTime int `json:"message_auto_delete_time"`
}

type ChatBackground struct {
	Type *BackgroundType `json:"type"`
}

type BackgroundFill interface {
	GetType() string
}

type BackgroundFillSolid struct {
	Type  string `json:"type"`
	Color int    `json:"color"`
}

func (back *BackgroundFillSolid) GetType() string {
	return back.Type
}

type BackgroundFillGradient struct {
	Type          string `json:"type"`
	TopColor      int    `json:"top_color"`
	BottomColor   int    `json:"bottom_color"`
	RotationAngle int    `json:"rotation_angle"`
}

func (back *BackgroundFillGradient) GetType() string {
	return back.Type
}

type BackgroundFillFreeformGradient struct {
	Type  string `json:"type"`
	Color []int  `json:"color"`
}

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
	Type             string    `json:"type"`
	Document         *Document `json:"document"`
	DarkThemeDimming int       `json:"dark_theme_dimming"`
	IsBlurred        bool      `json:"is_blurred,omitempty"`
	IsMoving         bool      `json:"is_moving,omitempty"`
}

func (back *BackgroundTypeWallpaper) getType() string {
	return back.Type
}

type BackgroundTypePattern struct {
	Type      string         `json:"type"`
	Document  *Document      `json:"document"`
	Fill      BackgroundFill `json:"fill"`
	Intensity int            `json:"intensity"`
	IsBlurred bool           `json:"is_blurred,omitempty"`
	IsMoving  bool           `json:"is_moving,omitempty"`
}

func (back *BackgroundTypePattern) getType() string {
	return back.Type
}

type BackgroundTypeChatTheme struct {
	Type      string `json:"type"`
	ThemeName string `json:"theme_name"`
}
