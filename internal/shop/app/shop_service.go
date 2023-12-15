package app

import (
	"golang-project-template/internal/shop/domain"
)

type ShopService interface {
	Create(domain.NewShop) (int, error)
	GetShopById(int) (*domain.Shop, error)
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

func (s *shopService) Create(req domain.NewShop) (int, error) {

	err := s.shopFactory.Validate(req)
	if err != nil {
		return 0, err
	}

	ok, err := s.userRepository.UserExists(req.GetOwnerId())
	if err != nil {
		return 0, err
	}
	if !ok {
		return 0, domain.ErrUserNotExists
	}

	shopNameExists, err := s.repository.CheckShopNameExists(req.GetName())
	if err != nil {
		return 0, err
	}

	if shopNameExists {
		return 0, domain.ErrShopNameExists
	}

	return s.repository.Save(req)
}

func (s *shopService) GetShopById(id int) (*domain.Shop, error) {
	return s.repository.FindShopById(id)
}

func (s *shopService) GetAllShops(limit, offset int, search string) ([]domain.Shop, error) {
	err := s.shopFactory.GetAllShopsInputValidate(limit, offset, search)
	if err != nil {
		return []domain.Shop{}, err
	}
	return s.repository.FindAllShops(limit, offset, search)
}
