package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	FirstName       string    `json:"first_name" gorm:"not null"`
	LastName        string    `json:"last_name" gorm:"not null"`
	Email           string    `gorm:"unique;not null"`
	Password         string    `gorm:"not null" json:"password"`
    ResetToken       string    `gorm:"default:null" json:"reset_token"`
    ResetTokenExpiry *time.Time `gorm:"default:null" json:"reset_token_expiry"`
}