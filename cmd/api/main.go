package main

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"log/slog"
	"mytheresa-promotions/internal/api"
	"mytheresa-promotions/internal/infra"
	"mytheresa-promotions/internal/product"
	"net/http"
	"os"
)

const DSN = "./var/products.db"
const PromotionsResultLimit = 5
const Port = "8080"
const DiscountCategory = "boots"
const DiscountCategoryPercentage = 30
const DiscountSku = "000003"
const DiscountSkuPercentage = 15

func main() {

	db, err := gorm.Open(sqlite.Open(DSN), &gorm.Config{TranslateError: true})
	if err != nil {
		os.Exit(1)
	}

	defer func() {
		sqlDB, _ := db.DB()
		if sqlDB != nil {
			sqlDB.Close()
		}
	}()

	handler := api.NewHandler(
		product.NewService(
			infra.NewGormRepository(db),
			product.NewCompositePromotion(
				product.NewCategoryPromotion(DiscountCategory, DiscountCategoryPercentage),
				product.NewSkuPromotion(DiscountSku, DiscountSkuPercentage),
			),
			PromotionsResultLimit,
		),
	)

	http.HandleFunc("GET /products/{sku}", handler.GetOne())
	http.HandleFunc("POST /products/{sku}", handler.Update())
	http.HandleFunc("DELETE /products/{sku}", handler.Delete())
	http.HandleFunc("POST /products", handler.Create())
	http.HandleFunc("GET /products", handler.Promotions())

	slog.Error(http.ListenAndServe(":"+Port, nil).Error())
}
