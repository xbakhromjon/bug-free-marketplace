package adapter

import (
	"database/sql"
	"errors"
	"fmt"
	"golang-project-template/internal/shop/domain"
)

type productAdapter struct {
	db *sql.DB
}

func NewProductAdapter(db *sql.DB) domain.ProductRepository {

	return &productAdapter{db: db}
}

func (p *productAdapter) Save(product *domain.Product) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (p *productAdapter) FindById(id int) (*domain.Product, error) {
	row, err := p.db.Query(`select p.id, p.name, p.price, p.shop_id from products p where p.id = $1`, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &domain.ErrProductNotFound{fmt.Sprintf("Product not found with %d id", id)}
		}
	}
	var product domain.Product
	err = row.Scan(&product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *productAdapter) FindAllByShopId(shopId int) ([]*domain.Product, error) {
	//TODO implement me
	panic("implement me")
}
