package main

import "github.com/jmoiron/sqlx"

type SalesService interface {
	SaveSale(sale Sale) error
	CreateSalesTable() error
}

type Sale struct {
	CustomerName    string
	Description     string
	Price           string
	Quantity        int
	MerchantName    string
	MerchantAddress string
}

type salesService struct {
	db sqlx.DB
}
