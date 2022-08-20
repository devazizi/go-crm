package entity

import (
	"gorm.io/gorm"
	"time"
)

type Token struct {
	gorm.Model
	UserID     int       `json:"user_id"`
	Hash       string    `json:"hash"`
	ExpiryDate time.Time `json:"expiry_date"`
	User       User
}
