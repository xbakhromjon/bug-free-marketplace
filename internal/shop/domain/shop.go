package domain

type Shop struct {
	id        int
	name      string
	ownerId   int
	createdAt string
	updatedAt string
}

type NewShop struct {
	name    string `json:"name"`
	ownerId int    `json:"owner_id"`
}

func (s *NewShop) GetName() string {
	return s.name
}

func (s *NewShop) GetOwnerId() int {
	return s.ownerId
}

func (s *NewShop) SetName(name string) {
	s.name = name
}

func (s *NewShop) SetOwnerId(ownerId int) {
	s.ownerId = ownerId
}

func MakeNewShop(name string, ownerId int) NewShop {
	return NewShop{
		name:    name,
		ownerId: ownerId,
	}
}

func (s *Shop) GetId() int {
	return s.id
}

func (s *Shop) GetName() string {
	return s.name
}

func (s *Shop) GetOwnerId() int {
	return s.ownerId
}

func (s *Shop) GetCreatedAt() string {
	return s.createdAt
}

func (s *Shop) GetUpdatedAt() string {
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
func (s *Shop) SetCreateAt(createdAt string) {
	s.createdAt = createdAt
}
func (s *Shop) SetUpdatedAt(updatedAt string) {
	s.updatedAt = updatedAt
}
