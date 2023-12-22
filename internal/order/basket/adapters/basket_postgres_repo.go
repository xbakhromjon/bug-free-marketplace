package adapters

import (
	"github.com/jackc/pgx"
	"golang-project-template/internal/order/basket/domain"
)

type basketPostgresRepo struct {
	db *pgx.Conn
}

func NewBasketRepository(db *pgx.Conn) domain.BasketRepository {
	return &basketPostgresRepo{db: db}
}

func (b basketPostgresRepo) CreateBasket(userId int) (id int, err error) {
	row := b.db.QueryRow("INSERT INTO cart(user_id,purchased) VALUES($1) RETURNING id;", userId)
	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func (b basketPostgresRepo) AddItem(cart *domain.BasketItems) (id int, err error) {
	row := b.db.QueryRow("INSERT INTO cart_items(cart_id,product_id, quantity) VALUES ($1,$2,$3) RETURNING id",
		cart.CartId, cart.ProductId, cart.Quantity)
	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
