package models

import (
	"time"

	"github.com/LigeronAhill/planify/internal/enums/category"
	"github.com/LigeronAhill/planify/internal/enums/priority"
	"github.com/LigeronAhill/planify/internal/enums/status"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	AuthorID    uint
	Author      User
	ExecutorID  uint
	Executor    User
	Category    category.Category
	Priority    priority.Priority
	Status      status.Status
	Name        string
	Description string
	Deadline    time.Time
}

func NewTask(author *User) *Task {
	deadline := time.Now().Add(time.Hour * 24 * 31)
	return &Task{
		AuthorID:   author.ID,
		Author:     *author,
		ExecutorID: author.ID,
		Executor:   *author,
		Category:   category.TASK,
		Priority:   priority.NORMAL,
		Status:     status.NEW,
		Deadline:   deadline,
	}
}

func (t *Task) SetExecutor(executor *User) *Task {
	t.ExecutorID = executor.ID
	t.Executor = *executor
	return t
}

func (t *Task) SetCategory(category category.Category) *Task {
	t.Category = category
	return t
}

func (t *Task) SetPriority(priority priority.Priority) *Task {
	t.Priority = priority
	return t
}

func (t *Task) SetStatus(status status.Status) *Task {
	t.Status = status
	return t
}

func (t *Task) SetName(name string) *Task {
	t.Name = name
	return t
}

func (t *Task) SetDescription(description string) *Task {
	t.Description = description
	return t
}

func (t *Task) SetDeadline(deadline time.Time) *Task {
	t.Deadline = deadline
	return t
}
