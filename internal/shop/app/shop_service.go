package app

import "golang-project-template/internal/shop/domain"

type ShopService interface {
	Create(req domain.NewShop) (int, error)
}
