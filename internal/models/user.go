package models

import (
	"time"
)

type User struct {
	UserID    int
	FirstName string
	LastName  string
	Username  string
	Created   time.Time
	Updated   time.Time
}
