package app

import (
	"golang-project-template/internal/shop/domain"
)

type ShopService interface {
	Create(domain.NewShop) (int, error)
}

type shopService struct {
	repository domain.ShopRepository
}

func NewShopService(repository domain.ShopRepository) ShopService {
	return &shopService{repository: repository}
}

func (s *shopService) Create(req domain.NewShop) (int, error) {
	err := ValidateShop(&req)

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

func ValidateShop(newShop *domain.NewShop) error {

	if newShop.Name == "" {
		return domain.ErrEmptyShopName
	}

	if len(newShop.Name) > 128 {
		return domain.ErrInvalidShopName
	}

	return nil
}
