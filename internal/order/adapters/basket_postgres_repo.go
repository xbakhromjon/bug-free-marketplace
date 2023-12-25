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
		SELECT b.*, bi.id as ItemId, bi.product_id, bi.quantity
		FROM basket b
		INNER JOIN basket_items bi ON b.id = bi.basket_id
		WHERE b.id = $1`
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

func (b *basketRepo) GetActiveBasket(userID int) (*domain.Basket, error) {
	row := b.db.QueryRow("SELECT b.id FROM basket b WHERE b.user_id = $1 AND b.purchased = false", userID)
	var basket domain.Basket
	if err := row.Scan(&basket.Id, &basket.UserId, &basket.Purchased); err != nil {
		return nil, err
	}
	return &basket, nil
}

func (b *basketRepo) MarkBasketAsPurchased(userId, basketId int) error {
	_, err := b.db.Exec("UPDATE basket SET purchased = true WHERE user_id = $1 AND basketId = $2", userId, basketId)
	if err != nil {
		return err
	}
	return nil
}
