package domain

import "errors"

var (
	ErrOrderNotFound = errors.New("order not found")
	OrderWaiting     = "Waiting"
	OrderTaken       = "Taken"
)
