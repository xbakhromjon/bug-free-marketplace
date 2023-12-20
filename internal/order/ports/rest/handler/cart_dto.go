package handler

type cart struct {
	userId int `json:"user_id"`
}
type cartItems struct {
	cartId    int `json:"cart_id"`
	productId int `json:"product_id"`
	quantity  int `json:"quantity"`
}
