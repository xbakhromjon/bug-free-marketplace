package domain

type BasketRepository interface {
	CreateBasket(userId int) (int, error)
}
