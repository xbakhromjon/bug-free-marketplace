package domain

type ProductFactory struct {
}

func (p *ProductFactory) CreateNewSearchModel(name string, priceFrom int, priceTo int) *ProductSearchModel {

	return &ProductSearchModel{
		name:      name,
		priceFrom: priceFrom,
		priceTo:   priceTo,
	}
}

func (p *ProductFactory) NewProduct(name string, price int, shopId int) *Product {

	return &Product{
		name:   name,
		price:  price,
		shopId: shopId,
	}
}

func (p *ProductFactory) CreateExistProduct(id int, name string, price int, shopId int) *Product {

	return &Product{
		id:     id,
		name:   name,
		price:  price,
		shopId: shopId,
	}
}
