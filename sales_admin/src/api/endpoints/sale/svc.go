package sale

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"github.com/icrowley/fake"
	"io"
	"sales-portal-api/database"
	"sales-portal-api/endpoints/customer"
	"sales-portal-api/endpoints/merchant"
	"sales-portal-api/endpoints/product"
	"sales-portal-api/models"
	"strconv"
	"strings"
)

type Svc interface {
	CreateSale(s *models.Sale) *models.Sale
	UpdateSale(s *models.Sale) *models.Sale
	GetSale(id *uint) *models.Sale
	ListSales(userId *uint) []*models.Sale
	DeleteSale(id *uint) *models.Sale
	GetSalesSummary() *models.SalesSummary
	UploadSalesCsv(salesData io.Reader) error
}

type svc struct {
	db *database.DbEngine
}

func NewSvc(db *database.DbEngine) Svc {
	return &svc{
		db: db,
	}
}

func (s *svc) CreateSale(sale *models.Sale) *models.Sale {
	return nil
}

func (s *svc) GetSale(id *uint) *models.Sale {
	return nil
}

func (s *svc) UpdateSale(sale *models.Sale) *models.Sale {
	return nil
}

func (s *svc) ListSales(userId *uint) []*models.Sale {
	return nil
}

func (s *svc) DeleteSale(id *uint) *models.Sale {
	return nil
}

func (s *svc) GetSalesSummary() *models.SalesSummary {
	row := s.db.DB.Raw(`select sum(total_price) from sale`).Row()

	var _revenue sql.NullFloat64
	if err := row.Scan(&_revenue); err != nil {
		s.db.Logger.Errorln(err)
		return nil
	}

	revenue := float32(_revenue.Float64)

	return &models.SalesSummary{TotalSalesRevenue: &revenue}
}

func (s *svc) UploadSalesCsv(salesData io.Reader) error {
	return s.LoadCsvRecords(csv.NewReader(salesData))
}

func (s *svc) LoadCsvRecords(reader *csv.Reader) error {
	data, err := reader.ReadAll()
	if err != nil {
		return err
	}

	// Clear previous data
	if err := s.clearData(); err != nil {
		return err
	}

	// Ignore header row
	for _, record := range data[1:] {
		if err := s.processRecord(record); err != nil {
			return err
		}
	}
	return nil
}

func (s *svc) clearData() error {
	tables := []string{"customer", "merchant", "product", "sale"}

	for _, table := range tables {
		if err := s.db.DB.Exec(fmt.Sprintf(`delete from %s`, table)).Error; err != nil {
			return err
		}
	}

	return nil
}

func (s *svc) processRecord(r []string) error {
	if len(r) == 6 {
		customerName := r[0]
		itemDesc := r[1]
		totalCost, _ := strconv.ParseFloat(r[2], 32)
		quantity, _ := strconv.Atoi(r[3])
		merchantName := r[4]
		merchantAddress := r[5]

		Customer := &models.Customer{
			FullName: strings.ToLower(customerName),
			Email:    fmt.Sprintf("%s@email.com", strings.TrimSpace(strings.ToLower(customerName))),
		}

		Customer = customer.CreateCustomer(s.db, Customer)

		Merchant := &models.Merchant{
			Name:    strings.ToLower(merchantName),
			Address: strings.ToLower(merchantAddress),
		}

		Merchant = merchant.CreateMerchant(s.db, Merchant)

		Product := &models.Product{
			Name:        fake.ProductName(),
			Description: itemDesc,
			MerchantID:  Merchant.ID,
			ItemPrice:   float32(totalCost) / float32(quantity),
		}

		Product = product.CreateProduct(s.db, Product)

		Sale := &models.Sale{
			CustomerID: Customer.ID,
			MerchantID: Merchant.ID,
			ProductID:  Product.ID,
			Quantity:   uint(quantity),
			TotalPrice: float32(totalCost),
		}

		Sale = CreateSale(s.db, Sale)

	} else {
		return fmt.Errorf("number of fields must equal 6")
	}
	return nil
}

func CreateSale(e *database.DbEngine, s *models.Sale) *models.Sale {
	err := e.DB.Create(&s).Error

	if err != nil && !strings.Contains(err.Error(), "duplicate key") {
		e.Logger.Errorln(err)
		return nil
	}

	return s
}
