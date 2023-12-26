package domain

const (
	ErrBasketCreationFailed = Err("Failed to create basket")
	ErrIDScanFailed         = Err("Failed to scan id")
	ErrBasketUpdateFailed   = Err("Failed to update basket")
	ErrDeleteItemFailed     = Err("Failed to delete Item")
	ErrAddItemFailed        = Err("Couldn't update quantity")
	ErrBasketItemsNotFound
)

type Err string

func (e Err) Error() string {
	return string(e)
}
