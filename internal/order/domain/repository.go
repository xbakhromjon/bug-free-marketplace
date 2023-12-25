package domain

type BasketRepository interface {
	CreateBasket(userId int) (int, error)
	AddItem(items *BasketItems) (int, error)
	GetAll(basketId int) ([]BasketItems, error)
	GetActiveBasket(basketID int) (*Basket, error)
	UpdateBasketItem(basketId, quantity int) error
	DeleteProduct(basketId, productId int) (int, error)
}
