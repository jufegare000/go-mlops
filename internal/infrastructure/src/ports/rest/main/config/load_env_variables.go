package config

import (
	"github.com/joho/godotenv"
	"log"
)

func SetUpEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("❌ Error loading .env: %v", err)
	}
}
