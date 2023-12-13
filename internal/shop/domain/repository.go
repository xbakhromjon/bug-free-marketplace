package domain

import (
	"errors"
	"golang-project-template/internal/common"
)

var saveShopError = errors.New("save shop error")

type ShopRepository interface {
	Save(shop NewShop) (int, error)
	CheckShopNameExists(string) (bool, error)
}

type UserRepo interface {
	UserExists(id int) (bool, error)
}

type ProductRepository interface {
	Save(product *Product) (int, error)
	FindById(id int) (*Product, error)
	FindAllByShopId(shopId int) ([]*Product, error)
	FindAll(model ProductSearchModel) ([]*Product, error)
	FindAllWithPageable(searchModel ProductSearchModel, pageable common.PageableRequest) (*common.PageableResult[Product], error)
}

const (
	ErrProductNotFound = Err("product not found")
	ErrEmptyShopName   = Err("shop name can not be empty")
	ErrInvalidShopName = Err("shop name max length must be 128 characters")
	ErrShopNameExists  = Err("This shop name already exists")
	ErrUserNotExists   = Err("No such user")
)

type Err string

func (e Err) Error() string {
	return string(e)
}
