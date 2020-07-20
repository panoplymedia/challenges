package mock_test

import (
	"errors"
	"reflect"
	"testing"

	. "github.com/panoplymedia/sales_admin/internal/mock"
	"github.com/panoplymedia/sales_admin/internal/sales"
)

func TestSalesServiceGetAll(t *testing.T) {
	expectedSales := []sales.Sale{
		{
			Customer: "Bob",
			Item:     sales.Item{Description: "Pants", Price: 100.00},
			Quantity: 3,
			Merchant: sales.Merchant{Name: "Pants R US", Address: "Pants Road"},
		},
		{
			Customer: "Alice",
			Item:     sales.Item{Description: "Shirt", Price: 100.00},
			Quantity: 3,
			Merchant: sales.Merchant{Name: "Shirt R US", Address: "Shirt Road"},
		},
	}
	expectedError := errors.New("error")

	service := SalesService{
		GetAllFn: func() ([]sales.Sale, error) {
			// Returning both sales and an error so that we can check that both
			// are passed up the stack.
			return expectedSales, expectedError

		},
	}

	sales, err := service.GetAll()

	if err != expectedError {
		t.Errorf("Unexpected error returned: expected %v, but got %v", expectedError, err)
	}

	if !reflect.DeepEqual(expectedSales, sales) {
		t.Errorf("Sales did not match: expected %v, but got %v", expectedSales, sales)
	}

	if !service.GetAllInvoked {
		t.Errorf("Expected GetAllInvoked to set to true, but was false")
	}
}

func TestSalesServiceAdd(t *testing.T) {
	input := []sales.Sale{
		{
			Customer: "Bob",
			Item:     sales.Item{Description: "Pants", Price: 100.00},
			Quantity: 3,
			Merchant: sales.Merchant{Name: "Pants R US", Address: "Pants Road"},
		},
		{
			Customer: "Alice",
			Item:     sales.Item{Description: "Shirt", Price: 100.00},
			Quantity: 3,
			Merchant: sales.Merchant{Name: "Shirt R US", Address: "Shirt Road"},
		},
	}
	expectedError := errors.New("error")

	service := SalesService{
		AddFn: func(sales ...sales.Sale) error {
			if !reflect.DeepEqual(input, sales) {
				t.Errorf("Sales did not match: expected %v, but got %v", input, sales)
			}

			return expectedError
		},
	}

	err := service.Add(input...)

	if err != expectedError {
		t.Errorf("Unexpected error returned: expected %v, but got %v", expectedError, err)
	}

	if !service.AddInvoked {
		t.Errorf("Expected GetAllInvoked to set to true, but was false")
	}
}
