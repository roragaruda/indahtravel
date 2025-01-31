package config

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	db_url := os.Getenv("DB_URL")
	dbConn, err := gorm.Open(postgres.Open(db_url), &gorm.Config{})
	if err != nil {
		panic("Failed connecting to Database")
	}

	DB = dbConn
}
