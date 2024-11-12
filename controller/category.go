package controller

import (
	"books/databases"
	"books/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListCategory(ctx *gin.Context) {
	var data []models.Categorys

	if err := databases.DB.Table("categorys").Find(&data).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Internal server error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "list data Category success",
		"data":    data,
	})
}

func CreateCategory(ctx *gin.Context) {
	var data models.Categorys

	if errInput := ctx.ShouldBind(&data); errInput != nil {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": errInput.Error(),
		})
		return
	}

	if err := databases.DB.Table("categorys").Create(&data).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Internal server error",
		})
		return
	}

	ctx.JSON(201, gin.H{
		"error":   false,
		"message": "create data Category success",
		"data":    data,
	})
}

func UpdateCategory(ctx *gin.Context) {
	var input models.InputCategorys
	var data models.Categorys

	if errInput := ctx.ShouldBind(&input); errInput != nil {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": errInput.Error(),
		})
		return
	}

	if err := databases.DB.Table("categorys").Where("id = ?", ctx.Param("id")).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(404, gin.H{
				"error":   true,
				"message": "Data Category Not Found",
			})
			return
		}

		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Internal server error",
		})
		return
	}

	// Updating the category name
	data.Name = input.Name

	// Saving the updated category
	if err := databases.DB.Save(&data).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Failed to update category",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "Category updated successfully",
		"data":    data,
	})
}

func DeleteCategory(ctx *gin.Context) {
	var data models.Categorys

	if err := databases.DB.Table("categorys").Where("id = ?", ctx.Param("id")).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(404, gin.H{
				"error":   true,
				"message": "Data Category Not Found",
			})
			return
		}

		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Internal server error",
		})
		return
	}

	if err := databases.DB.Table("categorys").Where("id = ?", ctx.Param("id")).Delete(&data).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Internal server error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "delete data Category success",
	})
}
