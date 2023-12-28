package domain

type OrderRepository interface {
	CreateOrder(order Order) error
	UpdateStatusOrder(orderID int, newStatus string) error
	GetOrderByID(orderID int) (Order, error)
	GetPaginatedOrders(offset, limit int) ([]Order, error)
}
