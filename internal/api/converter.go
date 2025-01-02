package api

import (
	"fmt"
	"mytheresa-promotions/internal/product"
)

func convertToResponse(p *product.Product) ProductResponse {

	var percentage *string
	if p.Price.DiscountPercentage != 0 {
		percentageStr := fmt.Sprintf("%d%%", p.Price.DiscountPercentage)
		percentage = &percentageStr
	}

	return ProductResponse{
		Sku:      p.Sku,
		Name:     p.Name,
		Category: p.Category,
		Price: PriceResponse{
			Original:           p.Price.Original,
			Final:              p.Price.Final,
			DiscountPercentage: percentage,
			Currency:           p.Price.Currency,
		},
	}
}

func convertCreateRequestToModel(request ProductCreateRequest) []*product.Product {
	models := make([]*product.Product, 0, len(request.Products))
	for _, p := range request.Products {
		models = append(models, &product.Product{
			Sku:      p.Sku,
			Name:     p.Name,
			Category: p.Category,
			Price: product.Price{
				Original:           p.Price,
				Final:              p.Price,
				DiscountPercentage: 0,
				Currency:           product.CurrencyEuro,
			},
		})
	}

	return models
}

func convertUpdateRequestToModel(sku string, request ProductUpdateRequest) *product.Product {
	return &product.Product{
		Sku:      sku,
		Name:     request.Name,
		Category: request.Category,
		Price: product.Price{
			Original:           request.Price,
			Final:              request.Price,
			DiscountPercentage: 0,
			Currency:           product.CurrencyEuro,
		},
	}
}
