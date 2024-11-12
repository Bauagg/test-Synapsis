package routers

import (
	"books/controller"
	middleware "books/middelware"

	"github.com/gin-gonic/gin"
)

func IndexRouter(app *gin.Engine) {
	router := app

	router.POST("/api/register", controller.Register)
	router.POST("/api/login", controller.Login)

	// category
	router.GET("/api/category", controller.ListCategory)
	router.POST("/api/category", controller.CreateCategory)
	router.PUT("/api/category/:id", controller.UpdateCategory)
	router.DELETE("/api/category/:id", controller.DeleteCategory)

	// book
	router.GET("/api/book", controller.ListBook)
	router.POST("/api/book", middleware.AuthMiddleware(), controller.CreateBook)
	router.PUT("/api/book/:id", middleware.AuthMiddleware(), controller.UpdateBook)
	router.DELETE("/api/book/:id", middleware.AuthMiddleware(), controller.DeleteBook)

	// rental book
	router.GET("/api/rental", middleware.AuthMiddleware(), controller.ListRentalBook)
	router.POST("/api/rental", middleware.AuthMiddleware(), controller.CreateRentalBook)
	router.PUT("/api/rental/:id", middleware.AuthMiddleware(), controller.UpdateRentalBook)
	router.DELETE("/api/rental/:id", middleware.AuthMiddleware(), controller.DeleteRental)
}
