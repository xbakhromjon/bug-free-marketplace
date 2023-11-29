package adapter

import (
	"golang-project-template/internal/shop/domain"

	"github.com/jmoiron/sqlx"
)

type shopPostgres struct {
	db *sqlx.DB
}

func NewShopPostgres(db *sqlx.DB) domain.ShopRepository {
	return &shopPostgres{db: db}
}

func (s *shopPostgres) Save(shop domain.NewShop) (int, error) {
	var id int

	createShopQuery := `
		INSERT INTO shop(
			name,
			owner_id
		)
		VALUES
			($1, $2)
		RETURNING
			id
	`

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
