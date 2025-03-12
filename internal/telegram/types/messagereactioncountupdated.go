package types

type MessageReactionCountUpdated struct {
	Chat      *Chat            `json:"chat"`
	MessageID int              `json:"message_id"`
	Date      UnixTime         `json:"date"`
	Reactions []*ReactionCount `json:"reactions"`
}

type ReactionCount struct {
	Type       *ReactionType
	TotalCount int
}
