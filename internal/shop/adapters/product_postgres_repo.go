package adapters

import (
	"database/sql"
	"errors"
	"golang-project-template/internal/shop/domain"

	"github.com/jackc/pgx"
)

type productPostgresRepo struct {
	db *pgx.Conn
}

func NewProductRepository(db *pgx.Conn) domain.ProductRepository {

	return &productPostgresRepo{db: db}
}

func (p *productPostgresRepo) Save(product *domain.Product) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (p *productPostgresRepo) FindById(id int) (*domain.Product, error) {
	row, err := p.db.Query(`select p.id, p.name, p.price, p.shop_id from products p where p.id = $1`, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrProductNotFound
		}
	}

	var product domain.Product
	err = row.Scan(&product)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *productPostgresRepo) FindAllByShopId(shopId int) ([]*domain.Product, error) {
	//TODO implement me
	panic("implement me")
}
