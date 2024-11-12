package main

import (
	"books/config"
	"books/databases"
	middleware "books/middelware"
	"books/migrations"
	"books/routers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// config env
	config.IntConfigEnv()

	// Connect to the database
	databases.ConnectDatabase()
	migrations.Migrate()

	app := gin.Default()

	// APP middleware
	app.Use(middleware.ErrorMiddleware())

	// Setup static file serving for images
	app.Static("/images", "./images-book")

	// router
	routers.IndexRouter(app)

	err := app.Run(config.APP_PORT)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
