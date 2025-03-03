package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	AuthorID    uint
	Author      User
	ExecutorID  uint
	Executor    User
	Category    string
	Priority    string
	Status      string
	Name        string
	Description string
	Deadline    time.Time
}
