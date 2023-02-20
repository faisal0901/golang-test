package Model

import (
	"time"

	"gorm.io/gorm"
)
type Merchant struct {
	gorm.Model
	ID          uint     `gorm:primaryKey"`
	Name    string
	Address string
	Phone   string
	Products []Product
	CreatedAt time.Time
    UpdatedAt time.Time
}
