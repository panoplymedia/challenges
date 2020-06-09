package main

import "github.com/jmoiron/sqlx"

type SalesService interface {
	SaveSales(sale []Sale) error
	CreateSalesTable() error
}

type Sale struct {
	CustomerName    string  `db:"customer_name"`
	Description     string  `db:"description"`
	Price           float64 `db:"price"`
	Quantity        int     `db:"quantity"`
	MerchantName    string  `db:"merchant_name"`
	MerchantAddress string  `db:"merchant_address"`
}

type salesService struct {
	db sqlx.DB
}
