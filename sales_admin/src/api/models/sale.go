package models

import "github.com/jinzhu/gorm"

type Sale struct {
	gorm.Model
	CustomerID uint
	ProductID  uint
	MerchantID uint
	Quantity   uint
	TotalPrice float32
}

type SalesSummary struct {
	TotalSalesRevenue *float32 `json:"totalSalesRevenue"`
}
