package product

import "mytheresa-promotions/internal/infra"

func convertFromInfraToModel(product *infra.Product) *Product {
	return &Product{
		Sku:      product.Sku,
		Name:     product.Name,
		Category: product.Category,
		Price: Price{
			Original:           product.Price,
			Final:              product.Price,
			DiscountPercentage: 0,
			Currency:           CurrencyEuro,
		},
	}
}

func convertSliceFromModelToInfra(products []*Product) []*infra.Product {
	result := make([]*infra.Product, 0, len(products))

	for _, p := range products {
		result = append(result, convertModelToInfra(p))
	}

	return result
}

func convertSliceFromInfraToModel(products []*infra.Product) []*Product {
	result := make([]*Product, 0, len(products))

	for _, p := range products {
		result = append(result, convertFromInfraToModel(p))
	}

	return result
}

func convertModelToInfra(p *Product) *infra.Product {
	return &infra.Product{
		Sku:      p.Sku,
		Name:     p.Name,
		Category: p.Category,
		Price:    p.Price.Original,
	}
}
