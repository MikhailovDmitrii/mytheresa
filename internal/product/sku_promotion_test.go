package product_test

import (
	"github.com/stretchr/testify/assert"
	"mytheresa-promotions/internal/product"
	"testing"
)

func TestSkuPromotion_Apply(t *testing.T) {
	promotion := product.NewSkuPromotion("1234", 10)

	p := promotion.Apply(&product.Product{
		Sku: "1234",
		Price: product.Price{
			Original:           100,
			Final:              100,
			DiscountPercentage: 0,
			Currency:           product.CurrencyEuro,
		},
	})

	assert.Equal(t, 90, p.Price.Final)
	assert.Equal(t, 10, p.Price.DiscountPercentage)
}

func TestSkuPromotion_ApplyWrongSku(t *testing.T) {
	promotion := product.NewSkuPromotion("1234", 10)

	p := promotion.Apply(&product.Product{
		Category: "another-category",
		Price: product.Price{
			Original:           100,
			Final:              100,
			DiscountPercentage: 0,
			Currency:           product.CurrencyEuro,
		},
	})

	assert.Equal(t, 100, p.Price.Final)
	assert.Equal(t, 0, p.Price.DiscountPercentage)
}

func TestSkuPromotion_CanApply(t *testing.T) {
	promotion := product.NewSkuPromotion("1234", 10)

	assert.True(t, promotion.CanApply(&product.Product{Sku: "1234"}))
	assert.False(t, promotion.CanApply(&product.Product{Sku: "another-sku"}))
}
