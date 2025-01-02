package product

type compositePromotion struct {
	promotions []Promotion
}

func NewCompositePromotion(promotions ...Promotion) Promotion {
	return &compositePromotion{
		promotions: promotions,
	}
}

func (cp *compositePromotion) Apply(p *Product) *Product {
	if len(cp.promotions) == 0 {
		return p
	}

	biggestPercentage := 0
	promotionIndex := -1
	for i, promotion := range cp.promotions {
		if promotion.CanApply(p) && promotion.Percentage() > biggestPercentage {
			promotionIndex = i
			biggestPercentage = promotion.Percentage()
		}
	}

	if promotionIndex == -1 {
		return p
	}

	p1 := *p
	productWithPromotion := &p1
	productWithPromotion = cp.promotions[promotionIndex].Apply(productWithPromotion)

	return productWithPromotion
}

func (cp *compositePromotion) CanApply(p *Product) bool {
	canApply := false
	for _, promotion := range cp.promotions {
		canApply = canApply || promotion.CanApply(p)
	}

	return canApply
}

// Percentage doesn't make sense without product here because for each product it varies
func (cp *compositePromotion) Percentage() int {
	return 0
}
