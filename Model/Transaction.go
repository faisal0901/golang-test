package Model

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ID          uint     `gorm:primaryKey"`
	Qty        float32
	Price      float32
	CustomerID   uint   `gorm:"foreignKey:CustomerRefer"`
	ProductID   uint   `gorm:"foreignKey:ProductRefer"`
	CreatedAt time.Time
    UpdatedAt time.Time
}