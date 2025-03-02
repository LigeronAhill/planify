package storage

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStorage(t *testing.T) {
	ctx := context.Background()
	file := "test.db"
	r, err := New(ctx, file)
	assert.NoError(t, err)
	err = r.Close()
	assert.NoError(t, err)
}
