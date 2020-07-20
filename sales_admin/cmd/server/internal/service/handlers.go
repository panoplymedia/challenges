package service

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"

	"github.com/panoplymedia/sales_admin/internal/sales"
)

// listSales retrieves the list of sales from the database
func (s *Service) listSales(c echo.Context) error {
	// We should probably do something with pagination here to prevent huge payloads
	totalSales, err := s.sales.GetAll()
	if err != nil {
		c.Logger().Errorj(log.JSON{
			"error": err.Error(),
			"msg":   "Failed to retrieve the sales data from the database",
		})
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	err = c.Render(http.StatusOK, "sales-list",
		struct {
			Sales        []sales.Sale
			TotalRevenue float64
		}{
			Sales:        totalSales,
			TotalRevenue: sales.TotalRevenue(totalSales),
		},
	)

	return err
}

// UploadFileFormField is the name of the field that has the csv.
const UploadFileFormField = "file"

// uploadSales accepts a CSV of sales and inserts them into the database.
func (s *Service) uploadSales(c echo.Context) error {
	file, err := c.FormFile(UploadFileFormField)
	if err != nil {
		c.Logger().Warnj(log.JSON{
			"error": err.Error(),
			"msg":   fmt.Sprintf("Missing the '%s' field", UploadFileFormField),
		})

		return echo.NewHTTPError(http.StatusBadRequest, "the form did not have the file")
	}

	csv, err := file.Open()
	if err != nil {
		c.Logger().Warnj(log.JSON{
			"error": err.Error(),
			"msg":   fmt.Sprintf("Failed to open the uploaded file"),
		})

		return echo.NewHTTPError(http.StatusBadRequest, "could not open the file")
	}

	sales, err := sales.FromCSV(csv)
	if err != nil {
		c.Logger().Warnj(log.JSON{
			"error": err.Error(),
			"msg":   "Missing file field %s",
		})
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Failed to parse the CSV file: %s", err))
	}

	err = s.sales.Add(sales...)
	if err != nil {
		c.Logger().Errorj(log.JSON{
			"error": err.Error(),
			"msg":   "Failed to insert the parsed sales file into the database",
		})
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.Redirect(http.StatusSeeOther, "/sales")
}
