package app

import (
	basketapp "golang-project-template/internal/basket/app"
	"golang-project-template/internal/order/domain"
	productapp "golang-project-template/internal/shop/app"
	"math/rand"
	"time"
)

type OrderService interface {
	CreateOrder(order domain.Order) error
	GetOrderByID(orderID int) (domain.Order, error)
	GetAllOrders() ([]domain.Order, error)
	UpdateStatus(UserID int) error
}

func NewOrderService(repository domain.OrderRepository, basketUsecase basketapp.CartService, productUsecase productapp.ProductService) *orderService {
	return &orderService{
		repository:     repository,
		basketUseCase:  basketUsecase,
		productUsecase: productUsecase,
	}
}

type orderService struct {
	repository     domain.OrderRepository
	basketUseCase  basketapp.CartService
	productUsecase productapp.ProductService
}

func (o *orderService) CreateOrder(basketId int) error {

	cartItems, err := o.basketUseCase.GetAll(basketId)
	if err != nil {
		return err
	}

	var totalAmount int

	for _, item := range cartItems {
		// Get product details for each item in the basket
		product, err := o.productUsecase.GetOne(item.ProductId)
		if err != nil {
			return err
		}

		// Calculate total order amount
		totalAmount += product.Price * item.Quantity

	}

	status := domain.OrderWaiting

	number := string(rand.Intn(100000))

	order := domain.Order{
		Number:     number,
		CartId:     basketId,
		TotalPrice: totalAmount,
		Status:     status,
		CreatedAt:  time.Now(),
	}

	err = o.repository.CreateOrder(order)
	if err != nil {
		return err
	}
	return nil
}

func (o *orderService) GetOrderByID(orderID int) (domain.Order, error) {
	return o.repository.GetOrderByID(orderID)
}

func (o *orderService) GetAllOrders() ([]domain.Order, error) {
	return o.repository.GetAllOrders()
}

func (o *orderService) UpdateStatus(orderID int, newStatus string) error {
	newStatus = domain.OrderTaken
	return o.repository.UpdateStatus(orderID, newStatus)
}
