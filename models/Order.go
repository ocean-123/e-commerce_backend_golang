package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID     uint `json:"user_id"`
	User       User `gorm:"foreignKey:UserID"`
	OrderItems []OrderItem
	Total      float64 `json:"total"`
}
