package product

import (
	"sales-portal-api/database"
	"sales-portal-api/models"
	"strings"
)

type Svc interface {
	CreateProduct(p *models.Product) *models.Product
	ListProducts(ids []*uint) []*models.Product
	GetMerchantProducts(merchantId *uint) []*models.Product
	UpdateProduct(p *models.Product) *models.Product
	DeleteProduct(id *uint) *models.Product
}

type svc struct {
	db *database.DbEngine
}

func NewSvc(db *database.DbEngine) Svc {
	return &svc{
		db: db,
	}
}

func (s *svc) CreateProduct(p *models.Product) *models.Product {
	return nil
}

func (s *svc) ListProducts(ids []*uint) []*models.Product {
	return nil
}

func (s *svc) UpdateProduct(p *models.Product) *models.Product {
	return nil
}

func (s *svc) DeleteProduct(id *uint) *models.Product {
	return nil
}

func (s *svc) GetMerchantProducts(merchantId *uint) []*models.Product {
	return nil
}

func CreateProduct(e *database.DbEngine, p *models.Product) *models.Product {
	err := e.DB.FirstOrCreate(&p, "merchant_id = $1 and name = $2", p.MerchantID, p.Name).Error

	if err != nil && !strings.Contains(err.Error(), "duplicate key") {
		e.Logger.Errorln(err)
		return nil
	}

	return p
}
