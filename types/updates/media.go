package updates

type Animation struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Width        int        `json:"width"`
	Height       int        `json:"height"`
	Duration     int        `json:"duration"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
	FileName     string     `json:"file_name,omitempty"`
	MimeType     string     `json:"mime_type,omitempty"`
	FileSize     int64      `json:"file_size,omitempty"`
}

type Audio struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Duration     int        `json:"duration"`
	Performer    string     `json:"performer,omitempty"`
	Title        string     `json:"title,omitempty"`
	FileName     string     `json:"file_name,omitempty"`
	MimeType     string     `json:"mime_type,omitempty"`
	FileSize     int64      `json:"file_size,omitempty"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
}

type Document struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
	FileName     string     `json:"file_name,omitempty"`
	MimeType     string     `json:"mime_type,omitempty"`
	FileSize     int64      `json:"file_size,omitempty"`
}

type Video struct {
	FileId         string      `json:"file_id"`
	FileUniqueId   string      `json:"file_unique_id"`
	Width          int         `json:"width"`
	Height         int         `json:"height"`
	Duration       int         `json:"duration"`
	Thumbnail      PhotoSize   `json:"thumbnail,omitempty"`
	Cover          []PhotoSize `json:"cover,omitempty"`
	StartTimestamp int64       `json:"start_timestamp,omitempty"`
	FileName       string      `json:"file_name,omitempty"`
	MimeType       string      `json:"mime_type,omitempty"`
	FileSize       int64       `json:"file_size,omitempty"`
}

type Voice struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	Duration     int    `json:"duration"`
	MimeType     string `json:"mime_type,omitempty"`
	FileSize     int    `json:"file_size,omitempty"`
}

type VideoNote struct {
	FileId       string    `json:"file_id"`
	FileUniqueId string    `json:"file_unique_id"`
	Length       int       `json:"length"`
	Duration     int       `json:"duration"`
	Thumbnail    PhotoSize `json:"thumbnail,omitempty"`
	FileSize     int       `json:"file_size,omitempty"`
}

type Sticker struct {
	FileID           string        `json:"file_id"`
	FileUniqueID     string        `json:"file_unique_id"`
	Type             string        `json:"type"`
	Width            int           `json:"width"`
	Height           int           `json:"height"`
	IsAnimated       bool          `json:"is_animated"`
	IsVideo          bool          `json:"is_video"`
	Thumbnail        *PhotoSize    `json:"thumbnail,omitempty"`
	Emoji            string        `json:"emoji,omitempty"`
	SetName          string        `json:"set_name,omitempty"`
	PremiumAnimation *File         `json:"premium_animation,omitempty"`
	MaskPosition     *MaskPosition `json:"mask_position,omitempty"`
	CustomEmojiID    string        `json:"custom_emoji_id,omitempty"`
	NeedsRepainting  bool          `json:"needs_repainting,omitempty"`
	FileSize         int64         `json:"file_size,omitempty"`
}

type PaidMediaInfo struct {
	StarCount int         `json:"star_count"`
	PaidMedia []PaidMedia `json:"paid_media"`
}

type PaidMedia interface {
	GetType() string
}

type PaidMediaPreview struct {
	Type     string `json:"type"`
	Width    int    `json:"width,omitempty"`
	Height   int    `json:"height,omitempty"`
	Duration int    `json:"duration,omitempty"`
}

func (paid *PaidMediaPreview) GetType() string {
	return paid.Type
}

type PaidMediaPhoto struct {
	Type  string      `json:"type"`
	Photo []PhotoSize `json:"photo"`
}

func (paid *PaidMediaPhoto) GetType() string {
	return paid.Type
}

type PaidMediaVideo struct {
	Type  string `json:"type"`
	Video Video  `json:"video"`
}

func (paid *PaidMediaVideo) GetType() string {
	return paid.Type
}

type PhotoSize struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	FileSize     int64  `json:"file_size,omitempty"`
}

type File struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	FileSize     int    `json:"file_size,omitempty"`
	FilePath     string `json:"file_path,omitempty"`
}

type MaskPosition struct {
	Point  string  `json:"point"`
	XShift float64 `json:"x_shift"`
	YShift float64 `json:"y_shift"`
	Scale  float64 `json:"scale"`
}

type Location struct {
	Latitude             float64 `json:"latitude"`
	Longitude            float64 `json:"longitude"`
	HorizontalAccuracy   float64 `json:"horizontal_accuracy,omitempty"`
	LivePeriod           int     `json:"live_period,omitempty"`
	Heading              int     `json:"heading,omitempty"`
	ProximityAlertRadius int     `json:"proximity_alert_radius,omitempty"`
}

type Dice struct {
	Emoji string `json:"emoji"`
	Value int    `json:"value"`
}
