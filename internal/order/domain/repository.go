package domain

type CartRepository interface {
}
type CartItemsRepository interface {
	AddItem(productId int, userId int, quantity int) error
	GetCardId(userId int) (carId int)
	RemoveItem(productId int, userId int) error
	GetAll(userId int) ([]*CartItems, error)
	RemoveAll(userId int) error
}
