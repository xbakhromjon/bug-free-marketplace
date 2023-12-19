package app

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"golang-project-template/internal/basket/domain"
	"testing"
)

type MockCartRepository struct {
	CreateCartFunc                      func(cart *domain.Cart) (int, error)
	CreateCartItemFunc                  func(cItems *domain.CartItems) (int, error)
	GetByUserIdFunc                     func(userID int) (*domain.Cart, error)
	GetCardItemFunc                     func(cartId int) (*domain.CartItems, error)
	GetCartByIdFunc                     func(id int) (*domain.Cart, error)
	GetCartItemByCartIdAndProductIdFunc func(cartID, productID int) (*domain.CartItems, error)
	UpdateCartItemFunc                  func(cartId, productId, quantity int) error
}

func newMockCartRepo() domain.CartRepository {
	return &MockCartRepository{}
}

func (m *MockCartRepository) Create(cart *domain.Cart) (int, error) {
	return m.CreateCartFunc(cart)
}

func (m *MockCartRepository) CreateCardItem(cItems *domain.CartItems) (int, error) {
	return m.CreateCartItemFunc(cItems)
}

func (m *MockCartRepository) GetByUserId(userID int) (*domain.Cart, error) {
	return m.GetByUserIdFunc(userID)
}

func (m *MockCartRepository) GetCardItem(cartId int) (*domain.CartItems, error) {
	return m.GetCardItemFunc(cartId)
}

func (m *MockCartRepository) GetCart(id int) (*domain.Cart, error) {
	return m.GetCartByIdFunc(id)
}

func (m *MockCartRepository) GetCartItemByCartIdAndProductId(cartId, productId int) (*domain.CartItems, error) {
	return m.GetCartItemByCartIdAndProductIdFunc(cartId, productId)
}

func (m *MockCartRepository) UpdateCartItem(cartId, productId, quantity int) error {
	return m.UpdateCartItemFunc(cartId, productId, quantity)
}

func (m *MockCartRepository) DeleteProduct(cardId, productId int) error {
	//TODO implement me
	panic("implement me")
}

func TestCartServiceImpl_CreateCart(t *testing.T) {
	mockRepo := &MockCartRepository{
		CreateCartFunc: func(cart *domain.Cart) (int, error) {
			return 1, nil
		},
	}
	service := NewCartService(mockRepo)

	userID := 12
	cartID, err := service.CreateCart(userID)

	assert.NoError(t, err)
	assert.Equal(t, 12, cartID)
}

func TestCartServiceImpl_CreateCartItem(t *testing.T) {

	mockRepo := &MockCartRepository{
		CreateCartItemFunc: func(cItems *domain.CartItems) (int, error) {
			return 1, nil
		},
	}
	service := NewCartService(mockRepo)

	request := domain.CartItems{
		Id:        1,
		CartId:    1,
		ProductId: 1,
		Quantity:  1,
	}

	err := service.CreateCartItem(request)

	assert.NoError(t, err)

	t.Run("error in cart item creation", func(t *testing.T) {
		expectedErr := errors.New("sasdssf")
		mockRepo := &MockCartRepository{
			CreateCartItemFunc: func(cItems *domain.CartItems) (int, error) {
				return 1, expectedErr
			},
		}
		service := NewCartService(mockRepo)

		request := domain.CartItems{
			Id:        1,
			CartId:    1,
			ProductId: 1,
			Quantity:  10,
		}

		err := service.CreateCartItem(request)
		assert.Error(t, err)
		assert.True(t, errors.Is(err, expectedErr), "error type mismatch")
		assert.Equal(t, expectedErr.Error(), err.Error(), "error message mismatch")
	})
}

func TestCartServiceImpl_GetCartItem(t *testing.T) {
	mockRepo := &MockCartRepository{
		GetByUserIdFunc: func(userID int) (*domain.Cart, error) {
			return &domain.Cart{Id: 1, UserId: userID}, nil
		},
	}
	service := NewCartService(mockRepo)

	userId := 2
	cart, err := service.GetBasket(userId)

	assert.NoError(t, err)
	assert.Equal(t, userId, cart.Id)
}
