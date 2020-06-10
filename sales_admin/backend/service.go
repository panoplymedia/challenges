package main

import (
	"encoding/csv"
	"fmt"
	"mime/multipart"
	"os"
	"strconv"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var schema = `
CREATE TABLE IF NOT EXISTS sale (
	customer_name TEXT,
	description TEXT,
	price FLOAT,
	quantity INTEGER,
	merchant_name TEXT,
	merchant_address TEXT
)`

func NewSalesService() (SalesService, error) {
	dbHost := os.Getenv("PG_HOST")
	dbName := os.Getenv("PG_DB")
	dbPassword := os.Getenv("PG_PASS")
	dbPort := os.Getenv("PG_PORT")
	dbUser := os.Getenv("PG_USER")

	dbInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbName, dbPassword, dbUser)

	db, err := sqlx.Connect("postgres", dbInfo)
	if err != nil {
		return nil, err
	}

	return &salesService{db: *db}, nil
}

func (s *salesService) CreateSalesTable() error {
	_, err := s.db.Exec(schema)
	if err != nil {
		return err
	}

	return nil
}

func (s *salesService) SaveSales(sales []Sale) error {
	_, err := s.db.Query("DELETE FROM sale")
	if err != nil {
		return err
	}

	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	for _, sale := range sales {
		tx.Exec("INSERT INTO sale (customer_name, description, price, quantity, merchant_name, merchant_address) VALUES ($1, $2, $3, $4, $5, $6)", sale.CustomerName, sale.Description, sale.Price, sale.Quantity, sale.MerchantName, sale.MerchantAddress)
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (s *salesService) CalculateRevenue() (float64, error) {
	res := s.db.QueryRowx("SELECT SUM(price * quantity) FROM sale")
	var sum float64

	err := res.Scan(&sum)
	if err != nil {
		return 0.0, err
	}

	return sum, nil
}

func processCSV(file multipart.File) ([]Sale, error) {
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 0

	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	sales := []Sale{}

	for _, row := range data[1:] {
		price, err := strconv.ParseFloat(row[2], 64)
		if err != nil {
			return nil, err
		}

		quantity, err := strconv.Atoi(row[3])
		if err != nil {
			return nil, err
		}

		sale := Sale{
			CustomerName:    row[0],
			Description:     row[1],
			Price:           price,
			Quantity:        quantity,
			MerchantName:    row[4],
			MerchantAddress: row[5],
		}

		sales = append(sales, sale)

	}

	return sales, nil
}
