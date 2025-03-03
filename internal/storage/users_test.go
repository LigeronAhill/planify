package storage

import (
	"testing"

	"github.com/LigeronAhill/planify/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestInsertUser(t *testing.T) {
	assert := assert.New(t)
	repo, err := New("test.db")
	assert.NoError(err)
	u := models.NewUser(42)
	u.SetFirstName("Bobby")
	u.SetLastName("Marly")
	u.SetUsername("Bananas")
	err = repo.InsertUser(u)
	assert.NoError(err)
}

func TestGetUser(t *testing.T) {
	assert := assert.New(t)
	repo, err := New("test.db")
	assert.NoError(err)
	u := models.NewUser(42)
	u.SetFirstName("Bobby")
	u.SetLastName("Marly")
	u.SetUsername("Bananas")
	err = repo.InsertUser(u)
	assert.NoError(err)
	user, err := repo.GetUser(42)
	assert.NoError(err)
	assert.Equal("Marly", user.LastName)
}

func TestGetAllUsers(t *testing.T) {
	assert := assert.New(t)
	repo, err := New("test.db")
	assert.NoError(err)
	u := models.NewUser(42)
	u.SetFirstName("Bobby")
	u.SetLastName("Marly")
	u.SetUsername("Bananas")
	err = repo.InsertUser(u)
	assert.NoError(err)
	users, err := repo.GetAllUsers()
	assert.NoError(err)
	assert.NotEmpty(users)
}

func TestDeleteUser(t *testing.T) {
	assert := assert.New(t)
	repo, err := New("test.db")
	assert.NoError(err)
	u := models.NewUser(42)
	u.SetFirstName("Bobby")
	u.SetLastName("Marly")
	u.SetUsername("Bananas")
	err = repo.InsertUser(u)
	assert.NoError(err)
	err = repo.DeleteUser(u.ID)
	assert.NoError(err)
}
