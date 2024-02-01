package app

import (
	"golang-project-template/internal/common"
	"golang-project-template/internal/shop/domain"
)

type ProductService interface {
	Add(req domain.NewProduct) (int, error)
	GetOne(id int) (*domain.Product, error)
	GetAllByShopId(shopId int) ([]*domain.Product, error)
	Filter(searchModel domain.ProductSearchModel) ([]*domain.Product, error)
	FilterByPageable(searchModel domain.ProductSearchModel, pageable common.PageableRequest) (*common.PageableResult[*domain.Product], error)
	UpdateProduct(productID int, product *domain.Product) error
}

func NewProductService(repository domain.ProductRepository, factory domain.ProductFactory) ProductService {

	return &productService{repository: repository, factory: factory}
}

type productService struct {
	repository domain.ProductRepository
	factory    domain.ProductFactory
}

func (p *productService) UpdateProduct(productID int, product *domain.Product) error {
	_, err := p.repository.UpdateProduct(productID, product)
	return err
}

func (p *productService) Add(req domain.NewProduct) (int, error) {
	product := p.factory.NewProduct(req.Name, req.Price, req.ShopId)
	id, err := p.repository.Save(product)
	if err != nil {
		return 0, err
	}
	return id, nil
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

func (p *productService) Filter(searchModel domain.ProductSearchModel) ([]*domain.Product, error) {
	products, err := p.repository.FindAll(searchModel)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (p *productService) FilterByPageable(searchModel domain.ProductSearchModel, pageable common.PageableRequest) (*common.PageableResult[*domain.Product], error) {
	content, totalCount, err := p.repository.FindAllWithPageable(searchModel, pageable)
	if err != nil {
		return nil, err
	}
	pageableResult := common.CreatePageableResult(content, totalCount)
	return pageableResult, nil
}
