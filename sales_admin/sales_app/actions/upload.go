package actions

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/x/responder"
	"github.com/henry-jackson/challenges/sales_admin/sales_app/models"
	"github.com/pkg/errors"
)

// UploadHandler handles csv uploads to the /upload endpoint
func UploadHandler(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	f, err := c.File("someFile")
	if err != nil {
		return errors.WithStack(err)
	}

	if !f.Valid() {
		return errors.WithStack(errors.New("invalid file upload"))
	}

	dir := filepath.Join(".", "uploads")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return errors.WithStack(err)
	}
	osFile, err := os.Create(filepath.Join(dir, f.String()))
	if err != nil {
		return errors.WithStack(err)
	}
	defer osFile.Close()
	buf := bytes.Buffer{}
	w := io.MultiWriter(&buf, osFile, os.Stdout)
	_, err = io.Copy(w, f)
	if err != nil {
		return errors.WithStack(err)
	}

	report, err := csv.NewReader(&buf).ReadAll()
	if err != nil {
		return errors.WithStack(err)
	}

	if len(report) < 1 {
		return errors.WithStack(errors.New("invalid file upload"))
	}

	columnIdxMap := map[string]int{}
	for idx, header := range report[0] {
		columnIdxMap[header] = idx
	}
	type Record struct {
		CustomerName    string
		ItemDesc        string
		ItemPrice       float64
		Quantity        int
		MerchantName    string
		MerchantAddress string
	}
	records := []Record{}
	for _, row := range report[1:] {
		price, err := strconv.ParseFloat(row[columnIdxMap["Item Price"]], 64)
		if err != nil {
			continue
		}
		qty, err := strconv.Atoi(row[columnIdxMap["Quantity"]])
		if err != nil {
			continue
		}
		rec := Record{
			CustomerName:    row[columnIdxMap["Customer Name"]],
			ItemDesc:        row[columnIdxMap["Item Description"]],
			ItemPrice:       price,
			Quantity:        qty,
			MerchantName:    row[columnIdxMap["Merchant Name"]],
			MerchantAddress: row[columnIdxMap["Merchant Address"]],
		}
		records = append(records, rec)
	}

	for _, rec := range records {
		customer := &models.Customer{
			Name: rec.CustomerName,
		}
		tx.Create(customer)

		merchant := &models.Merchant{
			Name:    rec.MerchantName,
			Address: rec.MerchantAddress,
		}
		tx.Create(merchant)

		product := &models.Product{
			Price:       rec.ItemPrice,
			MerchantID:  merchant.ID,
			Description: rec.ItemDesc,
		}
		tx.Eager().Create(product)

		order := &models.Order{
			CustomerID: customer.ID,
			ProductID:  product.ID,
			MerchantID: merchant.ID,
			Quantity:   rec.Quantity,
		}
		tx.Eager().Create(order)
	}

	orders := []models.Order{}

	// Retrieve all Orders from the DB
	if err := tx.Eager().All(&orders); err != nil {
		return err
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		c.Set("orders", orders)
		return c.Render(http.StatusOK, r.HTML("/orders/index.plush.html"))
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(200, r.JSON(orders))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(200, r.XML(orders))
	}).Respond(c)
}
