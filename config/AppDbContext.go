package Config

import (
	fmt "fmt"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func SetupDatabase() *gorm.DB {
	dsn := "Server=FAISAL;Database=DB_golang_test;Trusted_Connection=True;TrustServerCertificate=True;"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	return db
}