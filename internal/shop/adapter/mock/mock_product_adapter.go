package mock

import (
	"database/sql"
	"fmt"
	"golang-project-template/internal/shop/domain"
)

type mockProductAdapter struct {
}

func NewMockProductAdapter(db *sql.DB) domain.ProductRepository {
	return &mockProductAdapter{}
}

func (m *mockProductAdapter) Save(product *domain.Product) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (m *mockProductAdapter) FindById(id int) (*domain.Product, error) {
	if id == 1 {
		return &domain.Product{
			Id:     1,
			Name:   "T-shirt",
			Price:  100,
			ShopId: 1,
		}, nil
	}
	return nil, &domain.ErrProductNotFound{Err: fmt.Sprintf("Product not found with %d", id)}
}

func (m *mockProductAdapter) FindAllByShopId(shopId int) ([]*domain.Product, error) {
	//TODO implement me
	panic("implement me")
}
