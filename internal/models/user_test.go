package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	newUser := NewUser(42)
	newUser.SetFirstName("Bob")
	newUser.SetLastName("Marly")
	newUser.SetUsername("Banana")
	want := &User{
		username:  "Banana",
		firstName: "Bob",
		lastName:  "Marly",
	}
	assert.Equal(t, want.username, newUser.username)
	assert.Equal(t, "Bob Marly", newUser.FullName())
	assert.Equal(t, want.FirstName(), newUser.FirstName())
	assert.Equal(t, want.LastName(), newUser.LastName())
}
