package model

import (
	"time"

	"gorm.io/gorm"
)

type Tour struct {
	gorm.Model
	ID          int        `gorm:"primaryKey" json:"id"`
	Name        string     `json:"name" form:"name"`
	Description string     `json:"description" form:"description"`
	Status      int        `json:"status" form:"status"`
	TourItems   []TourItem `json:"tour_items" gorm:"foreignKey:TourID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

const (
	TourStatusInActive = 0
	TourStatusActive   = 1
)

func (t *Tour) GetStatusStr() string {
	switch t.Status {
	case TourStatusInActive:
		return "InActive"
	case TourStatusActive:
		return "Active"
	default:
		return "Unknown"
	}
}
