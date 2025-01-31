package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVars() {
	if os.Getenv("RAILWAY_ENVIRONMENT") == "" {
		err := godotenv.Load()
		if err != nil {
			log.Println("No .env file found, using system environment variables")
		}
	}
}
