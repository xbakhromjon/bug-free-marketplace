package app

import (
	basket "golang-project-template/internal/basket/domain"
	"golang-project-template/internal/shop/domain"
)

type CartService interface {
	CreateCart(userID int) (int, error)
	GetBasket(userID int) (*basket.Cart, error)
	GetCartItem(cartId int) ([]*basket.CartItems, error)
	AddProductToCart(userID, productID, quantity int) (*basket.Cart, error)
	UpdateProductQuantity(cartId, quantity int) error
	DeleteProductFromCart(userId, productId int) error
	IncrementProductQuantity(cartId int) error
	DecrementProductQuantity(cartId int) error
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

func (cs CartServiceImpl) GetCartItem(cartId int) ([]*basket.CartItems, error) {
	cItems, err := cs.cartRepo.GetCardItem(cartId)
	if err != nil {
		return nil, basket.ErrCartItemNotFound
	}
	return cItems, nil
}

func (cs CartServiceImpl) UpdateProductQuantity(cartId, quantity int) error {
	err := cs.cartRepo.UpdateCartItem(cartId, quantity)
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

func (cs CartServiceImpl) IncrementProductQuantity(cartId int) error {
	return cs.UpdateProductQuantity(cartId, +1)
}

func (cs CartServiceImpl) DecrementProductQuantity(cartId int) error {
	return cs.UpdateProductQuantity(cartId, -1)
}

func (cs CartServiceImpl) DeleteProductFromCart(cartId, productId int) error {
	_, err := cs.cartRepo.GetCartItemByCartIdAndProductId(cartId, productId)
	if err != nil {
		return basket.ErrProductNotFound
	}
	return cs.cartRepo.DeleteProduct(cartId, productId)
}
