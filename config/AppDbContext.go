package Config

import (
	fmt "fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func SetupDatabase() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dsn := os.Getenv("DB_LINK")
	fmt.Printf("DB_LINK value is: %s\n", dsn)
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	return db
}