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

func (b *Bot) sendTextMessage(ctx context.Context, chatID int, text string) error {
	op := fmt.Sprintf("sendMessage to: %d\n%s", chatID, text)
	slog.Debug(op)
	message := types.NewMessage(chatID, text)
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
	slog.Debug(fmt.Sprintf("Message sent: %+v", res.Result))
	return nil
}
