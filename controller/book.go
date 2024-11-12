package controller

import (
	"books/config"
	"books/databases"
	"books/models"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListBook(ctx *gin.Context) {
	var data []models.Books
	nameQuery := ctx.Query("name")
	categoryIDQuery := ctx.Query("category_id")

	query := databases.DB.Table("books").
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name", "role", "email", "created_at", "updated_at")
		}).
		Preload("Category")

		// Apply case-insensitive filtering on the book name if a name query is provided
	if nameQuery != "" {
		query = query.Where("LOWER(name) LIKE ?", "%"+nameQuery+"%")
	}

	// Apply filtering based on category_id if provided
	if categoryIDQuery != "" {
		query = query.Where("category_id = ?", categoryIDQuery)
	}

	// Execute the query
	if err := query.Find(&data).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Internal server error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "list data Book success",
		"data":    data,
	})
}

func CreateBook(ctx *gin.Context) {
	var input models.InputBook
	userID, _ := ctx.Get("userID")
	role, _ := ctx.Get("userRole")

	if errInput := ctx.ShouldBind(&input); errInput != nil {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": errInput.Error(),
		})
		return
	}

	if role != "admin" {
		ctx.JSON(401, gin.H{
			"error":   true,
			"message": "Unauthorized access: Admin privileges are required to create a book entry.",
		})
		return
	}

	file, err := ctx.FormFile("image")

	if err != nil {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": "Failed to upload image: " + err.Error(),
		})
		return
	}

	// Create directory if not exists
	imageDir := "./images-book"
	if _, err := os.Stat(imageDir); os.IsNotExist(err) {
		err = os.MkdirAll(imageDir, os.ModePerm)
		if err != nil {
			ctx.JSON(500, gin.H{
				"error":   true,
				"message": "Failed to create image directory: " + err.Error(),
			})
			return
		}
	}

	// Generate unique filename based on timestamp
	fileName := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
	filepath := filepath.Join(imageDir, fileName)

	// Save the file to the server
	if err := ctx.SaveUploadedFile(file, filepath); err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Failed to save image: " + err.Error(),
		})
		return
	}

	userIDUint, _ := strconv.ParseUint(fmt.Sprintf("%v", userID), 10, 32)

	data := models.Books{
		Name:        input.Name,
		Stock:       input.Stock,
		Description: input.Description,
		UserID:      uint(userIDUint),
		CategoryID:  input.CategoryID,
		Image:       config.URL_HOST + "/images/" + fileName,
	}

	if err := databases.DB.Table("books").Create(&data).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Internal server error",
		})
		return
	}

	ctx.JSON(201, gin.H{
		"error":   false,
		"message": "create data book success",
		"data":    data,
	})
}

func UpdateBook(ctx *gin.Context) {
	var input models.InputBook
	var data models.Books
	userID, _ := ctx.Get("userID")
	role, _ := ctx.Get("userRole")

	if errInput := ctx.ShouldBind(&input); errInput != nil {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": errInput.Error(),
		})
		return
	}

	if role != "admin" {
		ctx.JSON(401, gin.H{
			"error":   true,
			"message": "Unauthorized access: Admin privileges are required to update a book entry.",
		})
		return
	}

	if err := databases.DB.Table("books").Where("id = ? AND user_id = ?", ctx.Param("id"), userID).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(404, gin.H{
				"error":   true,
				"message": "data Book Not Found",
			})
			return
		}

		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Internal server error",
		})
		return
	}

	// Handle image upload if a new image file is provided
	file, _ := ctx.FormFile("image")
	imageDir := "./images-book"
	if file != nil {
		if data.Image != "" {
			// Get the old image file name and delete it
			fileName := filepath.Base(data.Image)
			oldFilePath := filepath.Join(imageDir, fileName)

			if err := os.Remove(oldFilePath); err != nil && !os.IsNotExist(err) {
				ctx.JSON(500, gin.H{
					"error":   true,
					"message": "Failed to delete old image: " + err.Error(),
				})
				return
			}
		}

		// Ensure the image directory exists
		if _, err := os.Stat(imageDir); os.IsNotExist(err) {
			err = os.MkdirAll(imageDir, os.ModePerm)
			if err != nil {
				ctx.JSON(500, gin.H{
					"error":   true,
					"message": "Failed to create image directory: " + err.Error(),
				})
				return
			}
		}

		// Generate a unique filename and save the new image
		fileName := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
		filepath := filepath.Join(imageDir, fileName)

		if err := ctx.SaveUploadedFile(file, filepath); err != nil {
			ctx.JSON(500, gin.H{
				"error":   true,
				"message": "Failed to save image: " + err.Error(),
			})
			return
		}

		data.Image = config.URL_HOST + "/images/" + fileName
	}

	data.Name = input.Name
	data.Stock = input.Stock
	data.Description = input.Description
	data.CategoryID = input.CategoryID

	if err := databases.DB.Table("books").Save(&data).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Failed to update book.",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "Book updated successfully.",
		"data":    data,
	})
}

func DeleteBook(ctx *gin.Context) {
	var data models.Books
	userID, _ := ctx.Get("userID")
	role, _ := ctx.Get("userRole")

	if role != "admin" {
		ctx.JSON(401, gin.H{
			"error":   true,
			"message": "Unauthorized access: Admin privileges are required to delete a book entry.",
		})
		return
	}

	if err := databases.DB.Table("books").Where("id = ? AND user_id = ?", ctx.Param("id"), userID).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(404, gin.H{
				"error":   true,
				"message": "data Book Not Found",
			})
			return
		}

		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Internal server error",
		})
		return
	}

	// Delete the associated image file if it exists
	if data.Image != "" {
		imageDir := "./images-book"
		fileName := filepath.Base(data.Image)
		imagePath := filepath.Join(imageDir, fileName)

		if err := os.Remove(imagePath); err != nil && !os.IsNotExist(err) {
			ctx.JSON(500, gin.H{
				"error":   true,
				"message": "Failed to delete book image: " + err.Error(),
			})
			return
		}
	}

	if err := databases.DB.Table("books").Delete(&data).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Internal server error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "Book deleted successfully",
	})
}
