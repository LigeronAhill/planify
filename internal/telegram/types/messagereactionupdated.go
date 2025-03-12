package types

type MessageReactionUpdated struct {
	Chat        *Chat           `json:"chat"`
	MessageID   int             `json:"message_id"`
	User        *User           `json:"user,omitempty"`
	ActorChat   Chat            `json:"actor_chat,omitempty"`
	Date        UnixTime        `json:"date"`
	OldReaction []*ReactionType `json:"old_reaction"`
	NewReaction []*ReactionType `json:"new_reaction"`
}
