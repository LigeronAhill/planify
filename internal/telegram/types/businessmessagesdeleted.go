package types

type BusinessMessagesDeleted struct {
	BusinessConnectionID string `json:"business_connection_id"`
	Chat                 *Chat  `json:"chat"`
	MessageIDs           []int  `json:"message_ids"`
}
