package config

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	db_url := os.Getenv("DB_URL")
	if db_url == "" {
		log.Fatal("Environment variable DB_URL is not set")
	}

	log.Println("Connecting to database:", db_url) // Debugging log

	dbConn, err := gorm.Open(postgres.Open(db_url), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to Database:", err)
	}

	DB = dbConn
	log.Println("Database connected successfully")
}
