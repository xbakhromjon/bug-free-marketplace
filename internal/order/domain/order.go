package domain

import (
	"errors"
	"time"
)

var (
	ErrOrderNotFound = errors.New("Order not found")
)

const (
	OrderInProcess = "ORDER_IN_PROCCES"
	OrderReady     = "ORDER_READY"
	OrderPaid      = "ORDER_PAID"
	OrderCancelled = "ORDER_CANCELLED"
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
