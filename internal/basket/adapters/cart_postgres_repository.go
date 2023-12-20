package adapters

import (
	"errors"
	"github.com/jackc/pgx"
	basket "golang-project-template/internal/basket/domain"
)

type cartRepo struct {
	db *pgx.Conn
}

func (c cartRepo) GetCartItemByCartIdAndProductId(cartId, productId int) (*basket.CartItems, error) {
	row := c.db.QueryRow("SELECT * from card_items WHERE cart_id = $1 and product_id = $2",
		cartId, productId)
	var cItems basket.CartItems
	err := row.Scan(&cItems.Id, &cItems.CartId, &cItems.ProductId, &cItems.Quantity)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, basket.ErrCartItemNotFound
		}
	}
	return &cItems, nil
}

func (c cartRepo) GetCardItem(cartId int) ([]*basket.CartItems, error) {
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

func (c cartRepo) CreateCardItem(cart *basket.CartItems) (int, error) {
	_, err := c.db.Exec("INSERT INTO cart_items(cart_id,product_id, quantity) VALUES ($1,$2,$3) RETURNING id",
		cart.CartId, cart.CartId, cart.ProductId, cart.Quantity)
	if err != nil {
		return 0, basket.ErrCartItemCreationFailed
	}
	return cart.Id, nil
}

func (c cartRepo) Create(cart *basket.Cart) (int, error) {
	_, err := c.db.Exec("INSERT INTO cart(user_id)VALUES ($1) RETURNING id", cart.UserId)
	if err != nil {
		return 0, basket.ErrCartCreationFailed
	}
	return cart.Id, nil
}

func (c cartRepo) GetCart(id int) (*basket.Cart, error) {
	row := c.db.QueryRow("select id,user_id from cart where id = $1", id)
	var cart basket.Cart
	err := row.Scan(&cart.Id, &cart.UserId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, basket.ErrCartNotFound
		}
		return nil, err
	}
	return &cart, nil
}

func (c cartRepo) GetByUserId(userID int) (*basket.Cart, error) {
	row := c.db.QueryRow("select id,user_id from carts where user_id = $1", userID)
	var cart basket.Cart
	err := row.Scan(&cart.Id, &cart.UserId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, basket.ErrCartNotFound
		}
		return nil, err
	}
	return &cart, err
}

func (c cartRepo) UpdateCartItem(cartId, quantity int) error {
	_, err := c.db.Exec("UPDATE cart_items SET quantity = quantity + $1 WHERE cart_id = $2",
		quantity, cartId)
	if err != nil {
		return basket.ErrCartUpdateFailed
	}
	return nil
}

func (c cartRepo) DeleteProduct(cartId, productId int) error {
	_, err := c.db.Exec("delete from cart_items where cart_id = $1 AND product_id = $2", cartId, productId)
	return err
}

func NewCartRepository(db *pgx.Conn) basket.CartRepository {
	return &cartRepo{db: db}
}
