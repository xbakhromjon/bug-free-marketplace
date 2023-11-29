package app

import (
	"errors"
	"golang-project-template/internal/shop/domain"
	"testing"
)

func (m *MockShopRepository) Save(shop domain.NewShop) (int, error) {
	return m.SaveFunc(shop)
}

type MockShopRepository struct {
	SaveFunc func(shop domain.NewShop) (int, error)
}

func TestCreateShop(t *testing.T) {

	mockRepo := &MockShopRepository{}
	service := NewShopService(mockRepo)

	t.Run("Create Shop successfully", func(t *testing.T) {
		expectedID := 1
		mockRepo.SaveFunc = func(shop domain.NewShop) (int, error) {
			return expectedID, nil
		}

		createdID, err := service.Create(domain.NewShop{Name: "Test Shop", OwnerId: 1})

		if err != nil {
			t.Errorf("Expected shop ID %d, but got %d", expectedID, createdID)
		}

	})
	t.Run("Error saving shop", func(t *testing.T) {
		expectedErr := errors.New("mocked error")
		mockRepo.SaveFunc = func(shop domain.NewShop) (int, error) {
			return 0, expectedErr
		}

		_, err := service.Create(domain.NewShop{Name: "Test Shop", OwnerId: 1})

		if err == nil {
			t.Error("Expected an error, but got nil")
		} else if err != expectedErr {
			t.Errorf("Expected error %v, but got %v", expectedErr, err)
		}
	})
}
