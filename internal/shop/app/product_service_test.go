package app

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"golang-project-template/internal/common"
	"golang-project-template/internal/shop/domain"
	"reflect"
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

func TestFilter(t *testing.T) {
	underTest := productService{repository: newMockProductRepo()}
	t.Run("result proper with repository returned result", func(t *testing.T) {
		searchModel := underTest.factory.CreateNewSearchModel("T-shirt", 10, 20)
		got, err := underTest.Filter(*searchModel)
		if err != nil {
			t.Errorf("expected ok but %q error occured", err)
		}

		want, _ := underTest.repository.FindAll(*searchModel)

		if !reflect.DeepEqual(want, got) {
			t.Errorf("want %+v but got %+v", want, got)
		}
	})
}

func TestFilterByPageable(t *testing.T) {
	underTest := productService{repository: newMockProductRepo()}
	t.Run("result proper with that repository returned result", func(t *testing.T) {
		searchModel := *underTest.factory.CreateNewSearchModel("T-shirt", 10, 20)
		pageableRequest := *common.CreateDefaultPageableRequest()
		got, err := underTest.FilterByPageable(searchModel, pageableRequest)
		if err != nil {
			t.Errorf("expected ok but %q error occured", err)
		}

		content, totalCount, _ := underTest.repository.FindAllWithPageable(searchModel, pageableRequest)
		want := common.CreatePageableResult(content, totalCount)
		if !reflect.DeepEqual(want, got) {
			t.Errorf("want %+v but got %+v", want, got)
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
			useCase := NewProductService(mockRepo, domain.ProductFactory{})

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

func (m *mockProductRepo) FindAll(searchModel domain.ProductSearchModel) ([]*domain.Product, error) {
	resp := newValidProductListWithName("T-shirt")
	return resp, nil
}

func (m *mockProductRepo) FindAllWithPageable(searchModel domain.ProductSearchModel, pageable common.PageableRequest) ([]*domain.Product, int, error) {
	result := newValidProductListWithName("T-shirt")
	return result, 1, nil
}

func newValidProduct() *domain.Product {
	return &domain.Product{
		Id:     1,
		Name:   "T-shirt",
		Price:  100,
		ShopId: 1,
	}
}

func newValidProductListWithName(name string) []*domain.Product {
	list := []*domain.Product{{
		Id:     1,
		Name:   name,
		Price:  100,
		ShopId: 1,
	}}
	return list
}
