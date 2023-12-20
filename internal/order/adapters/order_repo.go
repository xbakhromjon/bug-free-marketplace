package adapters

import (
	"database/sql"
	"fmt"
	"github.com/jackc/pgx"
	"golang-project-template/internal/order/domain"
)

type orderPostgresRepo struct {
	db *pgx.Conn
}

func newOrderPostgresRepo(db *pgx.Conn) domain.OrderRepository {
	return &orderPostgresRepo{db: db}
}

func (o *orderPostgresRepo) CreateOrder(order domain.Order) error {
	_, err := o.db.Exec("INSERT INTO orders (number, cart_id, total_price, status, created_at, updated_at) values ($1, $2, $3, $4, $5, $6)",
		order.Number, order.CartId, order.TotalPrice, order.Status, order.CreatedAt, order.UpdatedAt)

	return err
}

func (o *orderPostgresRepo) GetOrderByID(orderID int) (domain.Order, error) {
	var order domain.Order
	err := o.db.QueryRow(`select number,cart_id,total_price,status,created_at,updated_at from orders where order_Id = $1`, orderID).
		Scan(&order.Id, &order.Number, &order.TotalPrice, &order.Status, &order.CreatedAt, &order.UpdatedAt)
	if err != sql.ErrNoRows {
		return order, domain.ErrOrderNotFound
	}

	return order, nil
}

func (o *orderPostgresRepo) GetAllOrders() ([]domain.Order, error) {
	rows, err := o.db.Query(`SELECT number,cart_id,total_price,status,created_at,updated_at FROM orders`)
	if err != nil {
		return nil, err
	}

	orders := make([]domain.Order, 0)

	for rows.Next() {
		var order domain.Order
		if err := rows.Scan(&order.Id, order.Number, order.TotalPrice, order.Status, order.CreatedAt, order.UpdatedAt); err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, rows.Err()
}

func (o *orderPostgresRepo) UpdateStatus(orderID int, newStatus string) error {

	query := fmt.Sprintf("UPDATE orders SET status = $1 WHERE id=%d")
	_, err := o.db.Exec(query, newStatus, orderID)
	if err != nil {
		return err
	}
	return nil
}
