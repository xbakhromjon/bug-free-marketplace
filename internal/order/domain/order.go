package domain

import (
	"errors"
	"time"
)

var (
	ErrOrderNotFound = errors.New("Order not found")
)

type Order struct {
	ID         int
	Number     string
	BasketID   int
	TotalPrice int
	Status     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
