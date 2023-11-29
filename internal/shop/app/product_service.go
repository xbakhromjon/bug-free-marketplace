package app

import (
	"golang-project-template/internal/shop/domain"
)

type ProductService interface {
	Add(req domain.NewProduct) (int, error)
	GetOne(id int) (*domain.Product, error)
	GetAllByShopId(shopId int) ([]*domain.Product, error)
}

func NewProductService(repository domain.ProductRepository) ProductService {

	return &productService{repository: repository}
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
