package domain

import "time"

type UserFactory struct{}

func (f UserFactory) ParseModelToDomain(
	id int,
	name,
	phoneNumber,
	password,
	role string,
	createAt time.Time,
) *User {
	return &User{
		id:          id,
		name:        name,
		phoneNumber: phoneNumber,
		password:    password,
		role:        role,
		createAt:    time.Now().UTC(),
	}
}

func CreateUserFactory(newUser *NewUser) *User {
	currentTime := time.Now()
	return &User{
		name:        newUser.GetName(),
		phoneNumber: newUser.GetPhoneNumber(),
		password:    newUser.GetPassword(),
		role:        newUser.GetRole(),
		createAt:    currentTime,
		updatedAt:   currentTime,

		deletedAt: nil,
	}
}

const (
	ErrUserNotFound       = Err("user not found")
	ErrEmptyUserName      = Err("user name can not be empty")
	ErrEmptyPhoneNumber   = Err("phone number can not be empty")
	ErrInvalidCredentials = Err("bad credentials")

	//ErrInvalidName   = Err("shop name max length must be 128 characters")
	ErrPhoneNumberExists = Err("This phone number already exists")
)

type Err string

func (e Err) Error() string {
	return string(e)
}
