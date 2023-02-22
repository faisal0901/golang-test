package Test

import (
	"context"
	model "go-test/Model"
	repository "go-test/Repository"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func TestCreateData(t *testing.T) {
    // setup mock db
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("Failed to create mock db: %v", err)
    }
    defer db.Close()

    gormDB, err := gorm.Open(sqlserver.New(sqlserver.Config{
        DriverName: "sqlserver",
        DSN:        "Server=FAISAL;Database=DB_golang_test;Trusted_Connection=True;TrustServerCertificate=True;",

    }), &gorm.Config{})
    if err != nil {
        t.Fatalf("Failed to open gorm db: %v", err)
    }

    r := repository.NewRepository(gormDB)

    data := &model.Customer{
        Name:  "John",
        Email: "john@example.com",
    }

  
    mock.ExpectExec(`INSERT INTO "users" ("name","email") VALUES (?,?)`).WithArgs(data.Name, data.Email).WillReturnResult(sqlmock.NewResult(1, 1))
   

    err = r.CreateData(context.Background(), data)
    if err != nil {
        t.Errorf("Failed to create data: %v", err)
    }

   
}


