package adapters

import (
	"database/sql"
	"golang-project-template/internal/order/domain"
)

type cartPostgresRepo struct {
	db *sql.DB
	f  domain.CartFactory
}

func newCartPostgresRepo(db *sql.DB) domain.CartItemsRepository {

	return &cartPostgresRepo{db: db}
}

type Items struct {
	id        int
	cartId    int
	productId int
	quantity  int
}

func (c *cartPostgresRepo) AddItem(productId int, userId int, quantity int) error {
	cardId := c.GetCardId(userId)
	query := `
	insert into cart_items(CartId,productId,quantity) values($1,$2,$3)
 `
	_, err := c.db.Query(query, cardId, productId, quantity)
	if err != nil {
		return err
	}
	return nil
}
func (c *cartPostgresRepo) GetCardId(userId int) (carId int) {
	query := `
	select id from Cart where userId=$1
`
	id, err := c.db.Query(query, userId)
	if err != nil {
		return 0
	}
	err = id.Scan(&carId)
	if err != nil {
		return 0
	}
	return carId
}

func (c *cartPostgresRepo) RemoveItem(productId int, userId int) error {
	cardId := c.GetCardId(userId)
	query := `
	delete from Cart where cardId=$1 and  productId=$2
`
	_, err := c.db.Query(query, cardId, productId)
	if err != nil {
		return err
	}
	return nil
}
func (c *cartPostgresRepo) GetAll(userId int) ([]*domain.CartItems, error) {
	cardId := c.GetCardId(userId)
	query := `
	select * from CartItems where CartId=$1
`
	rows, err := c.db.Query(query, cardId)
	if err != nil {
		return nil, err
	}
	var CartItems []*domain.CartItems
	for rows.Next() {
		var Item Items
		err = rows.Scan(&Item.id, &Item.cartId, &Item.productId, &Item.quantity)
		if err != nil {
			return nil, err
		}
		CartItems = append(CartItems, c.f.ParseModelToDomain(Item.cartId, Item.productId, Item.quantity))
	}
	return CartItems, err
}
func (c *cartPostgresRepo) RemoveAll(userId int) error {
	cardId := c.GetCardId(userId)
	query := `
	delete from cartItems where cartId=1$
`
	_, err := c.db.Query(query, cardId)
	if err != nil {
		return err
	}
	return nil
}
