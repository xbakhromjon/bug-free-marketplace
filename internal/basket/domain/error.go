package domain

const (
	ErrCartItemNotFound   = Err("Cart item not found")
	ErrInvalidCartId      = Err("Invalid cart id")
	ErrInvalidId          = Err("Invalid id")
	ErrItemCreationFailed = Err("Cart Item creation failed")
	ErrCartUpdateFailed   = Err("Failed to update cart")
	ErrIDScanFailed       = Err("Failed to scan id")
	ErrAddItemFailed      = Err("Failed to add Item")
	ErrDeleteItemFailed   = Err("Failed to delete Item")
)

type Err string

func (e Err) Error() string {
	return string(e)
}
