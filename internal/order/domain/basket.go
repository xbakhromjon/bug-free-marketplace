package domain

type Basket struct {
	Id        int  `json:"id"`
	UserId    int  `json:"userId"`
	Purchased bool `json:"purchased"`
}

type BasketWithItems struct {
	Basket
	Items []BasketItems
}
