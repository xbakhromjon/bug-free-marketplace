package domain

type Cart struct {
	Id     int
	UserId int
}

type CartItems struct {
	Id        int
	CartId    int
	ProductId int
	Quantity  int
}
