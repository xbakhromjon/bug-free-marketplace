package domain

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
	ErrInvalidShopName = Err("invalid shop name")
)

type Err string

func (e Err) Error() string {
	return string(e)
}
