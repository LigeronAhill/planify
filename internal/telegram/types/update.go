package types

type Update struct {
	UpdateID                int                          `json:"update_id"`
	Message                 *Message                     `json:"message"`
	EditedMessage           *Message                     `json:"edited_message"`
	ChannelPost             *Message                     `json:"channel_post"`
	EditedChannelPost       *Message                     `json:"edited_channel_post"`
	BusinessConnection      *BusinessConnection          `json:"business_connection"`
	BusinessMessage         *Message                     `json:"business_message"`
	EditedBusinessMessage   *Message                     `json:"edited_business_message"`
	DeletedBusinessMessages *BusinessMessagesDeleted     `json:"deleted_business_messages"`
	MessageReaction         *MessageReactionUpdated      `json:"message_reaction"`
	MessageReactionCount    *MessageReactionCountUpdated `json:"message_reaction_count"`
	InlineQuery             *InlineQuery                 `json:"inline_query"`
	ChosenInlineResult      *ChosenInlineResult          `json:"chosen_inline_result"`
	CallbackQuery           *CallbackQuery               `json:"callback_query"`
}
