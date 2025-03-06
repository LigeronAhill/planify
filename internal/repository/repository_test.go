package repository

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	ctx := context.TODO()
	assert := assert.New(t)
	repo, err := New(ctx, "test.db")
	assert.NoError(err)
	err = repo.Close()
	assert.NoError(err)
}

func TestMigrate(t *testing.T) {
	ctx := context.TODO()
	assert := assert.New(t)
	repo, err := New(ctx, "test.db")
	assert.NoError(err)
	defer func() {
		if err := repo.Close(); err != nil {
			panic(err)
		}
	}()
	err = repo.Migrate(ctx)
	assert.NoError(err)
}
