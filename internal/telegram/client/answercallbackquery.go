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

type AnswerCallbackQueryRequest struct {
	CallbackQueryID string `json:"callback_query_id"`
	Text            string `json:"text,omitempty"`
	ShowAlert       bool   `json:"show_alert,omitempty"`
	URL             string `json:"url,omitempty"`
	CacheTime       int    `json:"cache_time,omitempty"`
}

func (b *Bot) answerCallbackQuery(ctx context.Context, queryID, text string) error {
	op := "answer callback query"
	message := AnswerCallbackQueryRequest{
		CallbackQueryID: queryID,
		Text:            text,
		CacheTime:       3,
	}
	jsonBody, err := json.Marshal(message)
	if err != nil {
		return e.Wrap(op, err)
	}
	reader := bytes.NewReader(jsonBody)
	body, err := b.request(ctx, methods.AnswerCallbackQuery, reader)
	if err != nil {
		return e.Wrap(op, err)
	}
	var res types.Response[bool]
	err = json.Unmarshal(body, &res)
	if err != nil {
		return e.Wrap(op, err)
	}
	slog.Debug("message sent", slog.Bool("result", res.Result))
	return nil
}
