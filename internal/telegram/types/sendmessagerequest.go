package types

type SendMessageRequest struct {
	ChatID      int       `json:"chat_id"`
	Text        string    `json:"text"`
	ParseMode   ParseMode `json:"parse_mode,omitempty"`
	ReplyMarkup any       `json:"reply_markup,omitempty"`
}

type ParseMode string

const (
	MardownV2 ParseMode = "MarkdownV2"
	HTML      ParseMode = "HTML"
	Markdown  ParseMode = "Markdown"
)

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]*InlineKeyboardButton `json:"inline_keyboard"`
}

type InlineKeyboardButton struct {
	Text         string `json:"text"`
	URL          string `json:"url,omitempty"`
	CallbackData string `json:"callback_data,omitempty"`
}

type ReplyKeyboardMarkup struct {
	Keyboard              [][]*KeyboardButton `json:"keyboard"`
	IsPersistent          bool                `json:"is_persistent,omitempty"`
	ResizeKeyboard        bool                `json:"resize_keyboard,omitempty"`
	OneTimeKeyboard       bool                `json:"one_time_keyboard,omitempty"`
	InputFieldPlaceholder string              `json:"input_field_placeholder,omitempty"`
	Selective             bool                `json:"selective,omitempty"`
}

type KeyboardButton struct {
	Text            string                      `json:"text"`
	RequestUsers    *KeyboardButtonRequestUsers `json:"request_users,omitempty"`
	RequestChat     *KeyboardButtonRequestChat  `json:"request_chat,omitempty"`
	RequestContact  bool                        `json:"request_contact,omitempty"`
	RequestLocation bool                        `json:"request_location,omitempty"`
}

type KeyboardButtonRequestUsers struct {
	RequestID       int  `json:"request_id"`
	UserIsBot       bool `json:"user_is_bot,omitempty"`
	UserIsPremium   bool `json:"user_is_premium,omitempty"`
	MaxQuantity     int  `json:"max_quantity,omitempty"`
	RequestName     bool `json:"request_name,omitempty"`
	RequestUsername bool `json:"request_username,omitempty"`
	RequestPhoto    bool `json:"request_photo,omitempty"`
}

type KeyboardButtonRequestChat struct {
	RequestID       int  `json:"request_id"`
	ChatIsChannel   bool `json:"chat_is_channel,omitempty"`
	ChatIsForum     bool `json:"chat_is_forum,omitempty"`
	ChatHasUsername bool `json:"chat_has_username,omitempty"`
	ChatIsCreated   bool `json:"chat_is_created,omitempty"`
}

func NewMessage(chatID int, text string) *SendMessageRequest {
	return &SendMessageRequest{
		ChatID: chatID,
		Text:   text,
	}
}
