package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func Init() *gorm.DB {
	var err error
	var DB *gorm.DB
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		os.Getenv("ATOMONO_DB_HOST"),
		os.Getenv("ATOMONO_DB_USER"),
		os.Getenv("ATOMONO_DB_PASSWORD"),
		os.Getenv("ATOMONO_DB_NAME"),
		os.Getenv("ATOMONO_DB_PORT"),
	)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

	fmt.Println("Database connected")
	
	return DB
}