package client

import (
	"bytes"
	"context"
	"encoding/json"
	"log/slog"

	"github.com/LigeronAhill/planify/internal/e"
	"github.com/LigeronAhill/planify/internal/telegram/methods"
	"github.com/LigeronAhill/planify/internal/telegram/types"
)

type DeleteMessageRequest struct {
	ChatID    int `json:"chat_id"`
	MessageID int `json:"message_id"`
}

func (b *Bot) deleteMessage(ctx context.Context, chatID, messageID int) error {
	op := "deleting message"
	slog.Debug(op, slog.Int("chat", chatID), slog.Int("message", messageID))
	message := DeleteMessageRequest{
		ChatID:    chatID,
		MessageID: messageID,
	}
	jsonBody, err := json.Marshal(message)
	if err != nil {
		return e.Wrap(op, err)
	}
	reader := bytes.NewReader(jsonBody)
	body, err := b.request(ctx, methods.DeleteMessage, reader)
	if err != nil {
		return e.Wrap(op, err)
	}
	var res types.Response[bool]
	err = json.Unmarshal(body, &res)
	if err != nil {
		return e.Wrap(op, err)
	}
	slog.Debug("message deleted", slog.Bool("result", res.Result))
	return nil
}
