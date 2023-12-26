package domain

type OrderRepository interface {
	CreateOrder(order Order) error
	UpdateStatusOrder(orderID int, newStatus string) error
	GetOrderByID(orderID int) (Order, error)
	GetAllOrders() ([]Order, error)
}
