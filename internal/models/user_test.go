package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestUser(t *testing.T) {
	newUser := NewUser(42)
	newUser.SetFirstName("Bob")
	newUser.SetLastName("Marly")
	newUser.SetUsername("Banana")
	want := &User{
		Model: gorm.Model{
			ID: 42,
		},
		Username: "Banana",
	}
	assert.Equal(t, want.Username, newUser.Username)
	assert.Equal(t, "Bob Marly", newUser.FullName())
}
