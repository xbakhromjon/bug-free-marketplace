package domain

import "time"

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

func (f *ShopFactory) ParseModelToDomain(
	id int,
	name string,
	ownerId int,
	createdAt time.Time,
	updatedAt time.Time,
) *Shop {
	return &Shop{
		id:        id,
		name:      name,
		ownerId:   ownerId,
		createdAt: createdAt.Format(time.RFC1123),
		updatedAt: updatedAt.Format(time.RFC1123),
	}
}
