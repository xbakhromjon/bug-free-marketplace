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
		role:        "customer",
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
		deletedAt:   deletedAt, //
	}
}

//
//
