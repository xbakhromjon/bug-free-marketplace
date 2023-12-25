package domain

const (
	ErrIDScanFailed       = Err("Failed to scan id")
	ErrBasketUpdateFailed = Err("Failed to update basket")
	ErrDeleteItemFailed   = Err("Failed to delete Item")
)

type Err string

func (e Err) Error() string {
	return string(e)
}
