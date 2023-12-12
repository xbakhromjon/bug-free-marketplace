package domain

type ShopFactory struct {
	maxNameLen      int
	maxSearchLength int
}

func NewShopFactory(maxNameLen, maxSearchLength int) ShopFactory {
	return ShopFactory{maxNameLen: maxNameLen, maxSearchLength: maxSearchLength}
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

func (f *ShopFactory) GetAllShopsInputValidate(limit, offset int, search string) error {
	if limit <= 0 || limit > 100 {
		return ErrInvalidLimit
	}

	if offset < 0 {
		return ErrInvalidOffset
	}

	if len(search) > f.maxSearchLength {
		return ErrInvalidSearch
	}

	return nil
}
