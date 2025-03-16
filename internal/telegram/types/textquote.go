package types

type TextQuote struct {
	Text     string           `json:"text"`
	Entities []*MessageEntity `json:"entities,omitempty"`
	Position int              `json:"position,omitempty"`
	IsManual bool             `json:"is_manual,omitempty"`
}
