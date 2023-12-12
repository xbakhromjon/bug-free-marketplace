package app

import (
	"golang-project-template/internal/shop/domain"
	"testing"
	"time"
)

type mockShopRepo struct {
}

func newMockShopRepo() domain.ShopRepository {
	return &mockShopRepo{}
}

func (m *mockShopRepo) Save(shop domain.NewShop) (int, error) {
	if shop.Name == "" {
		return 0, domain.ErrEmptyShopName
	}

	return 1, nil
}

func (m *mockShopRepo) CheckShopNameExists(shopName string) (bool, error) {
	existingShopName := "existing_shop"

	if existingShopName == shopName {
		return true, nil
	}

	return false, nil
}

func (m *mockShopRepo) FindShopById(id int) (domain.Shop, error) {
	if id == 1 {
		return newValidShop(), nil
	}
	return domain.Shop{}, domain.ErrShopNotFound
}

type mockUserRepo struct {
}

func (u mockUserRepo) UserExists(id int) (bool, error) {
	if id == 99 {
		return false, domain.ErrUserNotExists
	}
	return true, nil
}

func TestCreateShop(t *testing.T) {

	underTest := shopService{
		repository:     newMockShopRepo(),
		shopFactory:    domain.NewShopFactory(20),
		userRepository: mockUserRepo{},
	}

	t.Run("Create Shop successfully", func(t *testing.T) {
		got, err := underTest.Create(
			domain.NewShop{Name: "testing shop name", OwnerId: 1},
		)
		want := 1
		if err != nil {
			t.Errorf("Error expected to be nil, bot got %v", err)
		} else if got != want {
			t.Errorf("want %v, but got %v", want, got)
		}

	})

	cases := []struct {
		label     string
		newShop   domain.NewShop
		wantedErr domain.Err
	}{
		{
			"empty shop name",
			domain.NewShop{
				Name:    "",
				OwnerId: 1,
			},
			domain.ErrEmptyShopName,
		},
		{
			"shop name exists",
			domain.NewShop{
				Name:    "existing_shop",
				OwnerId: 1,
			},
			domain.ErrShopNameExists,
		},
		{
			"invalid shop name",
			domain.NewShop{
				Name:    "shop name that contains more than 20 chars",
				OwnerId: 1,
			},
			domain.ErrInvalidShopName,
		},
		{
			"no such user",
			domain.NewShop{
				Name:    "random shop name",
				OwnerId: 99,
			},
			domain.ErrUserNotExists,
		},
	}
	for _, test := range cases {
		t.Run(test.label, func(t *testing.T) {
			_, gotErr := underTest.Create(test.newShop)
			if gotErr == nil {
				t.Error("Expected error but got nil")
			} else if gotErr != test.wantedErr {
				t.Errorf("Expected %v, but got %v", test.wantedErr, gotErr)
			}
		})
	}
}

func TestGetShopById(t *testing.T) {
	underTest := shopService{
		repository:     newMockShopRepo(),
		shopFactory:    domain.NewShopFactory(20),
		userRepository: mockUserRepo{},
	}

	t.Run("Get Shop By correct Id", func(t *testing.T) {
		got, err := underTest.GetShopById(1)
		want := newValidShop()
		if err != nil {
			t.Errorf("Error didn`t expected, but got %v", err)
		} else if got != want {
			t.Errorf("want %v, but got %v", want, got)
		}

	})

	t.Run("Get Shop By incorrect Id", func(t *testing.T) {
		_, err := underTest.GetShopById(2)
		want := domain.ErrShopNotFound
		if err == nil {
			t.Error("Expected err, but but didn`t get")
		} else if err != want {
			t.Errorf("want %q, but got %q", want, err)
		}

	})
}
func newValidShop() domain.Shop {
	return domain.Shop{
		Id:        1,
		Name:      "Default shop name",
		OwnerId:   1,
		CreatedAt: time.Now().Format(time.RFC1123),
		UpdatedAt: time.Now().Format(time.RFC1123),
	}
}
