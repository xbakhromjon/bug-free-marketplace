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
		role:        "user",
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
