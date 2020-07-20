package mock

import "github.com/panoplymedia/sales_admin/internal/sales"

// SalesService represents a mock implementation of sale.Service.
type SalesService struct {
	GetAllFn      func() ([]sales.Sale, error)
	GetAllInvoked bool

	AddFn      func(...sales.Sale) error
	AddInvoked bool
}

// GetAll invokes the mock implementation and marks the function as invoked.
func (s *SalesService) GetAll() ([]sales.Sale, error) {
	s.GetAllInvoked = true
	return s.GetAllFn()
}

// Add invokes the mock implementation and marks the function as invoked.
func (s *SalesService) Add(sales ...sales.Sale) error {
	s.AddInvoked = true
	return s.AddFn(sales...)
}
