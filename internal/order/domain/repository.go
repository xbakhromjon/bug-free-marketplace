package domain

type BasketRepository interface {
	CreateBasket(userId int) (int, error)
}

type OrderRepository interface {
	CreateOrder(order Order) error
	GetOrderByID(orderID int) (Order, error)
	GetAllOrders() ([]Order, error)
	UpdateStatusOrder(OrderID int, newStatus string) error
}
