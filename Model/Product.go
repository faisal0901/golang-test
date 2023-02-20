package Model

import (
	"time"

	"gorm.io/gorm"
)

type Product struct { 
	gorm.Model
	ID          uint   `gorm:primaryKey"`
	Name        string
	Description string
	MerchantID  uint   `gorm:"foreignKey:MerchantRefer"`
	Price       int 	
	Transaction []Transaction
	CreatedAt time.Time
    UpdatedAt time.Time
  } 

