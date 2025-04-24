package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://postgres:superadmin@localhost:5432/bank_assessment?sslmode=disable"
		log.Println("Warning: DATABASE_URL not found in environment, using default DSN.")
	}


	var db *gorm.DB
	var err error

	// Retry for 10 times to connect db image
	for i := 0; i < 10; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			DB = db
			fmt.Println("Database connection established")
			return
		}
		log.Printf("Retrying database connection (%d/10): %v", i+1, err)
		time.Sleep(3 * time.Second)
	}

	log.Fatalf("Failed to connect to database after retries: %v", err)
}
