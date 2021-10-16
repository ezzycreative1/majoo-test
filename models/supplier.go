package models

import (
	"time"

	"gorm.io/gorm"
)

type Supplier struct {
	gorm.Model
	ID          int64          `gorm:"primary_ke;index" "AUTO_INCREMENT" json:"id"`
	Name        string         `gorm:"not null" json:"name"`
	Email       string         `gorm:"not null,unique" json:"email,omitempty"`
	PhoneNumber string         `json:"phone_number"`
	Address     string         `json:"address"`
	Status      string         `json:"status"`
	CreatedAt   time.Time      `gorm:"null" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"null" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
