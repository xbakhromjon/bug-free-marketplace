package app

import (
	"errors"
	domain2 "golang-project-template/internal/order/domain"
	"golang-project-template/internal/shop/domain"
)

type CartService interface {
	Create(req domain2.CartItems) error
	AddProductToCart(userID, productID, quantity int) error
	DeleteProductFromCart(userId, productId int) error
}

func NewCartService(repo domain2.CartRepository) CartService {
	return CartServiceImpl{cartRepo: repo}
}

type CartServiceImpl struct {
	cartRepo    domain2.CartRepository
	productRepo domain.ProductRepository
}

func (cs CartServiceImpl) DeleteProductFromCart(userId, productId int) error {
	cart, err := cs.cartRepo.GetByUserId(userId)
	if err != nil {
		return err
	}

	_, err = cs.cartRepo.GetCardItem(cart.Id, productId)
	if err != nil {
		return errors.New("Product not found in the basket")
	}
	return cs.cartRepo.DeleteProduct(cart.Id, productId)
}

func (cs CartServiceImpl) AddProductToCart(userID, productID, quantity int) error {
	cart, err := cs.cartRepo.GetByUserId(userID)
	if err != nil {
		cart = &domain2.Cart{
			UserId: userID,
		}
		err := cs.cartRepo.Create(cart)
		if err != nil {
			return err
		}
	}

	cardItem, err := cs.cartRepo.GetCardItem(cart.Id, productID)
	if err != nil {
		cardItem = &domain2.CartItems{
			CartId:    cart.Id,
			ProductId: productID,
			Quantity:  quantity,
		}
		_, err := cs.cartRepo.CreateCardItem(cardItem)
		return err
	}

	// Case: if product already exixts in basket, update the quantity
	cardItem.Quantity += quantity
	err = cs.cartRepo.UpdateCartItem(userID, productID, cardItem.Quantity)
	return err
}

func (cs CartServiceImpl) Create(req domain2.CartItems) error {
	cart := &domain2.CartItems{
		Id:        req.Id,
		CartId:    req.CartId,
		ProductId: req.ProductId,
		Quantity:  req.Quantity,
	}
	_, err := cs.cartRepo.CreateCardItem(cart)
	if err != nil {
		return err
	}
	return nil
}
