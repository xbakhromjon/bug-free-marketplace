package domain

type Product struct {
	Id    int
	Name  string
	Price int
}

type ProductRequest struct {
	Id    int
	Name  string
	Price int
}

type ProductResponse struct {
	Id    int
	Name  string
	Price int
}

type ProductUseCase interface {
	Add(req *ProductRequest) (int, error)
	GetOne(id int) (*ProductResponse, error)
	GetAllShopProducts(shopId int) ([]*ProductResponse, error)
}
