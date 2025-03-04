package storage

import (
	"testing"

	"github.com/LigeronAhill/planify/internal/models"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
)

func user() *models.User {
	id := gofakeit.Uint16()
	name := gofakeit.FirstName()
	lastName := gofakeit.LastName()
	username := gofakeit.Username()
	u := models.NewUser(uint(id))
	u.SetFirstName(name)
	u.SetLastName(lastName)
	u.SetUsername(username)
	return u
}

func TestInsertUser(t *testing.T) {
	assert := assert.New(t)
	repo, err := New("test.db")
	assert.NoError(err)
	u := user()
	err = repo.InsertUser(u)
	assert.NoError(err)
}

func TestGetUser(t *testing.T) {
	assert := assert.New(t)
	repo, err := New("test.db")
	assert.NoError(err)
	u := user()
	err = repo.InsertUser(u)
	assert.NoError(err)
	user, err := repo.GetUser(u.ID)
	assert.NoError(err)
	assert.Equal(u.LastName, user.LastName)
}

func TestGetAllUsers(t *testing.T) {
	assert := assert.New(t)
	repo, err := New("test.db")
	assert.NoError(err)
	u := user()
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
	u := models.NewUser(69)
	u.SetFirstName("Bobby")
	u.SetLastName("Marly")
	u.SetUsername("Bananas")
	err = repo.InsertUser(u)
	assert.NoError(err)
	err = repo.DeleteUser(u.ID)
	assert.NoError(err)
}
