package product

type skuPromotion struct {
	sku        string
	percentage int
}

func NewSkuPromotion(sku string, percentage int) Promotion {
	return &skuPromotion{
		sku:        sku,
		percentage: percentage,
	}
}

func (c *skuPromotion) Apply(p *Product) *Product {
	if !c.CanApply(p) {
		return p
	}

	productWithDiscount := applyPercentageDiscount(*p, c.percentage)

	return &productWithDiscount
}

func (c *skuPromotion) CanApply(p *Product) bool {
	return p.Sku == c.sku
}

func (c *skuPromotion) Percentage() int {
	return c.percentage
}
