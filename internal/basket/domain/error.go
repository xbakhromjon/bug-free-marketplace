package domain

const (
	ErrCartNotFound           = Err("Cart not found")
	ErrCartItemNotFound       = Err("Cart item not found")
	ErrProductNotFound        = Err("Product not found in the cart")
	ErrCartCreationFailed     = Err("Cart creation failed")
	ErrCartItemCreationFailed = Err("Cart Item creation failed")
	ErrCartUpdateFailed       = Err("Failed to update cart")
)

type Err string

func (e Err) Error() string {
	return string(e)
}
