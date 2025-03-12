package client

import (
	"context"
	"testing"

	"github.com/LigeronAhill/planify/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestGetUpdates(t *testing.T) {
	assert := assert.New(t)
	ctx := context.TODO()
	c, err := config.Init("config.toml")
	assert.NoError(err)
	token := c.GetString("TELEGRAM_BOT_TOKEN")
	bot, err := New(ctx, token)
	assert.NoError(err)
	_, err = bot.getUpdates(ctx)
	assert.NoError(err)
}
