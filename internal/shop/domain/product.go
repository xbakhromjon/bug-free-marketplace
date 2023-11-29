package domain

type Product struct {
	Id    int
	Name  string
	Price int
}

type NewProduct struct {
	Id     int
	Name   string
	Price  int
	ShopId int
}

type ProductRepository interface {
	Save(product *Product) (int, error)
	GetById(id int) (*Product, error)
	GetAllByShopId(shopId int) ([]*Product, error)
}
