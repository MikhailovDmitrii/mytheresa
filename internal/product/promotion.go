package product

type Promotion interface {
	Apply(p *Product) *Product
	CanApply(p *Product) bool
	Percentage() int
}

func applyPercentageDiscount(p Product, percentage int) Product {
	p.Price.Final = p.Price.Original * (100 - percentage) / 100
	p.Price.DiscountPercentage = percentage

	return p
}
