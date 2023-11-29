package app

import "golang-project-template/internal/shop/domain"

type ShopService interface {
	Create(req domain.NewShop) (int, error)
}

type shopService struct {
	shopRepo domain.ShopRepository
}

func NewShopService(shopRepo domain.ShopRepository) ShopService {
	return &shopService{shopRepo: shopRepo}
}

func (s *shopService) Create(req domain.NewShop) (int, error) {
	return s.shopRepo.Save(req)
}
