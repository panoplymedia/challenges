package sales

import (
	"encoding/csv"
	"io"
	"strconv"

	"github.com/pkg/errors"
)

// Merchant is an entity that sells items to customers.
type Merchant struct {
	Name string
	// Address could be broken out further, but I will leave this as a string
	// for simplicity.
	Address string
}

// Item is something that a merchant sells to customers.
type Item struct {
	Description string
	Price       float64
}

// Sale is sale from a merchant to a customer.
type Sale struct {
	// Customer could be its own struct, but I will leave this as a string for
	// simplicity.
	Customer string
	Merchant Merchant
	Item     Item
	Quantity int64
}

// FromCSV converts a CSV file to a list of sales. The expected format of the
// CSV files is:
//    Customer Name,Item Description,Item Price,Quantity,Merchant Name,Merchant Address
// The first line will be dropped because we expect that to be the headings
func FromCSV(r io.Reader) ([]Sale, error) {
	if r == nil {
		return nil, errors.New("reader was nil")
	}

	parser := csv.NewReader(r)
	parser.FieldsPerRecord = 6
	parser.TrimLeadingSpace = true
	parser.ReuseRecord = true

	// Make sure the header is correct
	header, err := parser.Read()
	if err != nil {
		return nil, errors.Wrap(err, "failed to read the first line of the csv")
	}

	if err := validateHeader(header); err != nil {
		return nil, errors.Wrap(err, "invalid header: expected Customer Name,Item Description,Item Price,Quantity,Merchant Name,Merchant Address")
	}

	var sales []Sale
	for row, err := parser.Read(); err == nil; row, err = parser.Read() {
		price, err := strconv.ParseFloat(row[2], 64)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to convert '%s' into a price", row[2])
		}

		quantity, err := strconv.ParseInt(row[3], 10, 64)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to convert '%s' into a valid quantity", row[3])
		}

		sales = append(sales, Sale{
			Customer: row[0],
			Merchant: Merchant{Name: row[4], Address: row[5]},
			Item:     Item{Description: row[1], Price: price},
			Quantity: quantity,
		})
	}

	return sales, nil
}

func validateHeader(row []string) error {
	switch {
	case row[0] != "Customer Name":
		return errors.Errorf("field 0 invalid: expected Customer Name, but got %v", row[0])
	case row[1] != "Item Description":
		return errors.Errorf("field 1 invalid: expected Item Description, but got %v", row[1])
	case row[2] != "Item Price":
		return errors.Errorf("field 2 invalid: expected Item Price, but got %v", row[2])
	case row[3] != "Quantity":
		return errors.Errorf("field 3 invalid: expected Quantity, but got %v", row[3])
	case row[4] != "Merchant Name":
		return errors.Errorf("field 4 invalid: expected Merchant Name, but got %v", row[4])
	case row[5] != "Merchant Address":
		return errors.Errorf("field 5 invalid: expected Quantity, but got %v", row[5])
	}

	return nil
}

// SalesService handles retrieving and adding sales
type Service interface {
	GetAll() ([]Sale, error)
	Add(...Sale) error
}

// TotalRevenue returns the total revenue for the given sales.
func TotalRevenue(sales []Sale) float64 {
	var total float64
	for _, sale := range sales {
		total += float64(sale.Quantity) * sale.Item.Price
	}

	return total
}
