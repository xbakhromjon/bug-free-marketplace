package adapters

import (
	"database/sql"
	"errors"
	"github.com/jackc/pgx"
	domain2 "golang-project-template/internal/order/domain"
)

type cartRepo struct {
	db *pgx.Conn
	cf domain2.CartFactory
}

func (c cartRepo) GetCardItem(cartId, productId int) (*domain2.CartItems, error) {
	row := c.db.QueryRow("SELECT id, cart_id, product_id, quantity from card_items WHERE cart_id = ? AND product_id = ?", cartId, productId)
	var cartItems domain2.CartItems
	err := row.Scan(&cartItems.Id, &cartItems.CartId, &cartItems.ProductId, &cartItems.Quantity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("CartItem not found")
		}
		return nil, err
	}
	return &cartItems, nil
}

func (c cartRepo) CreateCardItem(cart *domain2.CartItems) (int, error) {
	_, err := c.db.Exec("INSERT INTO cart_items(cart_id,product_id, quantity) VALUES (?,?,?,?)",
		cart.CartId, cart.CartId, cart.ProductId, cart.Quantity)
	if err != nil {
		return 0, err
	}
	return cart.Id, nil
}

func (c cartRepo) Create(cart *domain2.Cart) error {
	_, err := c.db.Exec("INSERT INTO cart(user_id)VALUES (?)", cart.UserId)
	if err != nil {
		return err
	}
	return nil
}

func (c cartRepo) GetById(id int) (*domain2.CartItems, error) {
	row := c.db.QueryRow("select id,cart_id,product_id,quantity from cart_items where id = ?", id)
	var basket domain2.CartItems
	err := row.Scan(&basket.Id, &basket.CartId, &basket.ProductId, &basket.Quantity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("Basket not found")
		}
		return nil, err
	}
	return &basket, nil
}

func (c cartRepo) GetByUserId(userID int) (*domain2.Cart, error) {
	row := c.db.QueryRow("select id,user_id from carts where user_id = ?", userID)
	var cart domain2.Cart
	err := row.Scan(&cart.Id, &cart.UserId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("Cart not found")
		}
		return nil, err
	}
	return &cart, err
}

func (c cartRepo) UpdateCartItem(userId, productId, quantity int) error {
	_, err := c.db.Exec("UPDATE cart_items SET quantity = quantity + ? WHERE user_id = ? AND product_id = ?",
		quantity, userId, productId)
	if err != nil {
		return err
	}
	return nil
}

func (c cartRepo) DeleteProduct(cartId, productId int) error {
	_, err := c.db.Exec("delete from cart_items where cart_id = ? AND product_id = ?", cartId, productId)
	return err
}

func NewCartRepository(db *pgx.Conn) domain2.CartRepository {
	return &cartRepo{db: db}
}
