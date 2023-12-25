package adapters

import (
	"github.com/jackc/pgx"
	"golang-project-template/internal/order/domain"
)

type basketItemRepo struct {
	db *pgx.Conn
}

func NewBasketItemRepository(db *pgx.Conn) domain.BasketRepository {
	return &basketRepo{db: db}
}

func (b *basketRepo) AddItem(items *domain.BasketItems) (id int, err error) {
	row := b.db.QueryRow("INSERT INTO basket_items(basket_id,product_id,quantity) VALUES ($1,$2,$3) RETURNING id",
		items.BasketId, items.ProductId, items.Quantity)
	err = row.Scan(&id)
	if err != nil {
		return 0, domain.ErrIDScanFailed
	}
	return id, nil
}

func (b *basketRepo) GetAll(basketId int) ([]domain.BasketItems, error) {
	row, _ := b.db.Query("SELECT b.id, b.basket_id, b.product_id, b.quantity from basket_items b WHERE b.basket_id = $1", basketId)
	var Items []domain.BasketItems
	for row.Next() {
		var bItems domain.BasketItems
		err := row.Scan(&bItems.Id, &bItems.BasketId, &bItems.ProductId, &bItems.Quantity)
		if err != nil {
			return nil, domain.ErrIDScanFailed
		}
		Items = append(Items, bItems)
	}
	return Items, nil
}

func (b *basketRepo) GetActiveBasket(userID int) (*domain.Basket, error) {
	row := b.db.QueryRow("SELECT b.basket_id FROM basket b WHERE b.user_id = $1 AND b.purchased = false", userID)
	var basket domain.Basket
	if err := row.Scan(&basket.Id, &basket.UserId, &basket.Purchased); err != nil {
		return nil, err
	}
	return &basket, nil
}

func (b *basketRepo) UpdateBasketItem(basketItemId, quantity int) error {
	_, err := b.db.Exec("UPDATE basket_items SET quantity = quantity + $1 WHERE id = $2", quantity, basketItemId)
	if err != nil {
		return domain.ErrBasketUpdateFailed
	}
	return nil
}

func (b *basketRepo) DeleteProduct(basketItemId int) (id int, err error) {
	row := b.db.QueryRow("delete from basket_items where id = $1 AND product_id = $2 RETURNING id", basketItemId)
	if err != nil {
		return 0, domain.ErrDeleteItemFailed
	}
	err = row.Scan(&id)
	if err != nil {
		return 0, domain.ErrIDScanFailed
	}
	return id, nil
}
