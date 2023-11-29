package app

import "golang-project-template/internal/shop/domain"

type ShopUseCase interface {
	Create(req domain.NewShop) (int, error)
}
