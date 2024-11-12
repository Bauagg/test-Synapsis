package migrations

import (
	"books/databases"
	"books/models"
	"fmt"
)

func Migrate() {
	err := databases.DB.AutoMigrate(
		&models.Users{},
		&models.Books{},
		&models.Categorys{},
		&models.Rental{},
	)
	if err != nil {
		fmt.Println("Failed to migrate:", err)
		panic("Migration failed!")
	}
	fmt.Println("Migration successful!")
}
