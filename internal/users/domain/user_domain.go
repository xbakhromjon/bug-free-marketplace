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
	Name        string
	PhoneNumber string
	Password    string
}

// User-getter
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

func (u *User) GetDeletedAt() *time.Time {
	return u.deletedAt
}

// User-setter
func (u *User) SetID(id int) {
	u.id = id
}

func (u *User) SetName(name string) {
	u.name = name
}

func (u *User) SetPhoneNumber(phoneNumber string) {
	u.phoneNumber = phoneNumber
}

func (u *User) SetPassword(password string) {
	u.password = password
}

func (u *User) SetRole(role string) {
	u.role = role
}

func (u *User) SetCreateAt(createAt time.Time) {
	u.createAt = createAt
}

func (u *User) SetUpdatedAt(updatedAt time.Time) {
	u.updatedAt = updatedAt
}

func (u *User) SetDeletedAt(deletedAt *time.Time) {
	u.deletedAt = deletedAt
}

type UserRepository interface {
	Save(user *User) (int, error)
	FindOneByPhoneNumber(phoneNumber string) (*User, error)
	FindByID(userID int) (*User, error)
	UserExists(userID int) (bool, error)
	UserExistByPhone(phone string) (bool, error)
}
