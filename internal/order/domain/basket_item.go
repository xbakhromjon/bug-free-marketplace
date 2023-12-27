package domain

type BasketItems struct {
	Id        int `json:"id"`
	BasketId  int `json:"basketId"`
	ProductId int `json:"productId"`
	Quantity  int `json:"quantity"`
}
