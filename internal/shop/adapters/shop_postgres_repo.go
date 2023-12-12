package adapters

import (
	"context"
	"fmt"
	"golang-project-template/internal/common/postgres"
	"golang-project-template/internal/shop/domain"
	"time"
)

var (
	shopTableName = "shop"
	CreatedAt     time.Time
	UpdatedAt     time.Time
)

type shopPostgresRepo struct {
	db *postgres.PostgresDB
}

func NewShopRepository(db *postgres.PostgresDB) domain.ShopRepository {
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
		context.Background(),
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
		context.Background(),
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

	queryGetShopById := fmt.Sprintf(`
	SELECT 
		id,
		name,
		owner_id,
		created_at,
		updated_at
	FROM
		%s
	WHERE
		deleted_at IS NULL
	AND 
		id=$1
	
`, shopTableName)

	err := s.db.QueryRow(context.Background(),
		queryGetShopById,
		shopId,
	).Scan(
		&shop.Id,
		&shop.Name,
		&shop.OwnerId,
		&CreatedAt,
		&UpdatedAt,
	)

	if err != nil {
		return domain.Shop{}, err
	}

	shop.CreatedAt = CreatedAt.Format(time.RFC3339)
	shop.UpdatedAt = UpdatedAt.Format(time.RFC3339)

	return shop, nil
}

func (s *shopPostgresRepo) FindAllShops(limit, offset int, search string) ([]domain.Shop, error) {
	shops := []domain.Shop{}

	queryGetAllShops := fmt.Sprint(`
		SELECT 
			id,
			name,
			owner_id,
			created_at,
			updated_at
		FROM
			shop
		WHERE
			name ILIKE '%' || $1 || '%'
		AND
			deleted_at IS NULL
		LIMIT
			$2
		OFFSET
			$3
	`)

	row, err := s.db.Query(context.Background(), queryGetAllShops, search, limit, offset)

	if err != nil {
		return []domain.Shop{}, err
	}
	defer row.Close()

	for row.Next() {
		shop := domain.Shop{}
		err := row.Scan(
			&shop.Id,
			&shop.Name,
			&shop.OwnerId,
			&CreatedAt,
			&UpdatedAt,
		)
		if err != nil {
			return []domain.Shop{}, err
		}
		shop.CreatedAt = CreatedAt.Format(time.RFC1123)
		shop.UpdatedAt = UpdatedAt.Format(time.RFC1123)

		shops = append(shops, shop)
	}

	return shops, nil
}
