package database

import (
	"atomono-api/internal/models"
	"atomono-api/internal/models/master"

	"log"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
    log.Println("Running database migrations...")

    err := db.AutoMigrate(
        &models.Product{},
        &models.User{},
        &models.Review{},
        &models.ReplacesProduct{},
        &master.Brand{},
        &master.Category{},
        &master.Company{},
        &master.Country{},
    )

    if err != nil {
        log.Fatalf("Error during migration: %v", err)
    }

    log.Println("Database migration completed successfully")
}