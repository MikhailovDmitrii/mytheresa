package product_test

import (
	"github.com/stretchr/testify/assert"
	"mytheresa-promotions/internal/product"
	"testing"
)

func TestCompositePromotion_Apply(t *testing.T) {

	promotion := product.NewCompositePromotion(
		product.NewCategoryPromotion("some-category", 20),
		product.NewCategoryPromotion("some-category", 30),
		product.NewSkuPromotion("1234", 10),
	)

	p := promotion.Apply(&product.Product{
		Sku:      "1234",
		Category: "some-category",
		Price: product.Price{
			Original:           100,
			Final:              100,
			DiscountPercentage: 0,
			Currency:           product.CurrencyEuro,
		},
	})

	assert.Equal(t, 70, p.Price.Final)
	assert.Equal(t, 30, p.Price.DiscountPercentage)
}

func TestCompositePromotion_CanApply(t *testing.T) {
	promotion := product.NewCompositePromotion(
		product.NewCategoryPromotion("some-category", 20),
		product.NewSkuPromotion("1234", 10),
	)

	assert.True(t, promotion.CanApply(&product.Product{Category: "some-category"}))
	assert.True(t, promotion.CanApply(&product.Product{Category: "another-category", Sku: "1234"}))
	assert.False(t, promotion.CanApply(&product.Product{Category: "another-category", Sku: "another-sku"}))
}
