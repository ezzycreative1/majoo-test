package models

import (
	"time"

	"gorm.io/gorm"
)

type PurchaseOrder struct {
	gorm.Model
	ID           int64          `gorm:"primary_key;index" "AUTO_INCREMENT" json:"id"`
	Supplier     []Supplier     `gorm:"many2many:supplier_purchase_orders;"`
	price        int64          `gorm:"not null" json:"price"`
	PurchaseDate time.Time      `gorm:"not null" json:"purchase_date"`
	descripton   string         `gorm:"not null" json:"description"`
	CreatedAt    time.Time      `gorm:"null" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"null" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
