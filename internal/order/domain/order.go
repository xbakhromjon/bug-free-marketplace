package domain

import "time"

type Order struct {
	Id         int
	Number     string
	CartId     int
	TotalPrice int
	Status     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
type OrderRepository interface {
	CreateOrder(order Order) error
	GetOrderByID(orderID int) (Order, error)
	GetAllOrders() ([]Order, error)
	UpdateStatus(orderID int, newStatus string) error
}

type Orders struct {
	repo OrderRepository
}

func NewOrders(repo OrderRepository) *Orders {
	return &Orders{
		repo: repo,
	}
}
