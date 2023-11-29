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

const (
	ErrProductNotFound = Err("product not found")
)

type Err string

func (e Err) Error() string {
	return string(e)
}
