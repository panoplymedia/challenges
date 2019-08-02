package sale

import (
	"github.com/labstack/echo"
	"net/http"
)

type Handler interface {
	CreateSale(c echo.Context) error
	GetSale(c echo.Context) error
	UpdateSale(c echo.Context) error
	ListSales(c echo.Context) error
	DeleteSale(c echo.Context) error
	GetSalesSummary(c echo.Context) error
	UploadSalesCsv(c echo.Context) error
}

type handler struct {
	svc Svc
}

func NewHandler(svc Svc) Handler {
	return &handler{svc: svc}
}

func (h *handler) CreateSale(c echo.Context) error {
	return nil
}
func (h *handler) GetSale(c echo.Context) error {
	return nil
}
func (h *handler) UpdateSale(c echo.Context) error {
	return nil
}
func (h *handler) ListSales(c echo.Context) error {
	return nil
}
func (h *handler) DeleteSale(c echo.Context) error {
	return nil
}
func (h *handler) GetSalesSummary(c echo.Context) error {
	summary := h.svc.GetSalesSummary()
	if summary == nil {
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusOK, summary)
}

func (h *handler) UploadSalesCsv(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	return h.svc.UploadSalesCsv(src)
}
