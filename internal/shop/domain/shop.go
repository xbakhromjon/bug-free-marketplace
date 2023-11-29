package domain

type Shop struct {
	Id   int
	Name string
}

type NewShop struct {
	Id   int
	Name string
}

type ShopRepository interface {
	Save(shop Shop) (int, error)
}
