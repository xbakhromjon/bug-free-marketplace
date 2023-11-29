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
