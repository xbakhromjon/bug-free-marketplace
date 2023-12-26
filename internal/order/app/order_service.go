package app

import "golang-project-template/internal/order/domain"

type OrderService interface {
	GetOrderByID(orderID int) (domain.Order, error)
	GetAllOrders(page int) ([]domain.Order, error)
	MakeOrderReady(OrderID int) error
	MakeOrderPaid(OrderID int) error
	MakeOrderCancelled(OrderID int) error
}

func NewOrderService(repository domain.OrderRepository) OrderService {
	return &orderService{
		repository: repository,
	}
}

type orderService struct {
	repository domain.OrderRepository
}

func (o *orderService) GetOrderByID(orderID int) (domain.Order, error) {
	return o.repository.GetOrderByID(orderID)
}

func (o *orderService) GetAllOrders(page int) ([]domain.Order, error) {

	pageSize := 10

	offset := (page - 1) * pageSize

	return o.repository.GetPaginatedOrders(offset, pageSize)
}

func (o *orderService) MakeOrderReady(OrderID int) error {
	newStatus := domain.OrderReady
	return o.repository.UpdateStatusOrder(OrderID, newStatus)
}

func (o *orderService) MakeOrderPaid(OrderID int) error {
	newStatus := domain.OrderPaid
	return o.repository.UpdateStatusOrder(OrderID, newStatus)
}

func (o *orderService) MakeOrderCancelled(OrderID int) error {
	newStatus := domain.OrderCancelled
	return o.repository.UpdateStatusOrder(OrderID, newStatus)
}
