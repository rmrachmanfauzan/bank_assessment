package db

import (
    "log"

    "github.com/rmrachmanfauzan/bank_assessment/internal/model"
)

// RunMigrations migrates the database schemas
func RunMigrations() {
    err := DB.AutoMigrate(
        &model.User{},
        &model.Account{},
    )
    if err != nil {
        log.Fatalf("❌ Migration failed: %v", err)
    }

    log.Println("✅ Database migrated successfully")
}
