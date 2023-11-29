package app

import (
	"golang-project-template/internal/shop/domain"
	"testing"
)

func TestGetOneProduct(t *testing.T) {
	underTest := productService{repository: newMockProductRepo()}
	t.Run("with correct id", func(t *testing.T) {
		p := newValidProduct()
		got, _ := underTest.GetOne(p.Id)
		want := p

		if *got != *want {
			t.Errorf("want %+v but got %+v", want, got)
		}
	})

	t.Run("with invalid id", func(t *testing.T) {
		id := 2
		_, err := underTest.GetOne(id)
		want := domain.ErrProductNotFound

		if err == nil {
			t.Errorf("expected %T, got nil", want)
			return
		}

		if err != domain.ErrProductNotFound {
			t.Errorf("want %T, got %T", want, err)
		}
	})
}

type mockProductRepo struct {
}

func newMockProductRepo() domain.ProductRepository {
	return &mockProductRepo{}
}

func (m *mockProductRepo) FindById(id int) (*domain.Product, error) {
	if id == 1 {
		return newValidProduct(), nil
	}

	return nil, domain.ErrProductNotFound
}

func (m *mockProductRepo) Save(product *domain.Product) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (m *mockProductRepo) FindAllByShopId(shopId int) ([]*domain.Product, error) {
	//TODO implement me
	panic("implement me")
}

func newValidProduct() *domain.Product {
	return &domain.Product{
		Id:     1,
		Name:   "T-shirt",
		Price:  100,
		ShopId: 1,
	}
}
