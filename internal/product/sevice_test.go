package product_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"mytheresa-promotions/internal/infra"
	mock_infra "mytheresa-promotions/internal/mocks/infra"
	mock_product "mytheresa-promotions/internal/mocks/product"
	"mytheresa-promotions/internal/product"
	"testing"
)

func TestService_GetWithPromotions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_infra.NewMockRepository(ctrl)
	promotionMock := mock_product.NewMockPromotion(ctrl)

	promotionResultLimit := 2

	s := product.NewService(repositoryMock, promotionMock, promotionResultLimit)

	category := "some-category"
	priceLessThan := 0

	infraProducts := []*infra.Product{
		{
			Sku:      "1",
			Name:     "product-1",
			Category: "another-category",
			Price:    100,
		},
		{
			Sku:      "2",
			Name:     "product-2",
			Category: "some-category",
			Price:    200,
		},
	}

	repositoryMock.
		EXPECT().
		Search(category, priceLessThan, promotionResultLimit).
		Return(infraProducts, nil)

	products := []*product.Product{
		{
			Sku:      "1",
			Name:     "product-1",
			Category: "another-category",
			Price: product.Price{
				Original:           100,
				Final:              100,
				DiscountPercentage: 0,
				Currency:           product.CurrencyEuro,
			},
		},
		{
			Sku:      "2",
			Name:     "product-2",
			Category: "some-category",
			Price: product.Price{
				Original:           200,
				Final:              200,
				DiscountPercentage: 0,
				Currency:           product.CurrencyEuro,
			},
		},
	}

	expectedProducts := []*product.Product{
		products[0], // no promotion
		{ // some promotion applies but it doesn't matter here
			Sku:      "2",
			Name:     "product-2",
			Category: "some-category",
			Price: product.Price{
				Original:           200,
				Final:              150,
				DiscountPercentage: 25,
				Currency:           product.CurrencyEuro,
			},
		},
	}

	promotionMock.
		EXPECT().
		Apply(products[0]).
		Return(expectedProducts[0])

	promotionMock.
		EXPECT().
		Apply(products[1]).
		Return(expectedProducts[1])

	actualProducts, err := s.GetWithPromotions(category, priceLessThan)
	assert.NoError(t, err)
	assert.Equal(t, expectedProducts, actualProducts)

}

func TestService_GetWithPromotionsRepositoryError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_infra.NewMockRepository(ctrl)

	promotionResultLimit := 2

	s := product.NewService(repositoryMock, nil, promotionResultLimit)

	category := "some-category"
	priceLessThan := 0

	repositoryMock.
		EXPECT().
		Search(category, priceLessThan, promotionResultLimit).
		Return(nil, errors.New("repository error"))

	products, err := s.GetWithPromotions(category, priceLessThan)

	assert.Nil(t, products)
	assert.ErrorIs(t, err, product.InternalError)
}

func TestService_GetWithPromotionsEmptyPromotionSlice(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_infra.NewMockRepository(ctrl)

	promotionResultLimit := 2

	s := product.NewService(repositoryMock, nil, promotionResultLimit)

	category := "some-category"
	priceLessThan := 0

	infraProducts := []*infra.Product{
		{
			Sku:      "1",
			Name:     "product-1",
			Category: "another-category",
			Price:    100,
		},
		{
			Sku:      "2",
			Name:     "product-2",
			Category: "some-category",
			Price:    200,
		},
	}

	repositoryMock.
		EXPECT().
		Search(category, priceLessThan, promotionResultLimit).
		Return(infraProducts, nil)

	products, err := s.GetWithPromotions(category, priceLessThan)

	expected := []*product.Product{
		{
			Sku:      "1",
			Name:     "product-1",
			Category: "another-category",
			Price: product.Price{
				Original:           100,
				Final:              100,
				DiscountPercentage: 0,
				Currency:           product.CurrencyEuro,
			},
		},
		{
			Sku:      "2",
			Name:     "product-2",
			Category: "some-category",
			Price: product.Price{
				Original:           200,
				Final:              200,
				DiscountPercentage: 0,
				Currency:           product.CurrencyEuro,
			},
		},
	}

	assert.NoError(t, err)
	assert.Equal(t, expected, products)
}
