package domain

import (
	"time"
)

// User represents the user entity in the
type User struct {
	id          int
	name        string
	phoneNumber string
	password    string
	role        string
	createAt    time.Time
	updatedAt   time.Time
	deletedAt   *time.Time
}

type NewUser struct {
	name        string
	phoneNumber string
	password    string
	role        string
}

// User
func (u *User) GetID() int {
	return u.id
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) GetPhoneNumber() string {
	return u.phoneNumber
}

func (u *User) GetPassword() string {
	return u.password
}

func (u *User) GetRole() string {
	return u.role
}

func (u *User) GetCreatedAt() time.Time {
	return u.createAt
}

func (u *User) GetUpdatedAt() time.Time {
	return u.updatedAt
}

func (u *User) GetDeletedAt() time.Time {
	return *u.deletedAt
}

// NewUser
func (n *NewUser) GetName() string {
	return n.name
}

func (n *NewUser) GetPhoneNumber() string {
	return n.phoneNumber
}

func (n *NewUser) GetPassword() string {
	return n.password
}

func (n *NewUser) GetRole() string {
	return n.role
}

type UserRepository interface {
	Save(user *User) (int, error)
	FindOneByPhoneNumber(phoneNumber string) (*User, error)
}
