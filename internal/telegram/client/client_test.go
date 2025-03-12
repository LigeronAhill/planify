package client

import (
	"context"
	"testing"

	"github.com/LigeronAhill/planify/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	ctx := context.TODO()
	assert := assert.New(t)

	c, err := config.Init("config.toml")
	assert.NoError(err)

	token := c.GetString("TELEGRAM_BOT_TOKEN")

	_, err = New(ctx, token)
	assert.NoError(err)
}
