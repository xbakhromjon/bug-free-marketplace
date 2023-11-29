package domain

type Shop struct {
	Id     int
	Name   string
	UserId int
}

type NewShop struct {
	Id     int
	Name   string
	UserId int
}

type ShopRepository interface {
	Save(shop Shop) (int, error)
}
