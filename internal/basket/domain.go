package basket

type Basket struct {
	Id     int
	UserId int
}

type BasketItems struct {
	Id        int
	BasketId  int
	ProductId int
	Quantity  int
}
