package models

import "gorm.io/gorm"

type Books struct {
	gorm.Model
	UserID      uint      `json:"user_id" form:"user_id" binding:"required"`
	User        Users     `json:"user" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Name        string    `json:"name" form:"name" binding:"required"`
	Stock       int       `json:"stock" form:"stock" binding:"required"`
	Image       string    `json:"image" form:"image" binding:"required"`
	Description string    `json:"description" form:"description" binding:"required"`
	CategoryID  uint      `json:"category_id" form:"category_id" binding:"required"`
	Category    Categorys `json:"category" gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type InputBook struct {
	Name        string `json:"name" form:"name" binding:"required"`
	Stock       int    `json:"stock" form:"stock" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	CategoryID  uint   `json:"category_id" form:"category_id" binding:"required"`
}
