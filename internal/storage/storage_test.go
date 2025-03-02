package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStorage(t *testing.T) {
	file := "test.db"
	r, err := New(file)
	assert.NoError(t, err)
	err = r.Close()
	assert.NoError(t, err)
}
