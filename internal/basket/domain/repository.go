package domain

type CartRepository interface {
	Create(cart *Cart) (int, error)
	CreateCardItem(cart *CartItems) (int, error)
	GetCart(id int) (*Cart, error)
	UpdateCartItem(userId, productId, quantity int) error
	DeleteProduct(cardId, productId int) error
	GetByUserId(userID int) (*Cart, error)
	GetCardItem(cartId int) (*CartItems, error)
	GetCartItemByCartIdAndProductId(cartId, productId int) (*CartItems, error)
}
