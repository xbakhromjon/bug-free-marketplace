package app

import (
	"golang-project-template/internal/order/domain"
	"golang-project-template/internal/shop/app"
	"log"
	"math/rand"
	"time"
)

type OrderService interface {
	CreateOrder(basketID int) error
	GetOrderByID(orderID int) (domain.Order, error)
	GetAllOrders(page int) ([]domain.Order, error)
	MakeOrderReady(OrderID int) error
}

type orderService struct {
	repo           domain.OrderRepository
	basketService  BasketService
	productService app.ProductService
}

func NewOrderService(repo domain.OrderRepository, basketservice BasketService, productservice app.ProductService) OrderService {
	return &orderService{
		repo:           repo,
		basketService:  basketservice,
		productService: productservice,
	}
}

func (o *orderService) CreateOrder(basketID int) error {
	basketItems, err := o.basketService.GetAll(basketID)
	if err != nil {
		return err
	}

	var totalAmount int

	for _, item := range basketItems {

		product, err := o.productService.GetOne(item.ProductId)
		if err != nil {
			return err
		}

		// Calculate total order amount
		totalAmount += product.Price * item.Quantity

	}

	status := domain.OrderInProcess

	number := string(rand.Intn(100000))

	order := domain.Order{
		Number:     number,
		BasketID:   basketID,
		TotalPrice: totalAmount,
		Status:     status,
		CreatedAt:  time.Now(),
	}

	err = o.repo.CreateOrder(order)
	if err != nil {
		return err
	}
	return nil
}

func (o *orderService) GetOrderByID(orderID int) (domain.Order, error) {

	order, err := o.repo.GetOrderByID(orderID)
	if err != nil {
		log.Println("error in GetOrderByID(): ", err.Error())
		return domain.Order{}, err
	}

	return order, nil

}

func (o *orderService) GetAllOrders(page int) ([]domain.Order, error) {

	pageSize := 10

	offset := (page - 1) * pageSize
	orders, err := o.repo.GetPaginatedOrders(offset, pageSize)

	if err != nil {
		log.Println("error in GetAllOrders(): ", err.Error())
		return []domain.Order{}, err
	}

	return orders, nil
}

func (o *orderService) MakeOrderReady(OrderID int) error {
	newStatus := domain.OrderReady
	err := o.repo.UpdateStatusOrder(OrderID, newStatus)
	if err != nil {
		log.Println("error in MakeOrderReady()")
		return err
	}
	return nil
}

func (o *orderService) MakeOrderPaid(OrderID int) error {
	newStatus := domain.OrderPaid
	err := o.repo.UpdateStatusOrder(OrderID, newStatus)
	if err != nil {
		log.Println("error in MakeOrderPaid()")
		return err
	}
	return nil
}

func (o *orderService) MakeOrderCancelled(OrderID int) error {
	newStatus := domain.OrderCancelled
	err := o.repo.UpdateStatusOrder(OrderID, newStatus)
	if err != nil {
		log.Println("error in MakeOrderCancelled()")
		return err
	}
	return nil
}
