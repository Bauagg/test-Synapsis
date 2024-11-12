package controller

import (
	"books/databases"
	"books/models"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListRentalBook(ctx *gin.Context) {
	var data []models.Rental
	userID, _ := ctx.Get("userID")

	if err := databases.DB.Table("rentals").Where("user_id = ?", userID).Preload("Book").Find(&data).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Internal server error",
		})
		return
	}

	var rentalResponses []models.RentalResponse
	for _, rental := range data {
		rentalResponse := models.RentalResponse{
			ID:     rental.ID,
			UserID: rental.UserID,
			BookID: rental.BookID,
			Book: models.BookResponse{
				ID:          rental.Book.ID,
				Name:        rental.Book.Name,
				Stock:       rental.Book.Stock,
				Image:       rental.Book.Image,
				Description: rental.Book.Description,
				CategoryID:  rental.Book.CategoryID,
				Category:    rental.Book.Category,
			},
			Status:       rental.Status,
			DurationDays: rental.DurationDays,
			Qty:          rental.Qty,
			CreatedAt:    rental.CreatedAt,
			UpdatedAt:    rental.UpdatedAt,
		}
		rentalResponses = append(rentalResponses, rentalResponse)
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "list data Rental Book success",
		"data":    rentalResponses,
	})
}

func CreateRentalBook(ctx *gin.Context) {
	var input models.InputRental
	var book models.Books
	userID, _ := ctx.Get("userID")
	role, _ := ctx.Get("userRole")

	if errInput := ctx.ShouldBind(&input); errInput != nil {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": errInput.Error(),
		})
		return
	}

	if role != models.RoleUser {
		ctx.JSON(401, gin.H{
			"error":   true,
			"message": "Unauthorized access: Only users with the 'user' role can create a rental.",
		})
	}

	if err := databases.DB.Table("books").Where("id = ?", input.BookID).First(&book).Error; err != nil {
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

	// Check stock availability
	if book.Stock < int(input.Qty) {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": "Not enough stock available",
		})
		return
	}

	userIDUint, _ := strconv.ParseUint(fmt.Sprintf("%v", userID), 10, 32)

	data := models.Rental{
		UserID:       uint(userIDUint),
		BookID:       input.BookID,
		Status:       models.StatusSewa,
		DurationDays: time.Now().Add(time.Duration(input.DurationDays*24) * time.Hour),
		Qty:          input.Qty,
	}

	if err := databases.DB.Table("rentals").Create(&data).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Internal server error",
		})
		return
	}

	book.Stock -= int(input.Qty)

	if err := databases.DB.Table("books").Save(&book).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Internal server error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "create data Rental Book success",
		"data":    data,
	})
}

func UpdateRentalBook(ctx *gin.Context) {
	var data models.Rental
	var book models.Books
	userID, _ := ctx.Get("userID")
	role, _ := ctx.Get("userRole")

	if role != "user" {
		ctx.JSON(401, gin.H{
			"error":   true,
			"message": "Unauthorized access: Only users with the 'user' role can update a rental.",
		})
	}

	if err := databases.DB.Table("rentals").Where("id = ? AND user_id = ?", ctx.Param("id"), userID).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(404, gin.H{
				"error":   true,
				"message": "data Rental Book Not Found",
			})
			return
		}

		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Internal server error",
		})
		return
	}

	if data.Status == models.StatusDikembalikan {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": "book id " + ctx.Param("id") + " sudah di kembalikan",
		})
		return
	}

	data.Status = models.StatusDikembalikan
	if err := databases.DB.Table("rentals").Save(&data).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Internal server error",
		})
		return
	}

	if err := databases.DB.Table("books").Where("id = ?", data.BookID).First(&book).Error; err != nil {
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

	book.Stock += int(data.Qty)

	if err := databases.DB.Table("books").Save(&book).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Internal server error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "update data success",
		"data":    data,
	})
}

func DeleteRental(ctx *gin.Context) {
	var data models.Rental
	userID, _ := ctx.Get("userID")
	role, _ := ctx.Get("userRole")

	if role != "user" {
		ctx.JSON(401, gin.H{
			"error":   true,
			"message": "Unauthorized access: Only users with the 'user' role can update a rental.",
		})
	}

	if err := databases.DB.Table("rentals").Where("id = ? AND user_id = ?", ctx.Param("id"), userID).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(404, gin.H{
				"error":   true,
				"message": "data Rental Book Not Found",
			})
			return
		}

		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Internal server error",
		})
		return
	}

	if data.Status != models.StatusDikembalikan {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": "book id " + ctx.Param("id") + " belum di kembalikan",
		})
		return
	}

	if err := databases.DB.Table("rentals").Delete(&data).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Internal server error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "delete data Rental success",
	})
}
