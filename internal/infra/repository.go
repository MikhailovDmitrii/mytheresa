package infra

import (
	"errors"
	"gorm.io/gorm"
)

const BatchSize = 1000

var ProductNotFoundError error = errors.New("product is not found")
var ProductDuplicateSkuError error = errors.New("duplicate product sku")

type Repository interface {
	Create(products []*Product) ([]*Product, error)
	Update(p *Product) (*Product, error)
	GetBySku(sku string) (*Product, error)
	Search(category string, priceLessThan int, limit int) ([]*Product, error)
	Delete(sku string) error
}

type GormRepository struct {
	db *gorm.DB
}

func NewGormRepository(db *gorm.DB) Repository {
	db.CreateBatchSize = BatchSize
	return &GormRepository{
		db: db,
	}
}

func (r *GormRepository) Create(products []*Product) ([]*Product, error) {
	result := r.db.Create(products)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, ProductDuplicateSkuError
		}

		return nil, result.Error
	}

	return products, nil
}

func (r *GormRepository) Update(p *Product) (*Product, error) {
	saveResult := r.db.Model(&p).Updates(p)
	if saveResult.Error != nil {
		return nil, saveResult.Error
	}

	if saveResult.RowsAffected == 0 {
		return nil, ProductNotFoundError
	}

	return p, nil
}

func (r *GormRepository) GetBySku(sku string) (*Product, error) {
	var p *Product
	result := r.db.Take(&p, sku)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ProductNotFoundError
		}

		return nil, result.Error
	}

	return p, nil
}

func (r *GormRepository) Search(category string, priceLessThan int, limit int) ([]*Product, error) {
	var p []*Product
	query := r.db.Limit(limit)
	if category != "" {
		query = query.Where("category = ?", category)
	}
	if priceLessThan > 0 {
		query = query.Where("price <= ?", priceLessThan)
	}

	result := query.Find(&p)

	if result.Error != nil {
		return nil, result.Error
	}

	return p, nil
}

func (r *GormRepository) Delete(sku string) error {
	result := r.db.Delete(Product{Sku: sku})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ProductNotFoundError
	}

	return nil
}
