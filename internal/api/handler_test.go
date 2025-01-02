package api_test

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"mytheresa-promotions/internal/api"
	mock_product "mytheresa-promotions/internal/mocks/product"
	"mytheresa-promotions/internal/product"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	s := mock_product.NewMockService(ctrl)
	h := api.NewHandler(s)

	sku := "1234"
	req := httptest.NewRequest(http.MethodGet, "/products/"+sku, nil)
	req.SetPathValue("sku", sku)

	s.
		EXPECT().
		FindBySku(sku).
		Return(&product.Product{
			Sku: sku,
		}, nil)

	w := httptest.NewRecorder()
	h.GetOne()(w, req)

	var p *product.Product
	err := json.NewDecoder(w.Body).Decode(&p)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, p.Sku, sku)
}

func TestHandler_GetWithoutSku(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	s := mock_product.NewMockService(ctrl)
	h := api.NewHandler(s)

	sku := "1234"
	req := httptest.NewRequest(http.MethodGet, "/products/"+sku, nil)

	w := httptest.NewRecorder()
	h.GetOne()(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestHandler_GetNotFoundSku(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	s := mock_product.NewMockService(ctrl)
	h := api.NewHandler(s)

	sku := "1234"
	req := httptest.NewRequest(http.MethodGet, "/products/"+sku, nil)
	req.SetPathValue("sku", sku)

	s.
		EXPECT().
		FindBySku(sku).
		Return(nil, product.NotFoundError)

	w := httptest.NewRecorder()
	h.GetOne()(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestHandler_GetServiceInternalError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	s := mock_product.NewMockService(ctrl)
	h := api.NewHandler(s)

	sku := "1234"
	req := httptest.NewRequest(http.MethodGet, "/products/"+sku, nil)
	req.SetPathValue("sku", sku)

	s.
		EXPECT().
		FindBySku(sku).
		Return(nil, product.InternalError)

	w := httptest.NewRecorder()
	h.GetOne()(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestHandler_Promotions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	s := mock_product.NewMockService(ctrl)
	h := api.NewHandler(s)

	sku := "1234"
	category := "category-in-filters"
	priceLessThan := 10000
	req := httptest.NewRequest(
		http.MethodGet,
		fmt.Sprintf("/products/%s?category=%s&priceLessThan=%d", sku, category, priceLessThan),
		nil)
	req.SetPathValue("sku", sku)

	s.
		EXPECT().
		GetWithPromotions(category, priceLessThan).
		Return([]*product.Product{
			{
				Sku:      sku,
				Name:     "product-1",
				Category: "category-in-filters",
				Price: product.Price{
					Original:           100,
					Final:              20,
					DiscountPercentage: 80,
					Currency:           product.CurrencyEuro,
				},
			},
		}, nil)

	w := httptest.NewRecorder()
	h.Promotions()(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []*api.ProductResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)

	assert.Len(t, response, 1)
	assert.Equal(t, 100, response[0].Price.Original)
	assert.Equal(t, 20, response[0].Price.Final)
	assert.Equal(t, "80%", *response[0].Price.DiscountPercentage)

}
