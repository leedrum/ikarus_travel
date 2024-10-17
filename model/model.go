package model

import "gorm.io/gorm"

type Model interface {
	gorm.Model
}
