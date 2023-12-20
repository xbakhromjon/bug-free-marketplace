package adapters

import (
	"github.com/jackc/pgx"
	basket "golang-project-template/internal/basket/domain"
)

type cartRepo struct {
	db *pgx.Conn
}

func (c cartRepo) CreateBasket(userId int) (id int, err error) {
	row, err := c.db.Query("INSERT INTO cart(user_id)VALUES ($1) RETURNING id", userId)
	if err != nil {
		return 0, basket.ErrCartCreationFailed
	}
	err = row.Scan(&id)
	if err != nil {
		return 0, basket.ErrIDScanFailed
	}
	return id, nil
}

func (c cartRepo) GetAll(cartId int) ([]*basket.CartItems, error) {
	row, _ := c.db.Query("SELECT * from cart_items WHERE cart_id = $1", cartId)

	var Items []*basket.CartItems
	for row.Next() {
		var cItem *basket.CartItems
		err := row.Scan(&cItem.Id, &cItem.CartId, &cItem.ProductId, &cItem.Quantity)
		if err != nil {
			return nil, err
		}
		Items = append(Items, cItem)
	}
	return Items, nil
}

func (c cartRepo) AddItem(cart *basket.CartItems) (id int, err error) {
	row, err := c.db.Query("INSERT INTO cart_items(cart_id,product_id, quantity) VALUES ($1,$2,$3) RETURNING id",
		cart.CartId, cart.CartId, cart.ProductId, cart.Quantity)
	if err != nil {
		return 0, basket.ErrCartItemCreationFailed
	}
	err = row.Scan(&id)
	if err != nil {
		return 0, basket.ErrIDScanFailed
	}
	return id, nil
}

func (c cartRepo) UpdateCartItem(cartId, quantity int) error {
	_, err := c.db.Exec("UPDATE cart_items SET quantity = quantity + $1 WHERE cart_id = $2",
		quantity, cartId)
	if err != nil {
		return basket.ErrCartUpdateFailed
	}
	return nil
}

func (c cartRepo) DeleteProduct(cartId, productId int) (id int, err error) {
	row, err := c.db.Query("delete from cart_items where cart_id = $1 AND product_id = $2 RETURNING id", cartId, productId)
	if err != nil {
		return 0, basket.ErrDeleteItemFailed
	}
	err = row.Scan(&id)
	if err != nil {
		return 0, basket.ErrIDScanFailed
	}
	return id, nil
}

func NewCartRepository(db *pgx.Conn) basket.CartRepository {
	return &cartRepo{db: db}
}
