package models

import (
	"time"

	"github.com/LigeronAhill/planify/internal/enums"
)

type Task struct {
	ID          int
	AuthorID    int
	ExecutorID  int
	Title       string
	Description string
	Status      enums.Status
	Priority    enums.Priority
	Category    enums.Category
	Created     time.Time
	Updated     time.Time
}
