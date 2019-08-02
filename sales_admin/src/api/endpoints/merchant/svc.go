package merchant

import (
	"sales-portal-api/database"
	"sales-portal-api/models"
	"strings"
)

type Svc interface {
	CreateMerchant(p *models.Merchant) *models.Merchant
	GetMerchant(id *uint) *models.Merchant
	UpdateMerchant(p *models.Merchant) *models.Merchant
	DeleteMerchant(id *uint) *models.Merchant
}

type svc struct {
	db *database.DbEngine
}

func NewSvc(db *database.DbEngine) Svc {
	return &svc{
		db: db,
	}
}

func (s *svc) CreateMerchant(m *models.Merchant) *models.Merchant {
	return nil
}

func (s *svc) GetMerchant(id *uint) *models.Merchant {
	return nil
}

func (s *svc) UpdateMerchant(m *models.Merchant) *models.Merchant {
	return nil
}

func (s *svc) DeleteMerchant(id *uint) *models.Merchant {
	return nil
}

func CreateMerchant(e *database.DbEngine, m *models.Merchant) *models.Merchant {
	err := e.DB.FirstOrCreate(&m, "name = $1", m.Name).Error

	if err != nil && !strings.Contains(err.Error(), "duplicate key") {
		e.Logger.Errorln(err)
		return nil
	}

	return m
}
