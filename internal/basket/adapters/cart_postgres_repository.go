package adapters

import (
	"database/sql"
	"errors"
	"github.com/jackc/pgx"
	basket "golang-project-template/internal/basket/domain"
)

type cartRepo struct {
	db *pgx.Conn
}

func (c cartRepo) GetCardItem(cartId, productId int) (*basket.CartItems, error) {
	row := c.db.QueryRow("SELECT id, cart_id, product_id, quantity from card_items WHERE cart_id = ? AND product_id = ?", cartId, productId)
	var cartItems basket.CartItems
	err := row.Scan(&cartItems.Id, &cartItems.CartId, &cartItems.ProductId, &cartItems.Quantity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, basket.ErrCartItemNotFound
		}
	}
	return &cartItems, nil
}

func (c cartRepo) CreateCardItem(cart *basket.CartItems) (int, error) {
	_, err := c.db.Exec("INSERT INTO cart_items(cart_id,product_id, quantity) VALUES (?,?,?,?) RETURNING id",
		cart.CartId, cart.CartId, cart.ProductId, cart.Quantity)
	if err != nil {
		return 0, basket.ErrCartItemCreationFailed
	}
	return cart.Id, nil
}

func (c cartRepo) Create(cart *basket.Cart) (int, error) {
	_, err := c.db.Exec("INSERT INTO cart(user_id)VALUES (?) RETURNING id", cart.UserId)
	if err != nil {
		return 0, basket.ErrCartCreationFailed
	}
	return cart.Id, nil
}

func (c cartRepo) GetCart(id int) (*basket.Cart, error) {
	row := c.db.QueryRow("select id,user_id from cart where id = ?", id)
	var cart basket.Cart
	err := row.Scan(&cart.Id, &cart.UserId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, basket.ErrCartNotFound
		}
		return nil, err
	}
	return &cart, nil
}

func (c cartRepo) GetByUserId(userID int) (*basket.Cart, error) {
	row := c.db.QueryRow("select id,user_id from carts where user_id = ?", userID)
	var cart basket.Cart
	err := row.Scan(&cart.Id, &cart.UserId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, basket.ErrCartNotFound
		}
		return nil, err
	}
	return &cart, err
}

func (c cartRepo) UpdateCartItem(cartId, productId, quantity int) error {
	_, err := c.db.Exec("UPDATE cart_items SET quantity = ? WHERE cart_id = ? AND product_id = ?",
		quantity, cartId, productId)
	if err != nil {
		return basket.ErrCartUpdateFailed
	}
	return nil
}

func (c cartRepo) DeleteProduct(cartId, productId int) error {
	_, err := c.db.Exec("delete from cart_items where cart_id = ? AND product_id = ?", cartId, productId)
	return err
}

func NewCartRepository(db *pgx.Conn) basket.CartRepository {
	return &cartRepo{db: db}
}
