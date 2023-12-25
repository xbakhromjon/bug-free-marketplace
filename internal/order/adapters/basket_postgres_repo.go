package adapters

import (
	"github.com/jackc/pgx"
	"golang-project-template/internal/order/domain"
)

type basketRepo struct {
	db *pgx.Conn
}

func NewCartRepository(db *pgx.Conn) domain.BasketRepository {
	return &basketRepo{db: db}
}

type BasketWithItems struct {
	domain.Basket
	items []domain.BasketItems
}

func (b *basketRepo) CreateBasket(userId int) (id int, err error) {
	row := b.db.QueryRow("INSERT INTO basket(user_id,purchased) VALUES ($1,$2) RETURNING id;", userId)
	err = row.Scan(row)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (b *basketRepo) GetBasket(basketId int) (*BasketWithItems, error) {
	query := `
		SELECT b.*, bi.Id as ItemId, bi.ProductId, bi.Quantity
		FROM Basket b
		LEFT JOIN BasketItems bi ON b.Id = bi.BasketId
		WHERE b.Id = $1`
	rows, err := b.db.Query(query, basketId)
	if err != nil {
		return nil, err
	}
	var basketWithItems BasketWithItems
	for rows.Next() {
		var basket domain.Basket
		var items domain.BasketItems
		err := rows.Scan(
			&basket.Id, &basket.UserId, &basket.Purchased,
			&items.Id, &items.BasketId, &items.ProductId, &items.Quantity,
		)
		if err != nil {
			return nil, err
		}
		basketWithItems.items = append(basketWithItems.items, items)
		basketWithItems.Basket = basket
	}
	return &basketWithItems, nil
}
