package customer

import "github.com/labstack/echo"

type Handler interface {
	CreateCustomer(c echo.Context) error
	GetCustomer(c echo.Context) error
	ListCustomers(c echo.Context) error
	UpdateCustomer(c echo.Context) error
	DeleteCustomer(c echo.Context) error
}

type handler struct {
	svc Svc
}

func NewHandler(svc Svc) Handler {
	return &handler{
		svc: svc,
	}
}

func (h *handler) CreateCustomer(c echo.Context) error {
	return nil
}

func (h *handler) GetCustomer(c echo.Context) error {
	return nil
}

func (h *handler) ListCustomers(c echo.Context) error {
	return nil
}

func (h *handler) UpdateCustomer(c echo.Context) error {
	return nil
}

func (h *handler) DeleteCustomer(c echo.Context) error {
	return nil
}
