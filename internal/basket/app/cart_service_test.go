package app

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"golang-project-template/internal/basket/domain"
	"testing"
)

func TestCartService_CreateBasket(t *testing.T) {
	underTest := CartServiceImpl{cartRepo: &mockCartRepo{}}

	cart := mockCartForCreate(1)
	got, _ := underTest.CreateBasket(cart.UserId)
	want := 1

	assert.Equal(t, got, want)

	testCases := []struct {
		name          string
		cart          *domain.NewCart
		expectedError error
	}{
		{
			name:          "failed to create cart(Invalid id)",
			cart:          mockCreateInvalidId(0),
			expectedError: domain.ErrInvalidId,
		},
	}

	for _, tCase := range testCases {
		t.Run(tCase.name, func(t *testing.T) {
			_, err := underTest.CreateBasket(tCase.cart.UserId)
			assert.Equal(t, err, tCase.expectedError)
		})
	}
}

func TestCartServiceImpl_AddItem(t *testing.T) {
	underTest := CartServiceImpl{cartRepo: &mockCartRepo{}}

	cItems, _ := mockAddItem(1, 1, 100)
	got, _ := underTest.AddItem(cItems)
	want := 1
	assert.Equal(t, got, want)

}

func TestCartServiceImpl_GetAll(t *testing.T) {
	underTest := CartServiceImpl{cartRepo: &mockCartRepo{}}

	cartId := 1
	got, err := underTest.GetAll(cartId)
	want := mockGetAll(1)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	assert.Equal(t, got, want)

	t.Run("Error Scan Failed id", func(t *testing.T) {
		cartId := 0
		mockRepo := &mockCartRepo{}
		mockRepo.getAllItems, mockRepo.getAllErr = mockGetAll(cartId), errors.New("Error Scan Failed Id")
		underTest.cartRepo = mockRepo

		got, err := underTest.GetAll(cartId)
		assert.ErrorIs(t, err, domain.ErrIDScanFailed)
		assert.Nil(t, got)
	})

	t.Run("Empty Result", func(t *testing.T) {
		cartId := 0
		mockRepo := &mockCartRepo{}

		mockRepo.getAllItems = []domain.CartItems{}
		underTest.cartRepo = mockRepo

		got, err := underTest.GetAll(cartId)

		assert.NoError(t, err)
		assert.Empty(t, got)
	})
}

func TestCartServiceImpl_UpdateProductQuantity(t *testing.T) {
	underTest := CartServiceImpl{cartRepo: &mockCartRepo{}}

	got := underTest.UpdateProductQuantity(1, 10)
	want := mockUpdateItems(1, 10)
	assert.Equal(t, got, want)

	t.Run("Error Case - Invalid Cart ID", func(t *testing.T) {
		invalidCartId := 0
		quantity := 5
		mockRepo := &mockCartRepo{}
		underTest.cartRepo = mockRepo

		err := underTest.UpdateProductQuantity(invalidCartId, quantity)

		assert.ErrorIs(t, err, domain.ErrInvalidCartId)
	})

}

func mockGetAll(cartId int) []domain.CartItems {
	items := []domain.CartItems{
		{
			Id:        1,
			CartId:    cartId,
			ProductId: 1,
			Quantity:  10,
		},
	}
	return items
}

func mockUpdateItems(cartId, quantity int) error {
	updatedCart := &domain.CartItems{}
	updatedCart.Id = 1
	updatedCart.CartId = cartId
	updatedCart.ProductId = 1
	updatedCart.Quantity = quantity

	return nil
}

func mockAddItem(cartId, productId, quantity int) (*domain.CartItems, error) {
	newCart := &domain.CartItems{}
	newCart.CartId = cartId
	newCart.ProductId = productId
	newCart.Quantity = quantity

	if newCart.Quantity <= 0 {
		return nil, domain.ErrItemCreationFailed
	}

	return newCart, nil
}

func mockCreateInvalidId(id int) *domain.NewCart {
	newCart := &domain.NewCart{}
	newCart.UserId = id

	return newCart
}

func mockCartForCreate(UserId int) *domain.NewCart {
	newCart := &domain.NewCart{}
	newCart.UserId = UserId
	return newCart
}

type mockCartRepo struct {
	getAllItems []domain.CartItems
	getAllErr   error
}

func (m *mockCartRepo) CreateBasket(userId int) (int, error) {
	if userId <= 0 {
		return 0, domain.ErrInvalidId
	}
	return 1, nil
}

func (m *mockCartRepo) AddItem(cart *domain.CartItems) (int, error) {
	return 1, nil
}

func (m *mockCartRepo) GetAll(cartId int) ([]domain.CartItems, error) {
	if cartId == 0 {
		return nil, errors.New(("Error Scan Failed Id"))
	}
	return []domain.CartItems{{Id: 1, CartId: 1, ProductId: 1, Quantity: 10}}, nil
}

func (m *mockCartRepo) UpdateCartItem(cartId, quantity int) error {
	if cartId <= 0 {
		return domain.ErrInvalidCartId
	}
	return nil
}

func (m *mockCartRepo) DeleteProduct(cartId, productId int) (id int, err error) {
	return 1, nil
}

func NewMockCartRepo() domain.CartRepository {
	return &mockCartRepo{}
}
