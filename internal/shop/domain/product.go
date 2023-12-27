package domain

type Product struct {
	Id     int
	Name   string
	Price  int
	ShopId int
	Count  int
}

type NewProduct struct {
	Name   string
	Price  int
	ShopId int
}

type ProductSearchModel struct {
	name      string `json:"name"`
	priceFrom int    `json:"priceFrom"`
	priceTo   int    `json:"priceTo"`
}

func (p *ProductSearchModel) GetName() (string, bool) {
	if p.name == "" {
		return "", false
	}
	return p.name, true
}

func (p *ProductSearchModel) GetPriceFrom() (int, bool) {
	if p.priceFrom == 0 {
		return 0, false
	}
	return p.priceFrom, true
}

func (p *ProductSearchModel) GetPriceTo() (int, bool) {
	if p.priceTo == 0 {
		return 0, false
	}
	return p.priceTo, true
}
