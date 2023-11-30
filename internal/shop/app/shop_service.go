package app

import (
	"golang-project-template/internal/shop/domain"
)

type ShopService interface {
	Create(domain.NewShop) (int, error)
}

type shopService struct {
	repository  domain.ShopRepository
	shopFactory domain.ShopFactory
}

func NewShopService(
	repository domain.ShopRepository,
	shopFactory domain.ShopFactory,
) ShopService {

	return &shopService{repository: repository, shopFactory: shopFactory}
}

func (s *shopService) Create(req domain.NewShop) (int, error) {

	err := s.shopFactory.Validate(req)

	if err != nil {
		return 0, err
	}

	shopNameExists, err := s.repository.CheckShopNameExists(req.Name)
	if err != nil {
		return 0, err
	}

	if shopNameExists {
		return 0, domain.ErrShopNameExists
	}

	return s.repository.Save(req)
}
