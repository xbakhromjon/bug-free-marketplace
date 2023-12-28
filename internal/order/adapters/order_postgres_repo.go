package adapters

import (
	"database/sql"
	"fmt"
	"github.com/jackc/pgx"
	"golang-project-template/internal/order/domain"
)

type orderRepo struct {
	db *pgx.Conn
}

func NewOrderRepository(db *pgx.Conn) domain.OrderRepository {
	return &orderRepo{db: db}
}

func (o *orderRepo) CreateOrder(order domain.Order) error {
	_, err := o.db.Exec(`INSERT INTO orders (number, basket_id, total_price, status, created_at, updated_at) values ($1, $2, $3, $4, $5, $6)`,
		order.Number, order.BasketID, order.TotalPrice, order.Status, order.CreatedAt, order.UpdatedAt)

	return err
}

func (o *orderRepo) GetOrderByID(orderID int) (domain.Order, error) {
	var order domain.Order
	err := o.db.QueryRow(`SELECT id, number, basket_id, total_price, status, created_at, updated_at FROM orders WHERE id = $1`, orderID).
		Scan(&order.ID, &order.Number, &order.BasketID, &order.TotalPrice, &order.Status, &order.CreatedAt, &order.UpdatedAt)

	if err != sql.ErrNoRows {
		return order, domain.ErrOrderNotFound
	}
	return order, err
}

func (o *orderRepo) GetPaginatedOrders(offset, limit int) ([]domain.Order, error) {
	query := `
		SELECT id, number, basket_id, total_price, status, created_at, updated_at
		FROM orders
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := o.db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	orders := make([]domain.Order, 0)

	for rows.Next() {
		var order domain.Order

		err := rows.Scan(&order.ID, &order.Number, &order.BasketID, &order.TotalPrice, &order.Status, &order.CreatedAt, &order.UpdatedAt)
		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, rows.Err()
}

func (o *orderRepo) UpdateStatusOrder(orderID int, newStatus string) error {

	query := fmt.Sprintf("UPDATE orders SET status = $2 WHERE id=$1")
	_, err := o.db.Exec(query, orderID, newStatus)

	return err
}
