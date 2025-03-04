package storage

import (
	"testing"
	"time"

	"github.com/LigeronAhill/planify/internal/enums/priority"
	"github.com/LigeronAhill/planify/internal/enums/status"
	"github.com/LigeronAhill/planify/internal/models"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
)

func task(user *models.User) *models.Task {
	name := gofakeit.Sentence(2)
	desc := gofakeit.Sentence(10)
	t := models.NewTask(user).SetName(name).SetDescription(desc)
	return t
}

func TestInsertTask(t *testing.T) {
	assert := assert.New(t)
	repo, err := New("test.db")
	assert.NoError(err)
	user := user()
	task := task(user)
	_, err = repo.InsertTask(task)
	assert.NoError(err)
}

func TestUpdateTask(t *testing.T) {
	assert := assert.New(t)
	repo, err := New("test.db")
	assert.NoError(err)
	user := user()
	task := task(user)
	id, err := repo.InsertTask(task)
	assert.NoError(err)
	inserted, err := repo.GetTask(id)
	assert.NoError(err)
	dl := time.Now().Add(time.Hour * 3)
	inserted.SetStatus(status.PENDING).SetPriority(priority.HIGH).SetDeadline(dl)
	upd, err := repo.UpdateTask(inserted)
	assert.NoError(err)
	assert.Equal(inserted.Status, upd.Status)
	assert.Equal(inserted.Priority, upd.Priority)
	assert.Equal(inserted.Deadline, upd.Deadline)
}

func TestGetTask(t *testing.T) {
	assert := assert.New(t)
	repo, err := New("test.db")
	assert.NoError(err)
	user := user()
	task := task(user)
	id, err := repo.InsertTask(task)
	assert.NoError(err)
	inserted, err := repo.GetTask(id)
	assert.NoError(err)
	assert.Equal(id, inserted.ID)
}

func TestGetNotCompletedTasksByAuthor(t *testing.T) {
	assert := assert.New(t)
	repo, err := New("test.db")
	assert.NoError(err)
	author := user()
	task := task(author)
	id, err := repo.InsertTask(task)
	assert.NoError(err)
	inserted, err := repo.GetTask(id)
	assert.NoError(err)
	res, err := repo.GetNotCompletedTasksByAuthor(author.ID)
	assert.NoError(err)
	assert.NotEmpty(res)
	assert.Contains(res, inserted)
}

func TestGetNotCompletedTasksByExecutor(t *testing.T) {
	assert := assert.New(t)
	repo, err := New("test.db")
	assert.NoError(err)
	author := user()
	executor := user()
	task := task(author).SetExecutor(executor)
	id, err := repo.InsertTask(task)
	assert.NoError(err)
	inserted, err := repo.GetTask(id)
	assert.NoError(err)
	res, err := repo.GetNotCompletedTasksByExecutor(executor.ID)
	assert.NoError(err)
	assert.NotEmpty(res)
	assert.Contains(res, inserted)
}

func TestGetCompletedTasksByAuthor(t *testing.T) {
	assert := assert.New(t)
	repo, err := New("test.db")
	assert.NoError(err)
	author := user()
	task := task(author).SetStatus(status.DONE)
	id, err := repo.InsertTask(task)
	assert.NoError(err)
	inserted, err := repo.GetTask(id)
	assert.NoError(err)
	res, err := repo.GetCompletedTasksByAuthor(author.ID)
	assert.NoError(err)
	assert.NotEmpty(res)
	assert.Contains(res, inserted)
}

func TestGetCompletedTasksByExecutor(t *testing.T) {
	assert := assert.New(t)
	repo, err := New("test.db")
	assert.NoError(err)
	author := user()
	executor := user()
	task := task(author).SetExecutor(executor).SetStatus(status.DONE)
	id, err := repo.InsertTask(task)
	assert.NoError(err)
	inserted, err := repo.GetTask(id)
	assert.NoError(err)
	res, err := repo.GetCompletedTasksByExecutor(executor.ID)
	assert.NoError(err)
	assert.NotEmpty(res)
	assert.Contains(res, inserted)
}

func TestDeleteTask(t *testing.T) {
	assert := assert.New(t)
	repo, err := New("test.db")
	assert.NoError(err)
	user := user()
	task := task(user)
	id, err := repo.InsertTask(task)
	assert.NoError(err)
	err = repo.DeleteTask(id)
	assert.NoError(err)
	_, err = repo.GetTask(id)
	assert.Error(err)
}
