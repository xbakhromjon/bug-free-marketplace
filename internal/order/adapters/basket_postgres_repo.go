package adapters

import (
	"github.com/jackc/pgx"
	basketdomain "golang-project-template/internal/order/domain"
)

type basketRepo struct {
	db *pgx.Conn
}

func NewBasketRepository(db *pgx.Conn) basketdomain.BasketRepository {
	return &basketRepo{db: db}
}

func (b *basketRepo) CreateBasket(userId int) (id int, err error) {
	row := b.db.QueryRow("INSERT INTO basket(user_id,purchased) VALUES ($1,false) RETURNING id;", userId)
	err = row.Scan(&id)
	if err != nil {
		return 0, basketdomain.ErrIDScanFailed
	}
	return id, nil
}

func (b *basketRepo) GetBasketWithItems(basketId int) (*basketdomain.BasketWithItems, error) {
	query := `
		SELECT b.*, bi.id as ItemId, bi.product_id, bi.quantity
		FROM basket b
		INNER JOIN basket_items bi ON b.id = bi.basket_id
		WHERE b.id = $1`
	rows, err := b.db.Query(query, basketId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var basketWithItems basketdomain.BasketWithItems
	for rows.Next() {
		var basket basketdomain.Basket
		var items basketdomain.BasketItems
		err := rows.Scan(
			&basket.Id, &basket.UserId, &basket.Purchased,
			&items.Id, &items.BasketId, &items.ProductId, &items.Quantity,
		)
		if err != nil {
			return nil, err
		}
		basketWithItems.Items = append(basketWithItems.Items, items)
		basketWithItems.Basket = basket
	}
	return &basketWithItems, nil
}

func (b *basketRepo) GetActiveBasket(userID int) (*basketdomain.Basket, error) {
	row := b.db.QueryRow("SELECT b.id,b.user_id,b.purchased FROM basket as b WHERE b.user_id = $1 AND b.purchased = false", userID)
	var basket basketdomain.Basket
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
