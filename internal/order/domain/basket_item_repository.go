package domain

type BasketItemRepository interface {
	AddItem(items *BasketItems) (int, error)
	GetAll(basketId int) ([]BasketItems, error)
	UpdateBasketItem(bItemId, quantity int) error
	DeleteProduct(bItemId int) (int, error)
}
