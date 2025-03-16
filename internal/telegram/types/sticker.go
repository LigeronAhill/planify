package types

type Sticker struct {
	FileID           string        `json:"file_id"`
	FileUniqueID     string        `json:"file_unique_id"`
	Type             string        `json:"type"`
	Width            int           `json:"width"`
	Height           int           `json:"height"`
	IsAnimated       bool          `json:"is_animated"`
	IsVideo          bool          `json:"is_video"`
	Thumbnail        *PhotoSize    `json:"thumbnail"`
	Emoji            string        `json:"emoji"`
	SetName          string        `json:"set_name"`
	PremiumAnimation *File         `json:"premium_animation"`
	MaskPosition     *MaskPosition `json:"mask_position"`
	CustomEmojiID    string        `json:"custom_emoji_id"`
	NeedsRepainting  bool          `json:"needs_repainting"`
	FileSize         int           `json:"file_size"`
}

type MaskPosition struct {
	Point  string  `json:"point"`
	XShift float64 `json:"x_shift"`
	YShift float64 `json:"y_shift"`
	Scale  float64 `json:"scale"`
}
