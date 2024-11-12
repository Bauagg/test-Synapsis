package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DB_HOST         string
	DB_PORT         string
	DB_NAMA         string
	DB_USER         string
	DB_PASSWORD     string
	APP_PORT        string
	DB_TIMEZONE     string
	SECRETKEY_TOKEN string
	URL_HOST        string
)

func IntConfigEnv() {
	// Load .env file if exists
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, proceeding with environment variables")
	}

	// Fetch values from environment variables
	appPort := os.Getenv("APP_PORT")
	if appPort != "" {
		APP_PORT = appPort
	}

	// databases
	dbHost := os.Getenv("DB_HOST")
	if dbHost != "" {
		DB_HOST = dbHost
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort != "" {
		DB_PORT = dbPort
	}

	dbName := os.Getenv("DB_NAME")
	if dbName != "" {
		DB_NAMA = dbName
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser != "" {
		DB_USER = dbUser
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword != "" {
		DB_PASSWORD = dbPassword
	}

	dbTimezone := os.Getenv("DB_TIMEZONE")
	if dbTimezone != "" {
		DB_TIMEZONE = dbTimezone
	}

	secretToken := os.Getenv("SECRETKEY_TOKEN")
	if secretToken != "" {
		SECRETKEY_TOKEN = secretToken
	}

	urlHost := os.Getenv("URL_HOST")
	if urlHost != "" {
		URL_HOST = urlHost
	}
}
