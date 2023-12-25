package domain

import "golang-project-template/internal/order/adapters"

type BasketRepository interface {
	CreateBasket(userId int) (int, error)
	GetBasket(basketId int) (*adapters.BasketWithItems, error)
}
