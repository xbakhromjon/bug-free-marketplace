package domain

type Cart struct {
	id     int
	userId int
}

type CartItems struct {
	id        int
	cartId    int
	productId int
	quantity  int
}
