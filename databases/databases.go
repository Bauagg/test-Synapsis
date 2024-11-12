package databases

import (
	"books/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Build DSN (Data Source Name)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		config.DB_HOST, config.DB_USER, config.DB_PASSWORD, config.DB_NAMA, config.DB_PORT, config.DB_TIMEZONE)

	// Connect to the database
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = database
	fmt.Println("Database connected!")
}
