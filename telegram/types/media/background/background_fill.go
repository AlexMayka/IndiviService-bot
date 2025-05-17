package background

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
