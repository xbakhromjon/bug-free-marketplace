package adapters

import (
	"context"
	"database/sql"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"golang-project-template/internal/common"
	"golang-project-template/internal/shop/domain"
)

type productPostgresRepo struct {
	db      *pgx.Conn
	factory *domain.ProductFactory
}

func NewProductRepository(db *pgx.Conn, factory *domain.ProductFactory) domain.ProductRepository {
	return &productPostgresRepo{db: db, factory: factory}
}

func (p *productPostgresRepo) UpdateProduct(productID int, product *domain.Product) (*domain.Product, error) {
	_, err := p.db.Exec(context.Background(), "UPDATE product SET price=? WHERE id = ?", productID, product.GetPrice())
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p productPostgresRepo) Save(product *domain.Product) (int, error) {
	row := p.db.QueryRow(context.Background(), "INSERT INTO product (name, shop_id, price, count) VALUES ($1, $2, $3, $4) RETURNING id", product.GetName(), product.GetShopId(), product.GetPrice(), product.GetCount())
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
	products := sq.Select("p.id, p.name, p.price, p.shop_id").From("product p")

	if val, ok := searchModel.GetName(); ok {
		products.Where(sq.Eq{"name": val})
	}

	if val, ok := searchModel.GetPriceFrom(); ok {
		products.Where(sq.Gt{"price": val})
	}

	if val, ok := searchModel.GetPriceTo(); ok {
		products.Where(sq.Lt{"price": val})
	}

	// execute query
	rows, err := products.Query()
	if err != nil {
		return nil, err
	}
	var result []*domain.Product

	// scan rows and add to result
	for rows.Next() {
		var id int
		var name string
		var price int
		var shopId int
		err := rows.Scan(&id, &name, &price, &shopId)
		if err != nil {
			return nil, err
		}
		result = append(result, p.factory.CreateExistProduct(id, name, price, shopId))
	}
	return result, nil
}

func (p *productPostgresRepo) FindAllWithPageable(searchModel domain.ProductSearchModel, pageable common.PageableRequest) ([]*domain.Product, int, error) {
	products := sq.Select("p.id, p.name, p.price, p.shop_id").From("product p")
	productsCount := sq.Select("count(p.id) totalCount").From("product p")

	BuildProductFilterQuery(products, searchModel)
	BuildPageableQuery(products, pageable)
	BuildProductFilterQuery(productsCount, searchModel)

	// execute count query
	countRow := products.QueryRow()
	var totalCount int
	err := countRow.Scan(&totalCount)
	if err != nil {
		return nil, 0, err
	}

	// execute filter query
	rows, err := products.Query()
	if err != nil {
		return nil, 0, err
	}
	var result []*domain.Product

	// scan rows and add to result
	for rows.Next() {
		var id int
		var name string
		var price int
		var shopId int
		err := rows.Scan(&id, &name, &price, &shopId)
		if err != nil {
			return nil, 0, err
		}
		result = append(result, p.factory.CreateExistProduct(id, name, price, shopId))
	}

	return result, totalCount, nil
}

func BuildProductFilterQuery(base sq.SelectBuilder, searchModel domain.ProductSearchModel) {
	if val, ok := searchModel.GetName(); ok {
		base.Where(sq.Eq{"name": val})
	}

	if val, ok := searchModel.GetPriceFrom(); ok {
		base.Where(sq.Gt{"price": val})
	}

	if val, ok := searchModel.GetPriceTo(); ok {
		base.Where(sq.Lt{"price": val})
	}
}

func BuildPageableQuery(base sq.SelectBuilder, source common.PageableRequest) {
	page, ok := source.GetPage()
	if !ok {
		page = 1
	}

	size, ok := source.GetSize()
	if !ok {
		size = 10
	}
	base.Limit(size)
	base.Offset((page - 1) * size)

	if val, ok := source.GetSort(); ok {
		BuildSortQuery(base, val)
	}
}

func BuildSortQuery(base sq.SelectBuilder, source *common.SortRequest) {
	if val, ok := source.GetSort(); ok {
		base.OrderBy(val)
	}

	if val, ok := source.GetDirection(); ok {
		base.OrderBy(val)
	}
	// TODO: add direction

}
