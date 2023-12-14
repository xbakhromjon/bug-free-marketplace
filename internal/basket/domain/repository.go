package domain

type CartRepository interface {
	Create(cart *Cart) (int, error)
	CreateCardItem(cart *CartItems) (int, error)
	GetById(id int) (*CartItems, error)
	UpdateCartItem(userId, productId, quantity int) error
	DeleteProduct(cardId, productId int) error
	GetByUserId(userID int) (*Cart, error)
	GetCardItem(cartId, productId int) (*CartItems, error)
}
