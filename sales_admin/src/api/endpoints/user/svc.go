package user

import (
	"github.com/labstack/echo"
	"sales-portal-api/database"
)

type Svc interface {
	CreateUser(c echo.Context) error
	GetUser(c echo.Context) error
	UpdateUser(c echo.Context) error
	DeleteUser(c echo.Context) error
}

type svc struct {
	engine *database.DbEngine
}

func NewSvc(engine *database.DbEngine) Svc {
	return &svc{
		engine: engine,
	}
}

func (svc *svc) CreateUser(c echo.Context) error {

}

func (svc *svc) GetUser(c echo.Context) error {

}

func (svc *svc) UpdateUser(c echo.Context) error {

}

func (svc *svc) DeleteUser(c echo.Context) error {

}
