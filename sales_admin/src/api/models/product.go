package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	MerchantID  uint
	Name        string
	Description string
	ItemPrice   float32
}
