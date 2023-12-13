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

type NewBasket struct {
	Id        int
	CartId    string
	ProductId int
	Quantity  int
}

//
//type Order struct {
//	Id         int
//	Number     string
//	BasketId   int
//	TotalPrice int
//	Status     string
//	CreatedAt  time.Time
//	UpdatedAt  time.Time
//}
