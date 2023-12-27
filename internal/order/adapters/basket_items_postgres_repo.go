package adapters

import (
	"github.com/jackc/pgx"
	basketItem "golang-project-template/internal/order/domain"
)

type basketItemRepo struct {
	db *pgx.Conn
}

func NewBasketItemRepository(db *pgx.Conn) basketItem.BasketItemRepository {
	return &basketItemRepo{db: db}
}

func (b *basketItemRepo) AddItem(items *basketItem.BasketItems) (id int, err error) {
	row := b.db.QueryRow("INSERT INTO basket_items(basket_id,product_id,quantity) VALUES ($1,$2,$3) RETURNING id",
		items.BasketId, items.ProductId, items.Quantity)
	err = row.Scan(&id)
	if err != nil {
		return 0, basketItem.ErrIDScanFailed
	}
	return id, nil
}

func (b *basketItemRepo) GetAll(basketId int) ([]basketItem.BasketItems, error) {
	row, err := b.db.Query("SELECT b.id, b.basket_id, b.product_id, b.quantity from basket_items b WHERE b.basket_id = $1", basketId)
	if err != nil {
		return nil, err
	}
	var Items []basketItem.BasketItems
	for row.Next() {
		var bItems basketItem.BasketItems
		err := row.Scan(&bItems.Id, &bItems.BasketId, &bItems.ProductId, &bItems.Quantity)
		if err != nil {
			return nil, basketItem.ErrIDScanFailed
		}
		Items = append(Items, bItems)
	}
	return Items, nil
}

func (b *basketItemRepo) UpdateBasketItem(bItemId, quantity int) error {
	_, err := b.db.Exec("UPDATE basket_items SET quantity = quantity + $1 WHERE id = $2", quantity, bItemId)
	if err != nil {
		return basketItem.ErrBasketUpdateFailed
	}
	return nil
}

func (b *basketItemRepo) DeleteProduct(bItemId int) (id int, err error) {
	row := b.db.QueryRow("delete from basket_items where id = $1 AND product_id = $2 RETURNING id", bItemId)
	if err != nil {
		return 0, basketItem.ErrDeleteItemFailed
	}
	err = row.Scan(&id)
	if err != nil {
		return 0, basketItem.ErrIDScanFailed
	}
	return id, nil
}
