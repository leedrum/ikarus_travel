package model

import (
	"time"

	"gorm.io/gorm"
)

const (
	RoleAdmin = "admin"
	RoleUser  = "user"
)

type User struct {
	gorm.Model
	ID                 int       `gorm:"primaryKey" json:"id"`
	Username           string    `json:"username" form:"username" gorm:"unique;index"`
	HashPassword       string    `json:"hash_password" form:"hash_password"`
	FullName           string    `json:"full_name" form:"full_name"`
	Email              string    `json:"email" form:"email"`
	PasswordChangegdAt time.Time `json:"password_changed_at"`
	Role               string    `json:"role" form:"role"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

func (u *User) IsAdmin() bool {
	return u.Role == RoleAdmin
}

func (u *User) IsUser() bool {
	return u.Role == RoleUser
}
