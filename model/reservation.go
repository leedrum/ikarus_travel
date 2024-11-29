package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	ReservationStatusInActive = iota
	ReservationStatusActive
	ReservationStatusCanceled
)

const (
	// Payment status
	PaymentStatusUnpaid  = iota // unpaid
	PaymentStatusPaid           // paid 100%
	PaymentStatusDeposit        // paid deposit
	PaymentStatusLater          // pay later
	PaymentStatusToGuide        // pay to guide
	PaymentOther                // other
)

type Reservation struct {
	gorm.Model
	ID            int       `gorm:"primaryKey" json:"id"`
	Code          string    `json:"code" form:"code" gorm:"unique;index"`
	CustomerName  string    `json:"customer_name" form:"customer_name"`
	Phone         string    `json:"phone" form:"phone"`
	HotelID       int       `json:"hotel_id" form:"hotel_id"` // HotelID is the foreign key
	RoomNote      string    `json:"room_note" form:"room_note"`
	Adults        int       `json:"adults" form:"adults"`
	AdultsPrice   int       `json:"adults_price" form:"adults_price"`
	Children      int       `json:"children" form:"children"`
	ChildrenPrice int       `json:"children_price" form:"children_price"`
	DepartureDate string    `json:"departure_date" form:"departure_date"`
	PickupTime    string    `json:"pickup_time" form:"pickup_time"`
	StartTime     string    `json:"start_time" form:"start_time"`
	EndTime       string    `json:"end_time" form:"end_time"`
	Status        int       `json:"status" form:"status" gorm:"default:1"`
	PaymentStatus int       `json:"payment_status" form:"payment_status" gorm:"default:0"`
	Note          string    `json:"note" form:"note"`
	UserID        int       `json:"user_id" form:"user_id"`
	TourItemID    int       `json:"tour_item_id"`
	Payments      []Payment `json:"payment" gorm:"foreignKey:ReservationID"`
	User          User      `json:"user" gorm:"references:ID"`
	TourItem      TourItem  `json:"tour_item" gorm:"references:ID"`
	Hotel         Hotel     `json:"hotel" gorm:"references:ID"`
	Tour          Tour      `json:"tour" gorm:"foreignKey:TourItemID;references:ID"`
}

func (u *Reservation) BeforeCreate(tx *gorm.DB) (err error) {
	uid := uuid.New()
	u.Code = "RES-" + uid.String()
	return nil
}

func (r *Reservation) TotalPrice() float64 {
	return float64(r.Adults*r.AdultsPrice + r.Children*r.ChildrenPrice)
}

func (r *Reservation) GetStatus() string {
	return r.GetStatusStr(r.Status)
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

func (r *Reservation) GetPaymentStatus() string {
	return r.GetPaymentStatusStr(r.PaymentStatus)
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

func (r *Reservation) GetPaidUSD() float64 {
	total := 0.0
	for _, payment := range r.Payments {
		if payment.Currency == CurrencyUSD {
			total += payment.Amount
		}
	}
	return total
}

func (r *Reservation) GetPaidVND() float64 {
	total := 0.0
	for _, payment := range r.Payments {
		if payment.Currency == CurrencyVND {
			total += payment.Amount
		}
	}
	return total
}
