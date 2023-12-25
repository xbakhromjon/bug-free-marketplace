package domain

type BasketItemRepository interface {
	AddItem(items *BasketItems) (int, error)
	GetAll(basketId int) ([]BasketItems, error)
	GetActiveBasket(userID int) (*Basket, error)
	UpdateBasketItem(bItemId, quantity int) error
	DeleteProduct(bItemId int) (int, error)
}
