package repository

import (
	"context"
	"testing"

	"github.com/LigeronAhill/planify/internal/models"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
)

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
	id := gofakeit.Int()
	firstName := gofakeit.FirstName()
	lastName := gofakeit.LastName()
	username := gofakeit.Username()
	user := models.NewUser(int(id)).SetFirstName(firstName).SetLastName(lastName).SetUsername(username)
	created, err := repo.InsertUser(ctx, user)
	assert.NoError(err)
	assert.Equal(user.Username(), created.Username())
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
	id := gofakeit.Int()
	firstName := gofakeit.FirstName()
	lastName := gofakeit.LastName()
	username := gofakeit.Username()
	user := models.NewUser(int(id)).SetFirstName(firstName).SetLastName(lastName).SetUsername(username)
	created, err := repo.InsertUser(ctx, user)
	assert.NoError(err)
	assert.Equal(user.Username(), created.Username())
	stored, err := repo.GetUser(ctx, id)
	assert.NoError(err)
	assert.Equal(stored.Username(), created.Username())
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
	id := gofakeit.Int()
	firstName := gofakeit.FirstName()
	lastName := gofakeit.LastName()
	username := gofakeit.Username()
	user := models.NewUser(int(id)).SetFirstName(firstName).SetLastName(lastName).SetUsername(username)
	created, err := repo.InsertUser(ctx, user)
	assert.NoError(err)
	assert.Equal(user.Username(), created.Username())
	err = repo.DeleteUser(ctx, created.ID())
	assert.NoError(err)
	_, err = repo.GetUser(ctx, created.ID())
	assert.Error(err)
}
