package product

import (
	"errors"
	"mytheresa-promotions/internal/infra"
)

var NotFoundError = errors.New("product not found")
var DuplicateSkuError = errors.New("product duplicate sku")
var InternalError = errors.New("internal error")

type Service interface {
	FindBySku(sku string) (*Product, error)
	BulkCreate(products []*Product) error
	Update(p *Product) (*Product, error)
	Delete(sku string) error
	GetWithPromotions(category string, priceLessThan int) ([]*Product, error)
}

type PromotionService struct {
	repository           infra.Repository
	promotion            Promotion
	promotionResultLimit int
}

func NewService(repository infra.Repository, promotion Promotion, promotionResultLimit int) Service {
	return &PromotionService{
		repository:           repository,
		promotion:            promotion,
		promotionResultLimit: promotionResultLimit,
	}
}

func (s *PromotionService) FindBySku(sku string) (*Product, error) {
	repoProduct, err := s.repository.GetBySku(sku)
	if err != nil {
		if errors.Is(err, infra.ProductNotFoundError) {
			return nil, NotFoundError
		}
		// todo log error
		return nil, InternalError
	}

	return convertFromInfraToModel(repoProduct), nil
}

func (s *PromotionService) BulkCreate(products []*Product) error {

	infraProducts := convertSliceFromModelToInfra(products)

	if _, err := s.repository.Create(infraProducts); err != nil {
		if errors.Is(err, infra.ProductDuplicateSkuError) {
			return DuplicateSkuError
		}

		return InternalError
	}

	return nil
}

func (s *PromotionService) Update(p *Product) (*Product, error) {
	product, err := s.repository.Update(convertModelToInfra(p))
	if err != nil {
		if errors.Is(err, infra.ProductNotFoundError) {
			return nil, NotFoundError
		}
		return nil, InternalError
	}

	return convertFromInfraToModel(product), nil
}

func (s *PromotionService) Delete(sku string) error {
	err := s.repository.Delete(sku)

	if errors.Is(err, infra.ProductNotFoundError) {
		return NotFoundError
	}

	if err != nil {
		return InternalError
	}

	return nil
}

func (s *PromotionService) GetWithPromotions(category string, priceLessThan int) ([]*Product, error) {
	infraProducts, err := s.repository.Search(category, priceLessThan, s.promotionResultLimit)
	if err != nil {
		return nil, InternalError
	}

	products := convertSliceFromInfraToModel(infraProducts)

	if s.promotion != nil {
		for i, p := range products {
			products[i] = s.promotion.Apply(p)
		}
	}

	return products, nil
}
