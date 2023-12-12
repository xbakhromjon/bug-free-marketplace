package adapters

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jackc/pgx/v5"
	"golang-project-template/internal/common"
	"golang-project-template/internal/shop/domain"
	"strings"
)

var productPostgresRepoInstance *productPostgresRepo

type productPostgresRepo struct {
	db *pgx.Conn
}

func GetProductRepositoryInstance(db *pgx.Conn) domain.ProductRepository {
	if productPostgresRepoInstance == nil {
		productPostgresRepoInstance = &productPostgresRepo{db: db}
	}

	return productPostgresRepoInstance
}

func NewProductRepository(db *pgx.Conn) domain.ProductRepository {
	return &productPostgresRepo{db: db}
}

func (p productPostgresRepo) Save(product *domain.Product) (int, error) {
	row := p.db.QueryRow(context.Background(), "INSERT INTO product (name, shop_id, price, count) VALUES ($1, $2, $3, $4) RETURNING id", product.Name, product.ShopId, product.Price, product.Count)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (p *productPostgresRepo) FindById(id int) (*domain.Product, error) {
	row, err := p.db.Query(context.Background(), `select p.id, p.name, p.price, p.shop_id from product p where p.id = ?`, id)
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

func (p *productPostgresRepo) FindAll(searchModel domain.ProductSearchModel) ([]*domain.Product, error) {
	// define base query
	baseQuery := strings.Builder{}
	baseQuery.WriteString(`
		select p.id, p.name, p.price, p.shop_id from product p  
	`)

	// define arguments to pass query
	var args []any

	// add where clauses by searchModel
	if searchModel.Search != "" {
		baseQuery.WriteString(
			`
		where (p.Name ilike '%' | ? | '%')
		`)
		// add argument
		args = append(args, searchModel.Search)
	}
	// ...

	// execute query
	rows, err := p.db.Query(context.Background(), baseQuery.String(), args...)
	if err != nil {
		return nil, err
	}

	var result []*domain.Product

	// scan rows and add to result
	for rows.Next() {
		var product domain.Product
		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.ShopId)
		if err != nil {
			return nil, err
		}
		result = append(result, &product)
	}
	return result, nil
}

func (p *productPostgresRepo) FindAllWithPageable(searchModel domain.ProductSearchModel, pageable common.PageableRequest) (*common.PageableResult[domain.Product], error) {

	return nil, nil
}
