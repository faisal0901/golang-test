package Model

type Token struct {
	ID         uint `gorm:primaryKey"`
	Token      string
	IsValid    byte
	CustomerID uint `gorm:"foreignKey:CustomerRefer"`
}