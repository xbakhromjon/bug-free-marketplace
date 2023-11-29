package app

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"golang-project-template/internal/shop/domain"
	"testing"
)

type MockProductRepository struct {
	SaveFunc func(product *domain.Product) (int, error)
}

func (m *MockProductRepository) GetById(id int) (*domain.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockProductRepository) GetAllByShopId(shopId int) ([]*domain.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockProductRepository) FindById(id int) (*domain.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockProductRepository) FindAllByShopId(shopId int) ([]*domain.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockProductRepository) Save(product *domain.Product) (int, error) {
	if m.SaveFunc != nil {
		return m.SaveFunc(product)
	}
	return 0, errors.New("SaveFunc is implemented")
}

func TestNewProductUseCase(t *testing.T) {
	newProduct := domain.NewProduct{
		Name:   "Test Product",
		Price:  100,
		ShopId: 1,
	}

	testCase := []struct {
		name          string
		mockSaveFunc  func(product *domain.Product) (int, error)
		expectedId    int
		expectedError error
	}{{
		name: "Successfully save",
		mockSaveFunc: func(product *domain.Product) (int, error) {
			return 1, nil
		},
	},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := &MockProductRepository{
				SaveFunc: tc.mockSaveFunc,
			}
			useCase := NewProductUseCase(mockRepo)

			id, err := useCase.Add(newProduct)

			assert.Equal(t, tc.expectedId, id, "ID mismatch")
			assert.Equal(t, tc.expectedError, err, "Error mismatch")
		})
	}
}
