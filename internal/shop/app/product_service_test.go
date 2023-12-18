package app

import (
	"errors"
	"github.com/stretchr/testify/assert"
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

func TestNewProduct(t *testing.T) {
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
	}}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := &mockProductRepo{
				SaveFunc: tc.mockSaveFunc,
			}
			useCase := NewProductService(mockRepo)

			id, err := useCase.Add(newProduct)

			assert.Equal(t, tc.expectedId, id, "Id mismatch")
			assert.Equal(t, tc.expectedError, err, "Error mismatch")
		})
	}
}

type mockProductRepo struct {
	SaveFunc func(product *domain.Product) (int, error)
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
	if m.SaveFunc != nil {
		return m.SaveFunc(product)
	}
	return 0, errors.New("Save func is implemented")
}
func (m *mockProductRepo) UpdateProduct(productID int, product *domain.Product) (*domain.Product, error) {
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
