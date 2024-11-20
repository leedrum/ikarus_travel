package model

import "gorm.io/gorm"

type TourItem struct {
	gorm.Model
	ID            int           `json:"id" form:"id"`
	TourID        int           `json:"tour_id" form:"tour_id"`
	DepartureDate string        `json:"departure_date" form:"departure_date"`
	Tour          Tour          `gorm:"references:ID"`
	Reservations  []Reservation `json:"reservations" gorm:"foreignKey:TourItemID"`
}
