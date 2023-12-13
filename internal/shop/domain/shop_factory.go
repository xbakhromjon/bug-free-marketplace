package domain

type ShopFactory struct {
	maxNameLen int
}

func NewShopFactory(maxNameLen int) ShopFactory {
	return ShopFactory{maxNameLen: maxNameLen}
}

func (f *ShopFactory) Validate(shop NewShop) error {
	if shop.Name == "" {
		return ErrEmptyShopName
	}

	if len(shop.Name) > f.maxNameLen {
		return ErrInvalidShopName
	}
	return nil
}
