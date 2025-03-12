package client

import (
	"context"
	"encoding/json"

	"github.com/LigeronAhill/planify/internal/e"
	"github.com/LigeronAhill/planify/internal/telegram/methods"
	"github.com/LigeronAhill/planify/internal/telegram/types"
)

func (b *Bot) getMe(ctx context.Context) (*types.User, error) {
	op := "GetMe"
	body, err := b.request(ctx, methods.GetMe, nil)
	if err != nil {
		return nil, e.Wrap(op, err)
	}
	var response types.Response[types.User]
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, e.Wrap(op, err)
	}
	return &response.Result, nil
}
