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
