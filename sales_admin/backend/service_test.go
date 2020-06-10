package main

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

func TestSaveSales(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	svc := salesService{db: *sqlxDB}

	sale := Sale{
		CustomerName:    "bill the cat",
		Description:     "smoked salmon",
		Price:           10.00,
		Quantity:        100,
		MerchantName:    "wheel of fish",
		MerchantAddress: "some address",
	}

	mock.ExpectQuery("DELETE FROM sale").WillReturnRows()
	mock.ExpectBegin()
	mock.ExpectCommit()

	err = svc.SaveSales([]Sale{sale})
	if err != nil {
		t.Fatalf("error while testing SaveSale: %s\n", err)
	}
}
