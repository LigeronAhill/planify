package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitConfig(t *testing.T) {
	c, err := Init("../../config_example.toml")
	assert.NoError(t, err)
	assert.Equal(t, "someToken", c.GetString("TELEGRAM_BOT_TOKEN"))
	_, err = Init("wrong.yml")
	assert.Error(t, err)
}
