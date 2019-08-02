package merchant

import "github.com/labstack/echo"

type Handler interface {
	CreateMerchant(c echo.Context) error
	GetMerchant(c echo.Context) error
	UpdateMerchant(c echo.Context) error
	DeleteMerchant(c echo.Context) error
}

type handler struct {
	svc Svc
}

func NewHandler(svc Svc) Handler {
	return &handler{
		svc: svc,
	}
}

func (h *handler) CreateMerchant(c echo.Context) error {
	return nil
}

func (h *handler) GetMerchant(c echo.Context) error {
	return nil
}

func (h *handler) UpdateMerchant(c echo.Context) error {
	return nil
}

func (h *handler) DeleteMerchant(c echo.Context) error {
	return nil
}
