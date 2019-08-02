package customer

import (
	"sales-portal-api/database"
	"sales-portal-api/models"
	"strings"
)

type Svc interface {
	CreateCustomer(Customer *models.Customer) *models.Customer
	GetCustomer(id *uint) *models.Customer
	ListCustomers() []*models.Customer
	UpdateCustomer(Customer *models.Customer) *models.Customer
	DeleteCustomer(id *uint) *models.Customer
}

type svc struct {
	engine *database.DbEngine
}

func NewSvc(engine *database.DbEngine) Svc {
	return &svc{
		engine: engine,
	}
}

func (s *svc) CreateCustomer(Customer *models.Customer) *models.Customer {
	return nil
}

func (s *svc) GetCustomer(id *uint) *models.Customer {
	return nil
}

func (s *svc) UpdateCustomer(Customer *models.Customer) *models.Customer {
	return nil
}

func (s *svc) DeleteCustomer(id *uint) *models.Customer {
	return nil
}

func (s *svc) ListCustomers() []*models.Customer {
	return nil
}

func CreateCustomer(e *database.DbEngine, c *models.Customer) *models.Customer {
	err := e.DB.FirstOrCreate(&c, "email = $1", c.Email).Error

	if err != nil && !strings.Contains(err.Error(), "duplicate key"){
		e.Logger.Errorln(err)
		return nil
	}

	return c
}
