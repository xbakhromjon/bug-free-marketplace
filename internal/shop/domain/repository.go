package domain

import "errors"

var saveShopError = errors.New("save shop error")

type ShopRepository interface {
	Save(shop NewShop) (int, error)
}

type ProductRepository interface {
	Save(product *Product) (int, error)
	FindById(id int) (*Product, error)
	FindAllByShopId(shopId int) ([]*Product, error)
}
