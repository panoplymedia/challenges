package sales_test

import (
	"io"
	"reflect"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	. "github.com/panoplymedia/sales_admin/internal/sales"
)

func TestTotalRevenue(t *testing.T) {
	tests := map[string]struct {
		sales           []Sale
		expectedRevenue float64
	}{
		"should return the sum of the sales": {
			sales: []Sale{
				{
					Customer: "Bob",
					Item:     Item{Description: "Pants", Price: 100.00},
					Quantity: 3,
					Merchant: Merchant{Name: "Pants R US", Address: "Pants Road"},
				},
				{
					Customer: "Alice",
					Item:     Item{Description: "Shirt", Price: 15.20},
					Quantity: 2,
					Merchant: Merchant{Name: "Shirt R US", Address: "Shirt Road"},
				},
			},
			expectedRevenue: 330.40,
		},

		"should return zero for empty args": {
			sales:           nil,
			expectedRevenue: 0,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			revenue := TotalRevenue(test.sales)
			if revenue != test.expectedRevenue {
				t.Errorf("The revenue did not match: expected %f, but got %f", test.expectedRevenue, revenue)
			}
		})
	}
}

func TestFromCSV(t *testing.T) {
	tests := map[string]struct {
		reader        io.Reader
		expectedSales []Sale
		errExpected   bool
	}{
		"should return a list of sales": {
			reader: strings.NewReader(`Customer Name,Item Description,Item Price,Quantity,Merchant Name,Merchant Address
			                           Jack Burton,Premium Cowboy Boots,149.99,1,Carpenter Outfitters,99 Factory Drive
			                           Ellen Ripley,Tank Top Undershirt,9.50,2,Hero Outlet,123 Main Street
			                           Lisbeth Salander,Black Hoodie,49.99,4,Stockholm Supplies,34 Other Avenue`),
			expectedSales: []Sale{
				{
					Customer: "Jack Burton",
					Item:     Item{Description: "Premium Cowboy Boots", Price: 149.99},
					Quantity: 1,
					Merchant: Merchant{Name: "Carpenter Outfitters", Address: "99 Factory Drive"},
				},
				{
					Customer: "Ellen Ripley",
					Item:     Item{Description: "Tank Top Undershirt", Price: 9.50},
					Quantity: 2,
					Merchant: Merchant{Name: "Hero Outlet", Address: "123 Main Street"},
				},
				{
					Customer: "Lisbeth Salander",
					Item:     Item{Description: "Black Hoodie", Price: 49.99},
					Quantity: 4,
					Merchant: Merchant{Name: "Stockholm Supplies", Address: "34 Other Avenue"},
				},
			},
			errExpected: false,
		},

		"should return an error missing columns": {
			reader: strings.NewReader(`Customer Name,Item Price,Quantity,Merchant Name,Merchant Address
			                           Jack Burton,149.99,1,Carpenter Outfitters,99 Factory Drive
			                           Ellen Ripley,9.50,2,Hero Outlet,123 Main Street
			                           Lisbeth Salander,49.99,4,Stockholm Supplies,34 Other Avenue`),
			expectedSales: nil,
			errExpected:   true,
		},

		"should return an error if reader isn't csv": {
			reader:        strings.NewReader("garbage in errors out"),
			expectedSales: nil,
			errExpected:   true,
		},
		"should return an error if the reader is nil": {
			reader:        nil,
			expectedSales: nil,
			errExpected:   true,
		},
		"should return empty list if the header is correct, but the body is empty": {
			reader:        strings.NewReader(`Customer Name,Item Description,Item Price,Quantity,Merchant Name,Merchant Address`),
			expectedSales: nil,
			errExpected:   false,
		},
		"should return an error is the header is missing": {
			reader: strings.NewReader(`Jack Burton,Premium Cowboy Boots,149.99,1,Carpenter Outfitters,99 Factory Drive
			                           Ellen Ripley,Tank Top Undershirt,9.50,2,Hero Outlet,123 Main Street
			                           Lisbeth Salander,Black Hoodie,49.99,4,Stockholm Supplies,34 Other Avenue`),
			expectedSales: nil,
			errExpected:   true,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			sales, err := FromCSV(test.reader)

			if !reflect.DeepEqual(test.expectedSales, sales) {
				t.Log(cmp.Diff(test.expectedSales, sales))
				t.Errorf("Sales did not match: expected %#v, got %#v", test.expectedSales, sales)
			}

			if (err != nil) != test.errExpected {
				t.Log(err)
				t.Errorf("Error expected was %v, but got %v", test.errExpected, err != nil)
			}
		})
	}
}
