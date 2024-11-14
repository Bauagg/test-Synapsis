package main

import (
	"books/config"
	"books/databases"
	middleware "books/middelware"
	"books/migrations"
	"books/routers"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// config env
	config.IntConfigEnv()

	// Connect to the database
	databases.ConnectDatabase()
	migrations.Migrate()

	app := gin.Default()

	// CORS middleware
	app.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodPatch,
			http.MethodHead,
			http.MethodConnect,
			http.MethodTrace,
		},
		AllowHeaders:  []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        12 * time.Hour,
	}))

	// APP middleware
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
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
