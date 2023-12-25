package domain

import "time"

type Shop struct {
	id        int
	name      string
	ownerId   int
	createdAt time.Time
	updatedAt time.Time
}

// type NewShop struct {
// 	name    string
// 	ownerId int
// }

// func (s *NewShop) GetName() string {
// 	return s.name
// }

// func (s *NewShop) GetOwnerId() int {
// 	return s.ownerId
// }

// func (s *NewShop) SetName(name string) {
// 	s.name = name
// }

// func (s *NewShop) SetOwnerId(ownerId int) {
// 	s.ownerId = ownerId
// }

// func MakeNewShop(name string, ownerId int) NewShop {
// 	return NewShop{
// 		name:    name,
// 		ownerId: ownerId,
// 	}
// }

func (s *Shop) GetId() int {
	return s.id
}

func (s *Shop) GetName() string {
	return s.name
}

func (s *Shop) GetOwnerId() int {
	return s.ownerId
}

func (s *Shop) GetCreatedAt() time.Time {
	return s.createdAt
}

func (s *Shop) GetUpdatedAt() time.Time {
	return s.updatedAt
}

func (s *Shop) SetId(id int) {
	s.id = id
}

func (s *Shop) SetName(name string) {
	s.name = name
}

func (s *Shop) SetOwnerId(ownerId int) {
	s.ownerId = ownerId
}
func (s *Shop) SetCreateAt(createdAt time.Time) {
	s.createdAt = createdAt
}
func (s *Shop) SetUpdatedAt(updatedAt time.Time) {
	s.updatedAt = updatedAt
}
