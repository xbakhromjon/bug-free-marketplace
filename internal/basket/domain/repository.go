package domain

type CartRepository interface {
	CreateBasket(userId int) (int, error)
	AddItem(cart *CartItems) (int, error)
	UpdateCartItem(cartId, quantity int) error
	DeleteProduct(cartId, productId int) (id int, err error)
	GetAll(cartId int) ([]*CartItems, error)
}
