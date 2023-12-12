package domain

type CartFactory struct {
}

func (f *CartFactory) ParseModelToDomain(carId int, productId int, quantity int) *CartItems {

	return &CartItems{
		cartId:    carId,
		productId: productId,
		quantity:  quantity,
	}
}
