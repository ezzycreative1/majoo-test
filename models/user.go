package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

// Users ..
type Users struct {
	User []User `json:"users"`
}

// User ..
type User struct {
	ID          int64          `gorm:"primary_key"; "AUTO_INCREMENT" json:"id"`
	MerchantID  int64          `gorm:"not null;index" json:"merchant_id"`
	OutletID    int64          `gorm:"not null;index" json:"outlet_id"`
	Fullname    string         `gorm:"not null" json:"fullname"`
	Email       string         `gorm:"not null,unique" json:"email,omitempty"`
	Password    string         `gorm:"not null" json:"password,omitempty"`
	PhoneNumber string         `json:"phone_number"`
	Address     string         `json:"address"`
	Status      string         `json:"status"`
	CreatedAt   time.Time      `gorm:"null" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"null" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

// Claims ..
type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// UserResponse ...
type UserResponse struct {
	Token string `json:"token"`
}

// Permission ..
type Permission struct {
	ID        int64          `gorm:"primary_key"; "AUTO_INCREMENT"`
	Name      string         `gorm:"not null"`
	Editor    bool           `gorm:"default:0"`
	Viewer    bool           `gorm:"default:0"`
	CreatedAt time.Time      `gorm:"not null"`
	UpdatedAt time.Time      `gorm:"not null"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
}

// AccessToken ..
type AccessToken struct {
	Token  string
	Expiry int64
	Name   string
	ID     string
}
