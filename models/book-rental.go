package models

import (
	"time"

	"gorm.io/gorm"
)

type RentalStatus string

const (
	StatusSewa         RentalStatus = "sewa"
	StatusDikembalikan RentalStatus = "dikembalikan"
)

type Rental struct {
	gorm.Model
	UserID       uint         `json:"user_id" binding:"required"`
	BookID       uint         `json:"book_id" binding:"required"`
	Book         Books        `json:"book" gorm:"foreignKey:BookID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Status       RentalStatus `json:"status" binding:"required" gorm:"type:text"`
	DurationDays time.Time    `json:"duration_days" binding:"required"`
	Qty          uint         `json:"qty" binding:"required"`
}

type InputRental struct {
	BookID       uint `json:"book_id" binding:"required"`
	DurationDays uint `json:"duration_days" binding:"required"`
	Qty          uint `json:"qty" binding:"required"`
}

// respon payload
type BookResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Stock       int       `json:"stock"`
	Image       string    `json:"image"`
	Description string    `json:"description"`
	CategoryID  uint      `json:"category_id"`
	Category    Categorys `json:"category"`
}

type RentalResponse struct {
	ID           uint         `json:"id"`
	UserID       uint         `json:"user_id"`
	BookID       uint         `json:"book_id"`
	Book         BookResponse `json:"book"`
	Status       RentalStatus `json:"status"`
	DurationDays time.Time    `json:"duration_days"`
	Qty          uint         `json:"qty"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
}
