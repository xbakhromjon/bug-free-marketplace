package adapters

import (
	"github.com/jackc/pgx"
	"golang-project-template/internal/order/basket/domain"
)

type basketPostgresRepo struct {
	db *pgx.Conn
}

func (b basketPostgresRepo) CreateBasket(userId int) (id int, err error) {
	row := b.db.QueryRow("INSERT INTO cart(user_id,purchased) VALUES($1,true) RETURNING id;", userId)
	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func NewBasketRepository(db *pgx.Conn) domain.BasketRepository {
	return &basketPostgresRepo{db: db}
}
