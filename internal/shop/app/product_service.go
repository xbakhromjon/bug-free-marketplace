package app

import "golang-project-template/internal/shop/domain"

type ProductUseCase interface {
	Add(req domain.NewProduct) (int, error)
	GetOne(id int) (*domain.Product, error)
	GetAllByShopId(shopId int) ([]*domain.Product, error)
}
