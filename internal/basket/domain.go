package basket

type Basket struct {
	Id        int
	UserId    int
	Purchased bool
}

type BasketItems struct {
	Id        int
	BasketId  int
	ProductId int
	Quantity  int
}
