package domain

type BasketRepository interface {
	CreateBasket(userId int) (int, error)
	GetBasketWithItems(basketId int) (*BasketWithItems, error)
	GetActiveBasket(userID int) (*Basket, error)
	MarkBasketAsPurchased(userId, basketId int) error
}
