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

func (s *shopPostgresRepo) CheckUserExists(ownerId int) (bool, error) {
	var exists bool
	queryCheckUserExists := fmt.Sprintf(`
		SELECT 
		EXISTS(
			SELECT
				phone_number
			FROM 
				%s
			WHERE 
				id = $1
			AND
				deleted_at 
			IS NULL
		);
			
		`, "users")

	err := s.db.QueryRow(
		queryCheckUserExists,
		ownerId,
	).Scan(
		&exists,
	)

	if err != nil {
		return false, err
	}

	return exists, nil
}
