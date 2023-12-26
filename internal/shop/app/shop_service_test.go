package app

import (
	"golang-project-template/internal/shop/domain"
	"reflect"
	"testing"
	"time"
)

type mockShopRepo struct {
	f domain.ShopFactory
}

func newMockShopRepo() domain.ShopRepository {
	return &mockShopRepo{}
}

func (m *mockShopRepo) Save(shop *domain.Shop) (int, error) {
	if shop.GetName() == "" {
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

func (m *mockShopRepo) FindAllShops(limit, offset int, search string) ([]domain.Shop, error) {

	if limit == 1 && offset == 1 && search == "" {
		return []domain.Shop{
			newValidShop(),
		}, nil
	}

	return []domain.Shop{}, nil
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

	service := shopService{
		repository:     newMockShopRepo(),
		shopFactory:    domain.NewShopFactory(20, 10),
		userRepository: mockUserRepo{},
	}

	t.Run("Create Shop successfully", func(t *testing.T) {
		reqShop := NewShop{"testing shop name", 1}
		got, err := service.Create(reqShop)
		want := 1
		if err != nil {
			t.Errorf("Error expected to be nil, bot got %v", err)
		} else if got != want {
			t.Errorf("want %v, but got %v", want, got)
		}

	})

	cases := []struct {
		label     string
		newShop   NewShop
		wantedErr domain.Err
	}{
		{
			"empty shop name",
			NewShop{"", 1},
			domain.ErrEmptyShopName,
		},
		{
			"shop name exists",
			NewShop{"existing shop", 1},
			domain.ErrShopNameExists,
		},
		{
			"invalid shop name",
			NewShop{"shop name that contains more than 20 chars", 1},
			domain.ErrInvalidShopName,
		},
		{
			"no such user",
			NewShop{"random shop name", 1},
			domain.ErrUserNotExists,
		},
	}
	for _, test := range cases {
		t.Run(test.label, func(t *testing.T) {
			_, gotErr := service.Create(test.newShop)
			if gotErr == nil {
				t.Error("Expected error but got nil")
			} else if gotErr != test.wantedErr {
				t.Errorf("Expected %v, but got %v", test.wantedErr, gotErr)
			}
		})
	}
}

func TestGetShopById(t *testing.T) {
	service := shopService{
		repository:     newMockShopRepo(),
		userRepository: mockUserRepo{},
	}

	t.Run("Get Shop By correct Id", func(t *testing.T) {
		got, err := service.GetShopById(1)
		want := newValidShop()
		if err != nil {
			t.Errorf("Error didn`t expected, but got %v", err)
		} else if !reflect.DeepEqual(want, got) {
			t.Errorf("want %v, but got %v", want, got)
		}

	})

	t.Run("Get Shop By incorrect Id", func(t *testing.T) {
		_, err := service.GetShopById(2)
		want := domain.ErrShopNotFound
		if err == nil {
			t.Error("Expected err, but but didn`t get")
		} else if err != want {
			t.Errorf("want %q, but got %q", want, err)
		}

	})
}

func TestGetAllShops(t *testing.T) {
	service := shopService{
		repository:     newMockShopRepo(),
		shopFactory:    domain.NewShopFactory(20, 10),
		userRepository: mockUserRepo{},
	}

	t.Run("GetAllShops successfully", func(t *testing.T) {
		want := []domain.Shop{
			newValidShop(),
		}
		got, err := service.GetAllShops(1, 1, "")

		if err != nil {
			t.Errorf("Didn`t expect error, but got %q", err)
		} else if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %v, but got %v", want, got)
		}
	})

	cases := []struct {
		label     string
		limit     int
		offset    int
		search    string
		wantedErr domain.Err
	}{
		{"invalid limit", 101, 1, "something", domain.ErrInvalidLimit},
		{"invalid offset", 28, -2, "something", domain.ErrInvalidOffset},
		{"invalid search", 28, 1, "long search string", domain.ErrInvalidSearch},
	}

	for _, test := range cases {
		t.Run(test.label, func(t *testing.T) {
			_, err := service.GetAllShops(test.limit, test.offset, test.search)

			if err == nil {
				t.Error("Expected error, but didn`t get")
			} else if err != test.wantedErr {
				t.Errorf("Expected %v, but got %v", test.wantedErr, err)
			}
		})
	}

}

func newValidShop() domain.Shop {
	res := domain.Shop{}
	res.SetId(1)
	res.SetName("Default")
	res.SetOwnerId(1)
	res.SetCreateAt(time.Now())
	res.SetUpdatedAt(time.Now())
	return res

}
