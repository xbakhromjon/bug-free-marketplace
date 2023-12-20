package app

import (
	basket "golang-project-template/internal/basket/domain"
	"golang-project-template/internal/shop/domain"
)

type CartService interface {
	CreateBasket(userID int) (int, error)
	AddItem(cart *basket.CartItems) (int, error)
	GetAll(cartId int) ([]*basket.CartItems, error)
	UpdateProductQuantity(cartId, quantity int) error
	DeleteProductFromCart(cartId, productId int) (id int, err error)
}

func NewCartService(repo basket.CartRepository) CartService {
	return CartServiceImpl{cartRepo: repo}
}

type CartServiceImpl struct {
	cartRepo    basket.CartRepository
	productRepo domain.ProductRepository
}

func (cs CartServiceImpl) CreateBasket(userID int) (int, error) {

	cartId, err := cs.cartRepo.CreateBasket(userID)
	if err != nil {
		return 0, basket.ErrIDScanFailed
	}
	return cartId, nil
}
func (cs CartServiceImpl) AddItem(cart *basket.CartItems) (int, error) {
	id, err := cs.cartRepo.AddItem(cart)
	if err != nil {
		return 0, basket.ErrAddItemFailed
	}
	return id, nil
}
func (cs CartServiceImpl) GetAll(cartId int) ([]*basket.CartItems, error) {
	cItems, err := cs.cartRepo.GetAll(cartId)
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

func (cs CartServiceImpl) DeleteProductFromCart(cartId, productId int) (id int, err error) {
	return cs.cartRepo.DeleteProduct(cartId, productId)
}
