package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Username     string
	PasswordHash string
	Email        *string
	Blocked      bool
	DateBlocked  *time.Time
}
