package models

import (
	"testing"
	"time"

	"github.com/LigeronAhill/planify/internal/enums/category"
	"github.com/LigeronAhill/planify/internal/enums/priority"
	"github.com/LigeronAhill/planify/internal/enums/status"
	"github.com/stretchr/testify/assert"
)

func user() *User {
	newUser := NewUser(42)
	newUser.SetFirstName("Bob")
	newUser.SetLastName("Marly")
	newUser.SetUsername("Banana")
	return newUser
}

func TestNewTask(t *testing.T) {
	newUser := user()
	newTask := NewTask(newUser)
	assert.Equal(t, newTask.AuthorID, newUser.ID)
	assert.Equal(t, &newTask.Author, newUser)
}

func TestSetExecutor(t *testing.T) {
	newUser := user()
	newTask := NewTask(newUser).SetExecutor(newUser)
	assert.Equal(t, &newTask.Executor, newUser)
	assert.Equal(t, newTask.ExecutorID, newUser.ID)
}

func TestSetCategory(t *testing.T) {
	user := user()
	task := NewTask(user).SetCategory(category.CALL)
	assert.Equal(t, task.Category, category.CALL)
}

func TestSetPriority(t *testing.T) {
	user := user()
	task := NewTask(user).SetPriority(priority.HIGH)
	assert.Equal(t, task.Priority, priority.HIGH)
}

func TestSetStatus(t *testing.T) {
	user := user()
	task := NewTask(user).SetStatus(status.DONE)
	assert.Equal(t, task.Status, status.DONE)
}

func TestSetName(t *testing.T) {
	user := user()
	task := NewTask(user).SetName("test name")
	assert.Equal(t, task.Name, "test name")
}

func TestSetDescription(t *testing.T) {
	user := user()
	task := NewTask(user).SetDescription("test description")
	assert.Equal(t, task.Description, "test description")
}

func SetDeadline(t *testing.T) {
	dl := time.Now().Add(time.Hour * 3)
	user := user()
	task := NewTask(user).SetDeadline(dl)
	assert.Equal(t, task.Deadline, dl)
}
