package models

import (
	"fmt"
	"time"
)

type User struct {
	id        int
	firstName string
	lastName  string
	username  string
	created   time.Time
	updated   time.Time
}

func NewUser(id int) *User {
	u := &User{}
	u.id = id
	u.created = time.Now()
	u.updated = time.Now()
	return u
}

func (u *User) SetFirstName(firstName string) *User {
	u.firstName = firstName
	return u
}

func (u *User) SetLastName(lastName string) *User {
	u.lastName = lastName
	return u
}

func (u *User) SetUsername(username string) *User {
	u.username = username
	return u
}

func (u *User) SetCreated(created time.Time) *User {
	u.created = created
	return u
}

func (u *User) SetUpdated(updated time.Time) *User {
	u.updated = updated
	return u
}

func (u *User) FullName() string {
	if len(u.lastName) != 0 {
		return fmt.Sprintf("%s %s", u.firstName, u.lastName)
	} else {
		return u.firstName
	}
}

func (u *User) ID() int {
	return u.id
}

func (u *User) FirstName() string {
	return u.firstName
}

func (u *User) LastName() string {
	return u.lastName
}

func (u *User) Username() string {
	return u.username
}

func (u *User) Created() time.Time {
	return u.created
}

func (u *User) Updated() time.Time {
	return u.updated
}
