package domain

type BasketRepository interface {
	CreateBasket(userId int) (int, error)
	AddItem(items *BasketItems) (int, error)
	GetAll(basketId int) ([]BasketItems, error)
	UpdateBasketItem(basketId, quantity int) error
	UpdateBasketStatus(basketId int) error
	DeleteProduct(basketId, productId int) (int, error)
}
