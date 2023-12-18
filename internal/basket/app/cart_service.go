package app

import (
	basket "golang-project-template/internal/basket/domain"
	"golang-project-template/internal/shop/domain"
)

type CartService interface {
	CreateCartItem(req basket.CartItems) error
	CreateCart(userID int) (int, error)
	GetBasket(userID int) (*basket.Cart, error)
	AddProductToCart(userID, productID, quantity int) (*basket.Cart, error)
	DeleteProductFromCart(userId, productId int) error
	IncrementProductQuantity(userId, productId int) error
	DecrementProductQuantity(userId, productId int) error
}

func NewCartService(repo basket.CartRepository) CartService {
	return CartServiceImpl{cartRepo: repo}
}

type CartServiceImpl struct {
	cartRepo    basket.CartRepository
	productRepo domain.ProductRepository
}

func (cs CartServiceImpl) CreateCart(userID int) (int, error) {
	cart := &basket.Cart{
		UserId: userID,
	}
	cartID, err := cs.cartRepo.Create(cart)
	if err != nil {
		return 0, basket.ErrCartCreationFailed
	}
	return cartID, nil
}

func (cs CartServiceImpl) CreateCartItem(req basket.CartItems) error {
	cart := &basket.CartItems{
		Id:        req.Id,
		CartId:    req.CartId,
		ProductId: req.ProductId,
		Quantity:  req.Quantity,
	}
	_, err := cs.cartRepo.CreateCardItem(cart)
	if err != nil {
		return basket.ErrCartItemCreationFailed
	}
	return nil
}

func (cs CartServiceImpl) GetBasket(userID int) (*basket.Cart, error) {
	cart, err := cs.cartRepo.GetByUserId(userID)
	if err != nil {
		return nil, basket.ErrCartNotFound
	}
	return cart, nil
}

func (cs CartServiceImpl) UpdateProductQuantity(userId, productId, quantity int) error {
	cart, err := cs.cartRepo.GetByUserId(userId)
	if err != nil {
		return basket.ErrCartNotFound
	}
	cardItem, err := cs.cartRepo.GetCardItem(cart.Id, productId)
	if err != nil {
		return basket.ErrProductNotFound
	}

	newQuantity := cardItem.Quantity + quantity
	if newQuantity < 0 {
		newQuantity = 0
	}
	cardItem.Quantity = newQuantity
	err = cs.cartRepo.UpdateCartItem(userId, productId, cardItem.Quantity)
	if err != nil {
		return basket.ErrCartUpdateFailed
	}
	return nil
}

func (cs CartServiceImpl) AddProductToCart(userID, productID, quantity int) (*basket.Cart, error) {
	cart, err := cs.cartRepo.GetByUserId(userID)
	if err != nil {
		cart = &basket.Cart{
			UserId: userID,
		}
		_, err := cs.cartRepo.Create(cart)
		if err != nil {
			return nil, basket.ErrCartNotFound
		}
	}
	_, err = cs.cartRepo.GetCart(cart.Id)
	if err != nil {
		cartItem := &basket.CartItems{
			CartId:    cart.Id,
			ProductId: productID,
			Quantity:  quantity,
		}
		_, err := cs.cartRepo.CreateCardItem(cartItem)
		if err != nil {
			return nil, basket.ErrCartItemCreationFailed
		}
	}
	updatedCart, err := cs.cartRepo.GetByUserId(userID)
	if err != nil {
		return nil, err
	}
	return updatedCart, nil
}

func (cs CartServiceImpl) IncrementProductQuantity(userId, productId int) error {
	return cs.UpdateProductQuantity(userId, productId, +1)
}

func (cs CartServiceImpl) DecrementProductQuantity(userId, productId int) error {
	return cs.UpdateProductQuantity(userId, productId, -1)
}

func (cs CartServiceImpl) DeleteProductFromCart(userId, productId int) error {
	cart, err := cs.cartRepo.GetByUserId(userId)
	if err != nil {
		return basket.ErrCartNotFound
	}
	_, err = cs.cartRepo.GetCardItem(cart.Id, productId)
	if err != nil {
		return basket.ErrProductNotFound
	}
	return cs.cartRepo.DeleteProduct(cart.Id, productId)
}
