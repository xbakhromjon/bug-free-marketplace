package app

import (
	"fmt"
	"golang-project-template/internal/order/domain"
)

type CartItemService interface {
	Add(productId int, userId int, quantity int) error
	Remove(productId int, userId int) error
	GetAll(userId int) ([]*domain.CartItems, error)
	RemoveAll(userId int) error
}

func NewCartItemService(repository domain.CartItemsRepository) CartItemService {

	return &cartItemService{repository: repository}
}

type cartItemService struct {
	repository domain.CartItemsRepository
}

func (c *cartItemService) Add(productId int, userId int, quantity int) error {
	err := c.repository.AddItem(productId, userId, quantity)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func (c *cartItemService) Remove(productId int, userId int) error {
	err := c.repository.RemoveItem(productId, userId)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
func (c *cartItemService) GetAll(userId int) ([]*domain.CartItems, error) {
	Items, err := c.repository.GetAll(userId)
	if err != nil {
		fmt.Println(err)
	}
	return Items, nil
}
func (c *cartItemService) RemoveAll(userId int) error {
	err := c.repository.RemoveAll(userId)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
