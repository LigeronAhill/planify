package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	file := "test.db"
	_, err := New(file)
	assert.NoError(t, err)
}

func TestMigrate(t *testing.T) {
	r, err := New("test.db")
	assert.NoError(t, err)
	err = r.Migrate()
	assert.NoError(t, err)
}
