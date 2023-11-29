package app

import (
	"database/sql"
	"golang-project-template/internal/shop/adapter"
	"golang-project-template/internal/shop/domain"
)

type ProductService interface {
	Add(req domain.NewProduct) (int, error)
	GetOne(id int) (*domain.Product, error)
	GetAllByShopId(shopId int) ([]*domain.Product, error)
}

func NewProductService(db *sql.DB) ProductService {

	return &productService{repository: adapter.NewProductAdapter(db)}
}

type productService struct {
	repository domain.ProductRepository
}

func (p *productService) Add(req domain.NewProduct) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (p *productService) GetOne(id int) (*domain.Product, error) {
	product, err := p.repository.FindById(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *productService) GetAllByShopId(shopId int) ([]*domain.Product, error) {
	//TODO implement me
	panic("implement me")
}
