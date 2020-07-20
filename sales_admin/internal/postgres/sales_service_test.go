package postgres_test

import (
	"reflect"
	"testing"

	"github.com/panoplymedia/sales_admin/internal/postgres"
	"github.com/panoplymedia/sales_admin/internal/sales"
)

func TestSalesService(t *testing.T) {
	db, err := postgres.Connect(connectionArgs)
	if err != nil {
		t.Fatalf("Unexpected error when attempting to connect to the database: %v", err)
	}
	service := postgres.SalesService{DB: db}
	expected := []sales.Sale{
		{
			Customer: "Tim",
			Item:     sales.Item{Description: "Shoes", Price: 10.00},
			Quantity: 2,
			Merchant: sales.Merchant{Name: "Shoes R US", Address: "Shoes Road"},
		},
		{
			Customer: "Betty",
			Item:     sales.Item{Description: "Belt", Price: 50.00},
			Quantity: 1,
			Merchant: sales.Merchant{Name: "Belts R US", Address: "Belt Road"},
		},
	}
	err = service.Add(expected...)
	if err != nil {
		t.Fatalf("Unexpected error returned when attempting add to the database: %v", err)
	}
	defer func() {
		// TODO(wh): This is pretty heavy handed, because if anything else is
		// accessing the database during the test it will be clobbered.
		db.Exec("DELETE FROM sales")
		db.Exec("DELETE FROM merchants")
		db.Exec("DELETE FROM items")
	}()

	sales, err := service.GetAll()
	if err != nil {
		t.Errorf("Unexpected error returned when attempting to call GetAll: %v", err)
	}

	if !reflect.DeepEqual(expected, sales) {
		t.Errorf("The returned sales did not match: expected %#v, but got %#v", expected, sales)
	}

}
