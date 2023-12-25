package app

import (
	"golang-project-template/internal/shop/domain"
)

type ShopService interface {
	Create(NewShop) (int, error)
	GetShopById(int) (domain.Shop, error)
	GetAllShops(int, int, string) ([]domain.Shop, error)
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

func (s *shopService) Create(req NewShop) (int, error) {

	shop, err := s.shopFactory.NewShop(req.Name, req.OwnerId)
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

	return s.repository.Save(shop)
}

func (s *shopService) GetShopById(id int) (domain.Shop, error) {
	return s.repository.FindShopById(id)
}

func (s *shopService) GetAllShops(limit, offset int, search string) ([]domain.Shop, error) {
	err := s.shopFactory.GetAllShopsInputValidate(limit, offset, search)
	if err != nil {
		return []domain.Shop{}, err
	}
	return s.repository.FindAllShops(limit, offset, search)
}
