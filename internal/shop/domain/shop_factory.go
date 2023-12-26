package domain

import "time"

type ShopFactory struct {
	maxNameLen      int
	maxSearchLength int
}

func NewShopFactory(maxNameLen, maxSearchLength int) ShopFactory {
	return ShopFactory{maxNameLen: maxNameLen, maxSearchLength: maxSearchLength}
}

func (f *ShopFactory) Validate(name string) error {
	if name == "" {
		return ErrEmptyShopName
	}

	if len(name) > f.maxNameLen {
		return ErrInvalidShopName
	}
	return nil
}

func (f *ShopFactory) NewShop(name string, ownerId int) (*Shop, error) {
	err := f.Validate(name)
	if err != nil {
		return &Shop{}, err
	}
	return &Shop{name: name, ownerId: ownerId, createdAt: time.Now().UTC(), updatedAt: time.Now().UTC()}, nil
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
) Shop {
	return Shop{
		id:        id,
		name:      name,
		ownerId:   ownerId,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}
