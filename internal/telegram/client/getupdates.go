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

func (b *Bot) getUpdates(ctx context.Context) ([]*types.Update, error) {
	op := "getting updates"
	slog.Debug(op, slog.Int("limit", b.limit), slog.Int("offset", b.offset))
	reqBody := types.UpdatesRequest{
		Limit:  10,
		Offset: b.offset,
	}
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, e.Wrap(op, err)
	}
	reader := bytes.NewReader(jsonBody)
	body, err := b.request(ctx, methods.GetUpdates, reader)
	if err != nil {
		return nil, e.Wrap(op, err)
	}
	var res types.Response[[]*types.Update]
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, e.Wrap(op, err)
	}
	updates := res.Result
	if len(updates) != 0 {
		b.offset = updates[len(updates)-1].UpdateID + 1
		return updates, nil
	} else {
		return nil, nil
	}
}
