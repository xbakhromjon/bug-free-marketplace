package app

import (
	"golang-project-template/internal/order/domain"
)

type BasketService interface {
	CreateBasket(userID int) (int, error)
	AddItem(userId int, basket *domain.BasketItems) (int, error)
	GetBasketWithItemsById(basketId int) (*domain.BasketWithItems, error)
	GetAll(basketId int) ([]domain.BasketItems, error)
	UpdateBasketQuantity(bItemId, quantity int) error
	MarkBasketAsPurchased(userId, basketId int) error
	DeleteProduct(bItemId int) (int, error)
}

func NewBasketService(basketRepo domain.BasketRepository, repoItem domain.BasketItemRepository) BasketService {
	return &basketService{
		basketRepo:     basketRepo,
		basketItemRepo: repoItem}
}

type basketService struct {
	basketRepo     domain.BasketRepository
	basketItemRepo domain.BasketItemRepository
}

func (b *basketService) CreateBasket(userID int) (int, error) {
	basketId, err := b.basketRepo.CreateBasket(userID)
	if err != nil {
		return 0, domain.ErrIDScanFailed
	}
	return basketId, nil
}

func (b *basketService) AddItem(userId int, bItems *domain.BasketItems) (int, error) {
	activeBasket, _ := b.basketRepo.GetActiveBasket(userId)
	if !activeBasket.Purchased {
		id, err := b.basketItemRepo.AddItem(bItems)
		if err != nil {
			return 0, domain.ErrAddItemFailed
		}
		return id, err
	}
	newBasketId, err := b.basketRepo.CreateBasket(userId)
	if err != nil {
		return 0, domain.ErrBasketCreationFailed
	}
	id, err := b.basketItemRepo.AddItem(&domain.BasketItems{
		BasketId:  newBasketId,
		ProductId: bItems.ProductId,
		Quantity:  bItems.Quantity,
	})
	if err != nil {
		return 0, domain.ErrAddItemFailed
	}
	return id, nil
}

func (b *basketService) GetBasketWithItemsById(basketId int) (*domain.BasketWithItems, error) {
	bWithItems, err := b.basketRepo.GetBasket(basketId)
	if err != nil {
		return nil, err
	}
	return bWithItems, nil
}

func (b *basketService) GetAll(basketId int) ([]domain.BasketItems, error) {
	cItems, err := b.basketItemRepo.GetAll(basketId)
	if err != nil {
		return nil, domain.ErrBasketItemsNotFound
	}
	return cItems, nil
}

func (b *basketService) UpdateBasketQuantity(bItemId, quantity int) error {
	err := b.basketItemRepo.UpdateBasketItem(bItemId, quantity)
	if err != nil {
		return domain.ErrBasketUpdateFailed
	}
	return nil
}

func (b *basketService) MarkBasketAsPurchased(userId, basketId int) error {
	err := b.basketRepo.MarkBasketAsPurchased(userId, basketId)
	if err != nil {
		return err
	}
	return nil
}

func (b *basketService) DeleteProduct(bItemId int) (int, error) {
	return b.basketItemRepo.DeleteProduct(bItemId)
}
