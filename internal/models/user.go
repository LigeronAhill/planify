package models

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Username  string
}

func NewUser(id uint) *User {
	u := &User{}
	u.ID = id
	return u
}

func (u *User) SetFirstName(firstName string) *User {
	u.FirstName = firstName
	return u
}

func (u *User) SetLastName(lastName string) *User {
	u.LastName = lastName
	return u
}

func (u *User) SetUsername(username string) *User {
	u.Username = username
	return u
}

func (u *User) FullName() string {
	if len(u.LastName) != 0 {
		return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
	} else {
		return u.FirstName
	}
}
