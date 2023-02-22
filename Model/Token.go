package Model

import "time"

type Token struct {
	ID         uint `gorm:primaryKey"`
	Token      string
	IsValid    byte
	CustomerID uint `gorm:"foreignKey:CustomerRefer"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}