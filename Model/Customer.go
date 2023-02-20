package Model

import (
	"time"
)

type Customer struct {
    ID        uint     `gorm:primaryKey"`
    Name      string
	Email     string    `gorm:"unique"`
    Password  string
    Phone     string
    Address   string
    Transaction []Transaction
    CreatedAt time.Time
    UpdatedAt time.Time
}