package product

const CurrencyEuro = "EUR"

type Price struct {
	Original           int
	Final              int
	DiscountPercentage int
	Currency           string
}

type Product struct {
	Sku      string
	Name     string
	Category string
	Price    Price
}
