package models

import "github.com/jinzhu/gorm"

type Merchant struct {
	gorm.Model
	Name string
	Address string
	Products []*Product
}
