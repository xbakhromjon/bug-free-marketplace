package domain

import (
	"time"
)

// User represents the user entity in the
type User struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	PhoneNumber string     `json:"phone_number"`
	Password    string     `json:"password"`
	Role        string     `json:"role"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

type UserRepository interface {
	Create(user *User) (*User, error)
	GetByID(userID int) (*User, error)
	GetByPhoneNumber(phoneNumber string) (*User, error)
	GetAll() ([]*User, error)
	UpdateByID(id int, name, phoneNumber, pass, role string) (*User, error)
	UpdateByPhoneNumber(name, phoneNumber, pass, role string) (*User, error)
	Delete(userID int) error
}

type UserUsecase interface {
	RegisterUser(user *User) (*User, error)
	LoginUser(phoneNumber, pass string) (bool, error)
	GetUserDataByID(userID int) (*User, error)
}
