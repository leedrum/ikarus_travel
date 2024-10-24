package model

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	ReservationStatusInActive = 0
	ReservationStatusActive   = 1
	ReservationStatusCanceled = 2

	// Payment status
	PaymentStatusUnpaid  = 0
	PaymentStatusPaid    = 1 // paid 100%
	PaymentStatusDeposit = 2 // paid deposit
	PaymentStatusLater   = 3 // pay later
	PaymentStatusToGuide = 4 // pay to guide
	PaymentOther         = 5 // other
)

type Reservation struct {
	gorm.Model
	ID            int    `gorm:"primaryKey" json:"id"`
	Code          string `json:"code" form:"code" gorm:"unique;index"`
	TourID        int    `json:"tour_id" form:"tour_id"` // TourID is the foreign key
	Tour          Tour   `json:"tour" form:"tour"`
	CustomerName  string `json:"customer_name" form:"customer_name"`
	Phone         string `json:"phone" form:"phone"`
	HotelID       int    `json:"hotel_id" form:"hotel_id"` // HotelID is the foreign key
	Hotel         Hotel  `json:"hotel" form:"hotel"`
	RoomNote      string `json:"room_note" form:"room_note"`
	Adults        int    `json:"adults" form:"adults"`
	AdultsPrice   int    `json:"adults_price" form:"adults_price"`
	Children      int    `json:"children" form:"children"`
	ChildrenPrice int    `json:"children_price" form:"children_price"`
	DepartureDate string `json:"departure_date" form:"departure_date"`
	PickupTime    string `json:"pickup_time" form:"pickup_time"`
	StartTime     string `json:"start_time" form:"start_time"`
	EndTime       string `json:"end_time" form:"end_time"`
	Status        int    `json:"status" form:"status" gorm:"default:1"`
	PaymentStatus int    `json:"payment_status" form:"payment_status" gorm:"default:0"`
	TotalPaid     int    `json:"total_paid" form:"total_paid" gorm:"default:0"`
	CurrencyPaid  string `json:"currency_paid" form:"currency_paid"`
	PaymentNote   string `json:"payment_note" form:"payment_note"`
	Note          string `json:"note" form:"note"`
}

func (u *Reservation) BeforeCreate(tx *gorm.DB) (err error) {
	uid := uuid.New()
	u.Code = "RES-" + uid.String()
	fmt.Println("BeforeCreate", u.Code)
	return nil
}

func (r *Reservation) GetStatusStr(n int) string {
	status := r.Status
	if n != 0 {
		status = n
	}

	switch status {
	case ReservationStatusInActive:
		return "InActive"
	case ReservationStatusActive:
		return "Active"
	case ReservationStatusCanceled:
		return "Canceled"
	default:
		return "Unknown"
	}
}

func (r *Reservation) GetPaymentStatusStr(n int) string {
	paymentStatus := r.PaymentStatus
	if n != 0 {
		paymentStatus = n
	}

	switch paymentStatus {
	case PaymentStatusUnpaid:
		return "Unpaid"
	case PaymentStatusPaid:
		return "Paid"
	case PaymentStatusDeposit:
		return "Deposit"
	case PaymentStatusLater:
		return "Pay Later"
	case PaymentStatusToGuide:
		return "Pay To Guide"
	case PaymentOther:
		return "Other"
	default:
		return "Unknown"
	}
}
