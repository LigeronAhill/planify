package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/LigeronAhill/planify/internal/e"
	"github.com/LigeronAhill/planify/internal/telegram/methods"
	"github.com/LigeronAhill/planify/internal/telegram/types"
)

func (b *Bot) sendMessage(ctx context.Context, message *types.SendMessageRequest) error {
	op := "sending message"
	slog.Debug(op, slog.Int("to chat", message.ChatID), slog.String("text", message.Text))
	jsonBody, err := json.Marshal(message)
	if err != nil {
		return e.Wrap(op, err)
	}
	reader := bytes.NewReader(jsonBody)
	body, err := b.request(ctx, methods.SendMessage, reader)
	if err != nil {
		return e.Wrap(op, err)
	}
	var res types.Response[types.Message]
	err = json.Unmarshal(body, &res)
	if err != nil {
		return e.Wrap(op, err)
	}
	slog.Debug("message sent", slog.String("message", fmt.Sprintf("%+v", res.Result)))
	return nil
}
