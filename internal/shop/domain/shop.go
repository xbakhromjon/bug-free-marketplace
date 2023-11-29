package domain

type Shop struct {
	Id     int
	Name   string
	UserId int
}

type NewShop struct {
	Name   string
	UserId int
}
