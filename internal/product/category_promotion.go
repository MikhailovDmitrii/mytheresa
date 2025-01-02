package product

type categoryPromotion struct {
	category   string
	percentage int
}

func NewCategoryPromotion(category string, percentage int) Promotion {
	return &categoryPromotion{
		category:   category,
		percentage: percentage,
	}
}

func (c *categoryPromotion) Apply(p *Product) *Product {
	if !c.CanApply(p) {
		return p
	}

	productWithDiscount := applyPercentageDiscount(*p, c.percentage)

	return &productWithDiscount
}

func (c *categoryPromotion) CanApply(p *Product) bool {
	return p.Category == c.category
}

func (c *categoryPromotion) Percentage() int {
	return c.percentage
}
