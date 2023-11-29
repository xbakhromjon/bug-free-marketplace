package app

import (
	"golang-project-template/internal/shop/domain"
	"testing"
)

func (m *mockShopRepo) Save(shop domain.NewShop) (int, error) {
	if shop.Name == "" {
		return 0, domain.ErrInvalidShopName
	}

	return 1, nil
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
	t.Run("invalid shop name", func(t *testing.T) {
		wanted := domain.ErrInvalidShopName

		_, err := underTest.Create(domain.NewShop{Name: "", OwnerId: 1})

		if err == nil {
			t.Error("Expected an error, but got nil")
		} else if err != wanted {
			t.Errorf("Expected error %v, but got %v", wanted, err)
		}
	})
}
