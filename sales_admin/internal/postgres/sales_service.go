package postgres

import (
	"database/sql"

	"github.com/panoplymedia/sales_admin/internal/sales"
	"github.com/pkg/errors"
)

// SalesService implement the sales.SalesService interface by using postgres as
// the datastore.
type SalesService struct {
	*sql.DB
}

func (s *SalesService) GetAll() ([]sales.Sale, error) {
	rows, err := s.Query(`SELECT
	                          sales.customer_name,
	                          items.description,
	                          items.price,
	                          sales.quantity as quantity,
	                          merchants.name,
	                          merchants.address
	                      FROM
	                          sales
	                      JOIN items     ON (sales.item_id = items.id)
	                      JOIN merchants ON (sales.merchant_id = merchants.id)`)

	if err != nil {
		return nil, errors.Wrap(err, "failed to retrieve the sales data")
	}
	defer rows.Close()

	var allSales []sales.Sale
	for rows.Next() {
		var sale sales.Sale
		err := rows.Scan(
			&sale.Customer,
			&sale.Item.Description,
			&sale.Item.Price,
			&sale.Quantity,
			&sale.Merchant.Name,
			&sale.Merchant.Address,
		)

		if err != nil {
			return nil, errors.Wrap(err, "failed to scan a row of the sales data")
		}

		allSales = append(allSales, sale)
	}

	if rows.Err() != nil {
		return nil, errors.Wrap(err, "failed to access rows of the sales data")
	}

	return allSales, nil
}

// Add inserts the provided sales into the database.
func (s *SalesService) Add(sales ...sales.Sale) (err error) {
	tx, err := s.Begin()
	if err != nil {
		return errors.Wrap(err, "failed to start the transaction")
	}

	defer func() {
		// We got an error somewhere in the function, so we need to rollback the
		// transaction.
		if err != nil {
			// TODO(wh): Log Error
			tx.Rollback()
		}
	}()

	insertMerchant, err := tx.Prepare(`INSERT INTO merchants (name, address) VALUES ($1, $2)
	                                   ON CONFLICT DO NOTHING`)
	if err != nil {
		return errors.Wrap(err, "failed to prepare insert merchants statement")
	}

	insertItem, err := tx.Prepare(`INSERT INTO items (description, price) VALUES ($1, $2)
	                               ON CONFLICT DO NOTHING`)
	if err != nil {
		return errors.Wrap(err, "failed to prepare insert items statement")
	}

	insertSale, err := tx.Prepare(`INSERT INTO sales (
	                                           customer_name, 
	                                           item_id, 
	                                           quantity, 
	                                           merchant_id) 
	                                    VALUES (
	                                            $1,
	                                            (SELECT id FROM items WHERE items.description = $2),
	                                            $3,
	                                            (SELECT id FROM merchants WHERE merchants.name = $4 AND merchants.address = $5)
	                                           )`)
	if err != nil {
		return errors.Wrap(err, "failed to prepare insert sales statement")
	}

	for _, s := range sales {
		_, err := insertMerchant.Exec(s.Merchant.Name, s.Merchant.Address)
		if err != nil {
			return errors.Wrapf(err, "failed to insert merchant %#v", s.Merchant)
		}

		_, err = insertItem.Exec(s.Item.Description, s.Item.Price)
		if err != nil {
			return errors.Wrapf(err, "failed to insert item %#v", s.Item)
		}

		_, err = insertSale.Exec(s.Customer, s.Item.Description,
			s.Quantity, s.Merchant.Name, s.Merchant.Address)
		if err != nil {
			return errors.Wrapf(err, "failed to insert sale %#v", s)
		}
	}

	if err := tx.Commit(); err != nil {
		return errors.Wrap(err, "failed to commit the sales data to the database")
	}

	return nil
}
