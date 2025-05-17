package attachments

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
