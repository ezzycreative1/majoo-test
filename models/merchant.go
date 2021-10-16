package models

import (
	"time"

	"gorm.io/gorm"
)

type Merchant struct {
	ID        int64          `gorm:"primary_ke;index" "AUTO_INCREMENT" json:"id"`
	Name      string         `gorm:"not null" json:"name"`
	Address   string         `json:"address"`
	CreatedAt time.Time      `gorm:"null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
