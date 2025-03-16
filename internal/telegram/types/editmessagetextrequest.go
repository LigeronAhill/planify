package types

type EditMessageTextRequest struct {
	ChatID          int       `json:"chat_id"`
	MessageID       int       `json:"message_id,omitempty"`
	InlineMessageID int       `json:"inline_message_id,omitempty"`
	Text            string    `json:"text"`
	ParseMode       ParseMode `json:"parse_mode,omitempty"`
	ReplyMarkup     any       `json:"reply_markup,omitempty"`
}
