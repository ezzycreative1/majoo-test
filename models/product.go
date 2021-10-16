package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID         int64          `gorm:"primary_key;index" "AUTO_INCREMENT" json:"id"`
	Name       string         `gorm:"not null" json:"name"`
	sku        string         `gorm:"not null" json:"sku"`
	image      string         `gorm:"not null" json:"image"`
	price      int64          `gorm:"not null" json:"price"`
	stock      int64          `gorm:"not null" json:"stock"`
	descripton string         `gorm:"not null" json:"description"`
	CreatedAt  time.Time      `gorm:"null" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"null" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
