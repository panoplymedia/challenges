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

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestCalculateTotal(t *testing.T) {
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

	columns := []string{"1000.00"}
	mock.ExpectQuery(`SELECT SUM\(price \* quantity\) FROM sale`).WillReturnRows(sqlmock.NewRows(columns))
	svc.CalculateRevenue()

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
