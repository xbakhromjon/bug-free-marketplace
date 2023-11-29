package domain

import "time"

type UserFactory struct{}

func (f UserFactory) MapUserData(
	id int,
	name,
	phoneNumber,
	role string,
	createdAt time.Time,
) *User {
	return &User{
		ID:          id,
		Name:        name,
		PhoneNumber: phoneNumber,
		Role:        role,
		CreatedAt:   createdAt,
	}
}
