package adapters

import (
	"fmt"
	"golang-project-template/internal/shop/domain"
	"time"

	"github.com/jackc/pgx"
)

var (
	shopTableName = "shop"
	CreatedAt     time.Time
	UpdatedAt     time.Time
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

	err := s.db.QueryRow(
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

	return shops, nil
}
