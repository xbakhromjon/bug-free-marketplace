package domain

import "time"

type UserFactory struct{}

func NewUserFactory() *UserFactory {
	return &UserFactory{}
}

func (f UserFactory) CreateMerchantUser(user *NewUser) *User {
	return &User{
		name:        user.name,
		phoneNumber: user.phoneNumber,
		password:    user.phoneNumber,
		role:        "merchant",
		createAt:    time.Now().UTC(),
		updatedAt:   time.Now().UTC(),
		deletedAt:   nil,
	}
}

func (f UserFactory) CreateCustomerUser(user *NewUser) *User {
	return &User{
		name:        user.name,
		phoneNumber: user.phoneNumber,
		password:    user.phoneNumber,
		role:        "user",
		createAt:    time.Now().UTC(),
		updatedAt:   time.Now().UTC(),
		deletedAt:   nil,
	}
}

func (f UserFactory) ParseModelToDomain(
	id int,
	name,
	phoneNumber,
	password,
	role string,
	createdAt time.Time,
	updatedAt time.Time,
	deletedAt *time.Time,
) *User {
	return &User{
		id:          id,
		name:        name,
		phoneNumber: phoneNumber,
		password:    password,
		role:        role,
		createAt:    createdAt,
		updatedAt:   updatedAt,
		deletedAt:   deletedAt,
	}
}

const (
	ErrUserNotFound       = Err("user not found")
	ErrEmptyUserName      = Err("user name can not be empty")
	ErrEmptyPhoneNumber   = Err("phone number can not be empty")
	ErrInvalidCredentials = Err("bad credentials")
	//ErrInvalidName   = Err("shop name max length must be 128 characters")
	ErrPhoneNumberExists = Err("this phone number already exists")
)

type Err string

func (e Err) Error() string {
	return string(e)
}
