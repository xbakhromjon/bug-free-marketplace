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

func (b *basketRepo) CreateBasket(userId int) (id int, err error) {
	row := b.db.QueryRow("INSERT INTO basket(user_id,purchased) VALUES ($1,false) RETURNING id;", userId)
	err = row.Scan(row)
	if err != nil {
		return 0, err
	}
	return id, nil
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
	row, _ := b.db.Query("SELECT * from basket_items WHERE basket_id = $1", basketId)
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

func (b *basketRepo) UpdateBasketItem(basketId, quantity int) error {
	_, err := b.db.Exec("UPDATE basket_items SET quantity = quantity + $1 WHERE basket_id = $2", quantity, basketId)
	if err != nil {
		return domain.ErrBasketUpdateFailed
	}
	return nil
}

func (b *basketRepo) UpdateBasketStatus(basketId int) error {
	_, err := b.db.Exec("UPDATE basket SET purchased = true WHERE id = $2", basketId)
	if err != nil {
		return err
	}
	return nil
}

func (b *basketRepo) DeleteProduct(basketId, productId int) (id int, err error) {
	row := b.db.QueryRow("delete from basket_items where basket_id = $1 AND product_id = $2 RETURNING id", basketId, productId)
	if err != nil {
		return 0, domain.ErrDeleteItemFailed
	}
	err = row.Scan(&id)
	if err != nil {
		return 0, domain.ErrIDScanFailed
	}
	return id, nil
}
