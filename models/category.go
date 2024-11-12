package models

import "gorm.io/gorm"

type Categorys struct {
	gorm.Model
	Name string `json:"name" binding:"required"`
}

type InputCategorys struct {
	Name string `json:"name" binding:"required"`
}
