package domain

type ShopRepository interface {
	Save(shop NewShop) (int, error)
	CheckShopNameExists(string) (bool, error)
	FindShopById(int) (*Shop, error)
	FindAllShops(int, int, string) ([]Shop, error)
}

type UserRepo interface {
	UserExists(id int) (bool, error)
}

type ProductRepository interface {
	Save(product *Product) (int, error)
	FindById(id int) (*Product, error)
	FindAllByShopId(shopId int) ([]*Product, error)
}

const (
	ErrProductNotFound = Err("product not found")
	ErrEmptyShopName   = Err("shop name can not be empty")
	ErrInvalidShopName = Err("shop name max length must be 255 characters")
	ErrShopNameExists  = Err("This shop name already exists")
	ErrUserNotExists   = Err("No such user")
	ErrShopNotFound    = Err("Shop not found")
	ErrInvalidLimit    = Err("Limit must be between 1 and 100")
	ErrInvalidOffset   = Err("offset must be non-negative")
	ErrInvalidSearch   = Err("too long search string")
)

type Err string

func (e Err) Error() string {
	return string(e)
}
