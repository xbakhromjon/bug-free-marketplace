package domain

import (
	"time"
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
