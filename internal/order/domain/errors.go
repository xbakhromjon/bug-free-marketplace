package domain

const (
	ErrBasketCreationFailed  = Err("Failed to create basket")
	ErrIDScanFailed          = Err("Failed to scan id")
	ErrBasketUpdateFailed    = Err("Failed to update basket")
	ErrDeleteItemFailed      = Err("Failed to delete Item")
	ErrAddItemFailed         = Err("Couldn't update quantity")
	ErrBasketNotFound        = Err("Basket not found")
	ErrGetActiveBasketFailed = Err("Get non ordered basket failed")
	ErrBasketItemsNotFound   = Err("Basket items not found")
)

type Err string

func (e Err) Error() string {
	return string(e)
}
