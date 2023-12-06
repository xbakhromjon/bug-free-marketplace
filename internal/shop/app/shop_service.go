package app

import (
	"golang-project-template/internal/shop/domain"
)

type ShopService interface {
	Create(domain.NewShop) (int, error)
}

type shopService struct {
	repository     domain.ShopRepository
	shopFactory    domain.ShopFactory
	userRepository domain.UserRepo
}

func NewShopService(
	repository domain.ShopRepository,
	shopFactory domain.ShopFactory,
	userRepository domain.UserRepo,

) ShopService {

	return &shopService{
		repository:     repository,
		shopFactory:    shopFactory,
		userRepository: userRepository,
	}
}

func (s *shopService) Create(req domain.NewShop) (int, error) {

	err := s.shopFactory.Validate(req)
	if err != nil {
		return 0, err
	}

	ok, err := s.userRepository.UserExists(req.OwnerId)
	if err != nil {
		return 0, err
	}
	if !ok {
		return 0, domain.ErrUserNotExists
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
