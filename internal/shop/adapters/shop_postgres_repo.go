package adapters

import (
	"fmt"
	"golang-project-template/internal/shop/domain"

	"github.com/jackc/pgx"
)

var (
	shopTableName = "shop"
)

type shopPostgresRepo struct {
	db *pgx.Conn
}

func NewShopRepository(db *pgx.Conn) domain.ShopRepository {
	return &shopPostgresRepo{db: db}
}

func (s *shopPostgresRepo) Save(shop domain.NewShop) (int, error) {
	var id int

	createShopQuery := fmt.Sprintf(`
		INSERT INTO %s(
			name,
			owner_id
		)
		VALUES
			($1, $2)
		RETURNING
			id
	`, shopTableName)

	err := s.db.QueryRow(
		createShopQuery,
		shop.Name,
		shop.OwnerId,
	).Scan(
		&id,
	)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *shopPostgresRepo) CheckShopNameExists(shopName string) (bool, error) {
	var exists bool
	queryCheckShopNameExists := fmt.Sprintf(`
		SELECT 
		EXISTS(
			SELECT
				name
			FROM 
				%s
			WHERE 
				name = $1
			AND 
				deleted_at
			IS NULL
		);
			
		`, shopTableName)

	err := s.db.QueryRow(
		queryCheckShopNameExists,
		shopName,
	).Scan(
		&exists,
	)

	if err != nil {
		return false, err
	}

	return exists, nil
}

func (s *shopPostgresRepo) FindShopById(shopId int) (domain.Shop, error) {
	shop := domain.Shop{}

	return shop, nil
}
