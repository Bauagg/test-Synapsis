package models

import "gorm.io/gorm"

type Role string

var (
	RoleUser  Role = "user"
	RoleAdmin Role = "admin"
)

type Users struct {
	gorm.Model
	Name     string `json:"name" binding:"required"`
	Role     Role   `json:"role" binding:"required" gorm:"type:text;default:'user'"`
	Email    string `json:"email" binding:"required" gorm:"unique"`
	Password string `json:"password" binding:"required"`
}

type InputRegister struct {
	Name     string `json:"name" binding:"required"`
	Role     Role   `json:"role" binding:"required" gorm:"type:text;default:'user'"`
	Email    string `json:"email" binding:"required" gorm:"unique"`
	Password string `json:"password" binding:"required"`
}

type InputLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
