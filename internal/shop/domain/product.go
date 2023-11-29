package domain

type Product struct {
	Id     int
	Name   string
	Price  int
	ShopId int
}

type NewProduct struct {
	Name   string
	Price  int
	ShopId int
}

type ProductRepository interface {
	Save(product *Product) (int, error)
	FindById(id int) (*Product, error)
	FindAllByShopId(shopId int) ([]*Product, error)
}
