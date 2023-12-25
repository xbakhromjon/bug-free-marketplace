package app

//
//import (
//	"golang-project-template/internal/order/domain"
//)
//
//type BasketService interface {
//	CreateBasket(userID int) (int, error)
//	AddItem(userId int, basket *domain.BasketItems) (int, error)
//	GetAll(cartId int) ([]domain.BasketItems, error)
//	UpdateBasketQuantity(bItemId, quantity int) error
//	DeleteProduct(bItemId, productId int) (int, error)
//}
//
//func NewCartService(repo domain.BasketRepository) BasketService {
//	return &basketService{basketRepo: repo}
//}
//
//type basketService struct {
//	basketRepo domain.BasketRepository
//}
//
//func (b *basketService) CreateBasket(userID int) (int, error) {
//	basketId, err := b.CreateBasket(userID)
//	if err != nil {
//		return 0, domain.ErrIDScanFailed
//	}
//	return basketId, nil
//}
//
//func (b *basketService) AddItem(userId int, bItem *domain.BasketItems) (int, error) {
//	activeBasket, _ := b.basketRepo.GetActiveBasket(userId)
//	if !activeBasket.Purchased {
//		id, err := b.basketRepo.AddItem(bItem)
//		if err != nil {
//			return 0, domain.ErrAddItemFailed
//		}
//		return id, err
//	}
//	basketId, _ := b.basketRepo.CreateBasket(activeBasket.UserId)
//	return basketId, nil
//}
//
//func (b *basketService) GetAll(basketId int) ([]domain.BasketItems, error) {
//	cItems, err := b.GetAll(basketId)
//	if err != nil {
//		return nil, domain.ErrBasketItemsNotFound
//	}
//	return cItems, nil
//}
//
//func (b *basketService) UpdateBasketQuantity(bItemId, quantity int) error {
//	err := b.basketRepo.UpdateBasketItem(bItemId, quantity)
//	if err != nil {
//		return domain.ErrBasketUpdateFailed
//	}
//	return nil
//}
//
//func (b *basketService) DeleteProduct(bItemId, productId int) (int, error) {
//	return b.basketRepo.DeleteProduct(bItemId, productId)
//}
