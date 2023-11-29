package domain

import "database/sql"

type Product struct {
	Id     int
	Name   string
	Price  int
	ShopId int
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
	FindById(id int) (*Product, error)
	FindAllByShopId(shopId int) ([]*Product, error)
}

type NewProductRepository interface {
	Add(product *NewProduct) error
}

type SqlProductRepo struct {
	db *sql.DB
}

func (repo *SqlProductRepo) Save(product *Product) (int, error) {
	_, err := repo.db.Exec("INSERT INTO products (name, price, shop_id) VALUES (?, ?, ?)", product.Name, product.Price, product.ShopId)
	if err != nil {
		return 0, err
	}
	return product.Id, nil
}
