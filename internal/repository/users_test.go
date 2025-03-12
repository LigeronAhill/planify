package repository

import (
	"context"
	"testing"
	"time"

	"github.com/LigeronAhill/planify/internal/models"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
)

func fakeuser() *models.User {
	id := gofakeit.Int()
	firstName := gofakeit.FirstName()
	lastName := gofakeit.LastName()
	username := gofakeit.Username()
	return &models.User{
		UserID:    id,
		FirstName: firstName,
		LastName:  lastName,
		Username:  username,
		Created:   time.Now(),
		Updated:   time.Now(),
	}
}

func TestGetAllUsers(t *testing.T) {
	ctx := context.TODO()
	assert := assert.New(t)
	repo, err := New(ctx, "test.db")
	assert.NoError(err)
	defer func() {
		if err := repo.Close(); err != nil {
			panic(err)
		}
	}()
	_, err = repo.GetAllUsers(ctx)
	assert.NoError(err)
}

func TestInsertUser(t *testing.T) {
	ctx := context.TODO()
	assert := assert.New(t)
	repo, err := New(ctx, "test.db")
	assert.NoError(err)
	defer func() {
		if err := repo.Close(); err != nil {
			panic(err)
		}
	}()
	user := fakeuser()
	created, err := repo.InsertUser(ctx, user)
	assert.NoError(err)
	assert.Equal(user.Username, created.Username)
}

func TestGetUser(t *testing.T) {
	ctx := context.TODO()
	assert := assert.New(t)
	repo, err := New(ctx, "test.db")
	assert.NoError(err)
	defer func() {
		if err := repo.Close(); err != nil {
			panic(err)
		}
	}()
	user := fakeuser()
	created, err := repo.InsertUser(ctx, user)
	assert.NoError(err)
	assert.Equal(user.Username, created.Username)
	stored, err := repo.GetUser(ctx, user.UserID)
	assert.NoError(err)
	assert.Equal(stored.Username, created.Username)
}

func TestDeleteUser(t *testing.T) {
	ctx := context.TODO()
	assert := assert.New(t)
	repo, err := New(ctx, "test.db")
	assert.NoError(err)
	defer func() {
		if err := repo.Close(); err != nil {
			panic(err)
		}
	}()
	user := fakeuser()
	created, err := repo.InsertUser(ctx, user)
	assert.NoError(err)
	assert.Equal(user.Username, created.Username)
	err = repo.DeleteUser(ctx, created.UserID)
	assert.NoError(err)
	_, err = repo.GetUser(ctx, created.UserID)
	assert.Error(err)
}
