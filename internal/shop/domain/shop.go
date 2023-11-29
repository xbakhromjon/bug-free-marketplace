package domain

type Shop struct {
	Id   int
	Name string
}

type ShopRequest struct {
	Id   int
	Name string
}

type ShopResponse struct {
	Id   int
	Name string
}

type ShopUseCase interface {
	Create(req ShopRequest) (int, error)
}

type ShopRepository interface {
	Save(shop Shop) (int, error)
}
