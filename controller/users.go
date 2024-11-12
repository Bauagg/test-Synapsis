package controller

import (
	"books/databases"
	"books/models"
	"books/utils"
	"regexp"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(ctx *gin.Context) {
	var input models.InputRegister
	var user models.Users

	if errorInput := ctx.ShouldBind(&input); errorInput != nil {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": errorInput.Error(),
		})
		return
	}

	// Validasi format email langsung di sini
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9.!#$%&'*+/=?^_` + "`{|}~^-]+@[a-zA-Z0-9-]+(?:\\.[a-zA-Z0-9-]+)*$")
	if !emailRegex.MatchString(input.Email) {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": "Format email tidak valid.",
		})
		return
	}

	err := databases.DB.Table("users").Where("email = ?", input.Email).First(&user).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		// Log internal errors for debugging purposes
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Internal server error",
		})
		return
	}

	if err == nil {
		// Email is already registered
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": "Email sudah terdaftar.",
		})
		return
	}

	// Validasi panjang password
	if len(input.Password) < 8 {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": "Password harus memiliki minimal 8 karakter.",
		})
		return
	}

	if input.Role != "user" && input.Role != "admin" {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": "Role harus berupa 'user' atau 'admin'.",
		})
		return
	}

	user.Email = input.Email
	user.Name = input.Name
	user.Password = *utils.HashPassword(&input.Password)
	user.Role = input.Role

	if err := databases.DB.Table("users").Create(&user).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Internal server error",
		})
		return
	}

	ctx.JSON(201, gin.H{
		"error":   false,
		"message": "Register success",
		"data": gin.H{
			"id":    user.ID,
			"email": user.Email,
			"name":  user.Name,
			"role":  user.Role,
		},
	})
}

func Login(ctx *gin.Context) {
	var input models.InputLogin
	var user models.Users

	if errInput := ctx.ShouldBind(&input); errInput != nil {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": errInput.Error(),
		})
		return
	}

	err := databases.DB.Table("users").Where("email = ?", input.Email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(401, gin.H{ // Status 401 untuk Unauthorized
				"error":   true,
				"message": "Invalid email or password.",
			})
			return
		}

		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Internal server error",
		})
		return
	}

	err = utils.VerifikasiHashPassword(&input.Password, &user.Password)
	if err != nil {
		ctx.JSON(401, gin.H{ // Status 401 untuk Unauthorized
			"error":   true,
			"message": "Invalid email or password.",
		})
		return
	}

	token, err := utils.SignToken(uint64(user.ID), user.Email, string(user.Role))
	if err != nil {
		ctx.JSON(500, gin.H{ // Status 500 for Internal Server Error
			"error":   true,
			"message": "Failed to generate token.",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "login success",
		"data": gin.H{
			"name":  user.Name,
			"email": user.Email,
			"role":  user.Role,
			"token": token,
		},
	})
}
