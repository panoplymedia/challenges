package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var schema = `
CREATE TABLE sale (
	customer_name TEXT,
	description TEXT,
	price TEXT,
	quantity INTEGER,
	merchant_name TEXT,
	merchant_address TEXT
)`

func NewSalesService() (SalesService, error) {
	dbInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		"ruby.db.elephantsql.com", "5432", "gcyotvfr", "hNSdE5i665siQjFY9KK-R-gTRU4FAG7k", "gcyotvfr")

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

func (s *salesService) SaveSale(sale Sale) error {
	return nil
}
