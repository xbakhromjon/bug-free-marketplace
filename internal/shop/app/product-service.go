package app

import "golang-project-template/internal/shop/domain"

type ProductUseCase interface {
	Add(req domain.NewProduct) (int, error)
	GetOne(id int) (*domain.Product, error)
	GetAllByShopId(shopId int) ([]*domain.Product, error)
}

type ProductUseCaseImpl struct {
	ProductRepo domain.ProductRepository
}

func NewProductUseCase(productRepo domain.ProductRepository) ProductUseCase {
	return &ProductUseCaseImpl{
		ProductRepo: productRepo,
	}
}

func (uc ProductUseCaseImpl) Add(req domain.NewProduct) (int, error) {
	product := &domain.Product{
		Id:     req.Id,
		Name:   req.Name,
		Price:  req.Price,
		ShopId: req.ShopId,
	}
	id, err := uc.ProductRepo.Save(product)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (uc *ProductUseCaseImpl) GetOne(id int) (*domain.Product, error) {
	return uc.ProductRepo.GetById(id)
}

func (uc *ProductUseCaseImpl) GetAllByShopId(shopId int) ([]*domain.Product, error) {
	return uc.ProductRepo.GetAllByShopId(shopId)
}
