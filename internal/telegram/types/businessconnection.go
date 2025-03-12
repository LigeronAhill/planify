package types

import (
	"encoding/json"
	"time"
)

type BusinessConnection struct {
	ID         string   `json:"id"`
	User       User     `json:"user"`
	UserChatID int      `json:"user_chat_id"`
	Date       UnixTime `json:"date"`
	CanReply   bool     `json:"can_reply"`
	IsEnabled  bool     `json:"is_enabled"`
}
type UnixTime struct {
	time.Time
}

func (u *UnixTime) UnmarshalJSON(b []byte) error {
	var timestamp int64
	err := json.Unmarshal(b, &timestamp)
	if err != nil {
		return err
	}
	u.Time = time.Unix(timestamp, 0)
	return nil
}

func (u UnixTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.Unix())
}
