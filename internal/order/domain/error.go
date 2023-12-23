package domain

const (
	ErrBasketUpdateFailed = Err("Failed to update basket")
	ErrIDScanFailed       = Err("Failed to scan id")
)

type Err string

func (e Err) Error() string {
	return string(e)
}
