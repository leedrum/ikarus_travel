package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID                 int `gorm:"primaryKey"`
	Username           string
	HashPassword       string
	FullName           string
	Email              string
	PasswordChangegdAt time.Time
	Role               string
	CreatedAt          time.Time
}
