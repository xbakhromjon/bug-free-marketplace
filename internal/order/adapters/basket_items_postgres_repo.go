package adapters

import (
	"github.com/jackc/pgx"
	"golang-project-template/internal/order/domain"
)

type basketItemRepo struct {
	db *pgx.Conn
}

func NewBasketItemRepository(db *pgx.Conn) domain.BasketItemRepository {
	return &basketItemRepo{db: db}
}

func (b basketItemRepo) AddItem(items *domain.BasketItems) (id int, err error) {
	row := b.db.QueryRow("INSERT INTO basket_items(basket_id,product_id,quantity) VALUES ($1,$2,$3) RETURNING id",
		items.BasketId, items.ProductId, items.Quantity)
	err = row.Scan(&id)
	if err != nil {
		return 0, domain.ErrIDScanFailed
	}
	return id, nil
}

func (b basketItemRepo) GetAll(basketId int) ([]domain.BasketItems, error) {
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

func (b basketItemRepo) UpdateBasketItem(bItemId, quantity int) error {
	_, err := b.db.Exec("UPDATE basket_items SET quantity = quantity + $1 WHERE id = $2", quantity, bItemId)
	if err != nil {
		return domain.ErrBasketUpdateFailed
	}
	return nil
}

func (b basketItemRepo) DeleteProduct(bItemId int) error {
	err := b.db.QueryRow("delete from basket_items where id = $1 AND product_id = $2 RETURNING id", bItemId)
	if err != nil {
		return domain.ErrDeleteItemFailed
	}
	return nil
}
