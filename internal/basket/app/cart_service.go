package app

import (
	"errors"
	domain2 "golang-project-template/internal/basket/domain"
	"golang-project-template/internal/shop/domain"
)

type CartService interface {
	CreateCartItem(req domain2.CartItems) error
	CreateCart(userID int) (int, error)
	GetBasket(userID int) (*domain2.Cart, error)
	AddProductToCart(userID, productID, quantity int) (*domain2.Cart, error)
	DeleteProductFromCart(userId, productId int) error
}

func NewCartService(repo domain2.CartRepository) CartService {
	return CartServiceImpl{cartRepo: repo}
}

type CartServiceImpl struct {
	cartRepo    domain2.CartRepository
	productRepo domain.ProductRepository
}

func (cs CartServiceImpl) CreateCart(userID int) (int, error) {
	cart := &domain2.Cart{
		UserId: userID,
	}
	cartID, err := cs.cartRepo.Create(cart)
	if err != nil {
		return 0, err
	}
	return cartID, nil
}

func (cs CartServiceImpl) CreateCartItem(req domain2.CartItems) error {
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

func (cs CartServiceImpl) GetBasket(userID int) (*domain2.Cart, error) {
	cart, err := cs.cartRepo.GetByUserId(userID)
	if err != nil {
		return nil, err
	}
	return cart, nil
}

func (cs CartServiceImpl) AddProductToCart(userID, productID, quantity int) (*domain2.Cart, error) {
	cart, err := cs.cartRepo.GetByUserId(userID)
	if err != nil {
		cart = &domain2.Cart{
			UserId: userID,
		}
		_, err := cs.cartRepo.Create(cart)
		if err != nil {
			return nil, err
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
		if err != nil {
			return nil, err
		}
	} else {
		cardItem.Quantity = quantity
		err = cs.cartRepo.UpdateCartItem(userID, productID, quantity)
		if err != nil {
			return nil, err
		}
	}

	updatedCart, err := cs.cartRepo.GetByUserId(userID)
	if err != nil {
		return nil, err
	}
	return updatedCart, nil
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
