package types

type Message struct {
	MessageID            int              `json:"message_id"`
	MessageThreadID      int              `json:"message_thread_id,omitempty"`
	From                 *User            `json:"from,omitempty"`
	SenderChat           *Chat            `json:"sender_chat,omitempty"`
	SenderBoostCount     int              `json:"sender_boost_count,omitempty"`
	SenderBusinessBot    *User            `json:"sender_business_bot,omitempty"`
	Date                 UnixTime         `json:"date"`
	BusinessConnectionID string           `json:"business_connection_id,omitempty"`
	Chat                 *Chat            `json:"chat"`
	ForwardOrigin        *MessageOrigin   `json:"forward_origin,omitempty"`
	IsTopicMessage       bool             `json:"is_topic_message,omitempty"`
	IsAutomaticForward   bool             `json:"is_automatic_forward,omitempty"`
	ReplyToMessage       *Message         `json:"reply_to_message,omitempty"`
	Quote                *TextQuote       `json:"quote,omitempty"`
	ReplyToStory         *Story           `json:"reply_to_story,omitempty"`
	ViaBot               *User            `json:"via_bot,omitempty"`
	EditDate             UnixTime         `json:"edit_date"`
	Text                 string           `json:"text,omitempty"`
	Entities             []*MessageEntity `json:"entities,omitempty"`
	Animation            *Animation       `json:"animation"`
	Audio                *Audio           `json:"audio"`
	Document             *Document        `json:"document"`
	Photo                []*PhotoSize     `json:"photo"`
	Sticker              *Sticker         `json:"sticker"`
	Story                *Story           `json:"story"`
}

type MessageOrigin struct {
	Type            string   `json:"type,omitempty"`
	Date            UnixTime `json:"date"`
	SenderUser      *User    `json:"sender_user,omitempty"`
	SenderUserName  string   `json:"sender_user_name,omitempty"`
	SenderChat      *Chat    `json:"sender_chat,omitempty"`
	AuthorSignature string   `json:"author_signature,omitempty"`
}
