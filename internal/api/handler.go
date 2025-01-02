package api

import (
	"encoding/json"
	"errors"
	"mytheresa-promotions/internal/product"
	"net/http"
	"strconv"
)

type Handler struct {
	service product.Service
}

func NewHandler(s product.Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) GetOne() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sku := r.PathValue("sku")
		if sku == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		var p *product.Product
		var err error

		p, err = h.service.FindBySku(sku)

		if errors.Is(err, product.NotFoundError) {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		responseWithSingleProduct(p, w)
	}
}

func (h *Handler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var createRequest ProductCreateRequest
		err := json.NewDecoder(r.Body).Decode(&createRequest)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		products := convertCreateRequestToModel(createRequest)

		if err = h.service.BulkCreate(products); err != nil {
			if errors.Is(err, product.DuplicateSkuError) {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		responseWithSlice(products, w)
	}
}

func (h *Handler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		sku := r.PathValue("sku")
		if sku == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var updateRequest ProductUpdateRequest
		err := json.NewDecoder(r.Body).Decode(&updateRequest)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		p := convertUpdateRequestToModel(sku, updateRequest)

		if p, err = h.service.Update(p); err != nil {
			if errors.Is(err, product.NotFoundError) {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		responseWithSingleProduct(p, w)
	}
}

func (h *Handler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		sku := r.PathValue("sku")
		if sku == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err := h.service.Delete(sku)

		if errors.Is(err, product.NotFoundError) {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		responseWithJson(nil, w)
	}
}

func (h *Handler) Promotions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var priceLessThan int
		var err error

		category := r.URL.Query().Get("category")
		priceLessThanParam := r.URL.Query().Get("priceLessThan")
		if priceLessThanParam != "" {
			priceLessThan, err = strconv.Atoi(priceLessThanParam)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}

		products, err := h.service.GetWithPromotions(category, priceLessThan)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		responseWithSlice(products, w)
	}
}

func responseWithSingleProduct(p *product.Product, w http.ResponseWriter) {
	var response ProductResponse

	if p != nil {
		response = convertToResponse(p)
	}

	responseWithJson(response, w)
}

func responseWithSlice(products []*product.Product, w http.ResponseWriter) {
	convertedProducts := make([]ProductResponse, 0, len(products))

	if len(products) != 0 {
		for _, p := range products {
			convertedProducts = append(convertedProducts, convertToResponse(p))
		}
	}

	responseWithJson(convertedProducts, w)
}

func responseWithJson(entities any, w http.ResponseWriter) {
	jsonResponse, err := json.Marshal(entities)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
