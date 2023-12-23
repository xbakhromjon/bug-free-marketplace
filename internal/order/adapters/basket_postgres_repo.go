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
	row := b.db.QueryRow("INSERT INTO basket(user_id,purchased) VALUES ($1,$2) RETURNING id;", userId)
	err = row.Scan(row)
	if err != nil {
		return 0, err
	}
	return id, nil
}
