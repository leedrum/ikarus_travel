package model

import "gorm.io/gorm"

type Hotel struct {
	gorm.Model
	ID          int    `gorm:"primaryKey" json:"id"`
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	Address     string `json:"address" form:"address"`
}
