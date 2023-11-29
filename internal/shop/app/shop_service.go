package app

import "golang-project-template/internal/shop/domain"

type ShopService interface {
	Create(req domain.NewShop) (int, error)
}

type shopService struct {
	repository domain.ShopRepository
}

func NewShopService(repository domain.ShopRepository) ShopService {
	return &shopService{repository: repository}
}

func (s *shopService) Create(req domain.NewShop) (int, error) {
	return s.repository.Save(req)
}
