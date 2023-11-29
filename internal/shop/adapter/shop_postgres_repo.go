package adapters

import (
	"database/sql"
	"fmt"
	"golang-project-template/internal/shop/domain"
)

var (
	shopTableName = "shop"
)

type shopPostgresRepo struct {
	db *sql.DB
}

func NewShopRepository(db *sql.DB) domain.ShopRepository {
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
