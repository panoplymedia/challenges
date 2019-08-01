package user

import (
	"github.com/labstack/echo"
)

type Handler interface {
	CreateUser(c echo.Context) error
	GetUser(c echo.Context) error
	UpdateUser(c echo.Context) error
	DeleteUser(c echo.Context) error
}

type handler struct {
	svc Svc
}

func NewHandler(svc Svc) Handler {
	return &handler{
		svc: svc,
	}
}

func (h *handler) CreateUser(c echo.Context) error {

}

func (h *handler) GetUser(c echo.Context) error {

}

func (h *handler) UpdateUser(c echo.Context) error {

}

func (h *handler) DeleteUser(c echo.Context) error {

}
