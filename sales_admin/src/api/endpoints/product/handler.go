package product

import "github.com/labstack/echo"

type Handler interface {
	CreateProduct(c echo.Context) error
	ListProducts(c echo.Context) error
	GetMerchantProducts(c echo.Context) error
	UpdateProduct(c echo.Context) error
	DeleteProduct(c echo.Context) error
}

type handler struct {
	svc Svc
}

func NewHandler(svc Svc) Handler {
	return &handler{
		svc: svc,
	}
}

func (h *handler) CreateProduct(c echo.Context) error {
	return nil
}

func (h *handler) ListProducts(c echo.Context) error {
	return nil
}

func (h *handler) GetMerchantProducts(c echo.Context) error {
	return nil
}

func (h *handler) UpdateProduct(c echo.Context) error {
	return nil
}

func (h *handler) DeleteProduct(c echo.Context) error {
	return nil
}
