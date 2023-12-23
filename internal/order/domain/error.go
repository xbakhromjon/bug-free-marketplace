package domain

const (
	ErrBasketUpdateFailed = Err("Failed to update basket")
	ErrIDScanFailed       = Err("Failed to scan id")
	ErrDeleteItemFailed   = Err("Failed to delete Item")
)

type Err string

func (e Err) Error() string {
	return string(e)
}
