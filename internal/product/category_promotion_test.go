package product_test

import (
	"github.com/stretchr/testify/assert"
	"mytheresa-promotions/internal/product"
	"testing"
)

func TestCategoryPromotion_Apply(t *testing.T) {
	promotion := product.NewCategoryPromotion("some-category", 10)

	p := promotion.Apply(&product.Product{
		Category: "some-category",
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

func TestCategoryPromotion_ApplyWrongCategory(t *testing.T) {
	promotion := product.NewCategoryPromotion("some-category", 10)

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

func TestCategoryPromotion_CanApply(t *testing.T) {
	promotion := product.NewCategoryPromotion("some-category", 10)

	assert.True(t, promotion.CanApply(&product.Product{Category: "some-category"}))
	assert.False(t, promotion.CanApply(&product.Product{Category: "another-category"}))
}
