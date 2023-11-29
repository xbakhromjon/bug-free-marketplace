package app

import (
	"errors"
	"golang-project-template/internal/shop/domain"
	"testing"
)

var mockedError = errors.New("mocked error")

func (m *mockShopRepo) Save(shop domain.NewShop) (int, error) {
	if shop.Name == "Test Shop" && shop.OwnerId == 1 {
		return 1, nil
	}

	return 0, mockedError
}

type mockShopRepo struct {
}

func newMockShopRepo() domain.ShopRepository {
	return &mockShopRepo{}
}

func TestCreateShop(t *testing.T) {

	underTest := shopService{repository: newMockShopRepo()}

	t.Run("Create Shop successfully", func(t *testing.T) {
		got, _ := underTest.Create(domain.NewShop{Name: "Test Shop", OwnerId: 1})
		want := 1

		if got != want {
			t.Errorf("want %v, but got %v", want, got)
		}

	})
	t.Run("Error saving shop", func(t *testing.T) {
		wanted := mockedError

		_, err := underTest.Create(domain.NewShop{Name: "failed name", OwnerId: 2})

		if err == nil {
			t.Error("Expected an error, but got nil")
		} else if err != wanted {
			t.Errorf("Expected error %v, but got %v", wanted, err)
		}
	})
}
