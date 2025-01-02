package api

type PriceResponse struct {
	Original           int     `json:"original"`
	Final              int     `json:"final"`
	DiscountPercentage *string `json:"discount_percentage"`
	Currency           string  `json:"currency"`
}

type ProductResponse struct {
	Sku      string        `json:"sku"`
	Name     string        `json:"name"`
	Category string        `json:"category"`
	Price    PriceResponse `json:"price"`
}

type ProductCreateRequest struct {
	Products []struct {
		Sku      string `json:"sku"`
		Name     string `json:"name"`
		Category string `json:"category"`
		Price    int    `json:"price"`
	} `json:"products"`
}

type ProductUpdateRequest struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    int    `json:"price"`
}
