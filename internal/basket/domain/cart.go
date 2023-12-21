package domain

type Cart struct {
	Id     int `json:"id"`
	UserId int `json:"userId"`
}

type CartItems struct {
	Id        int `json:"id"`
	CartId    int `json:"cartId"`
	ProductId int `json:"productId"`
	Quantity  int `json:"quantity"`
}

type NewCart struct {
	UserId int
}
