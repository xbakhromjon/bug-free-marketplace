package domain

const (
	ErrCartNotFound           = Err("Cart not found")
	ErrCartItemNotFound       = Err("Cart item not found")
	ErrCartCreationFailed     = Err("Cart creation failed")
	ErrCartItemCreationFailed = Err("Cart Item creation failed")
	ErrCartUpdateFailed       = Err("Failed to update cart")
	ErrIDScanFailed           = Err("Failed to scan id")
	ErrAddItemFailed          = Err("Failed to add Item")
	ErrDeleteItemFailed       = Err("Failed to delete Item")
)

type Err string

func (e Err) Error() string {
	return string(e)
}
