package model

import (
	"time"

	"gorm.io/gorm"
)

const (
	CurrencyUSD = iota
	CurrencyVND
)

const (
	PaymentMethodCash = iota
	PaymentMethodBankTransfer
	PaymentMethodCreditCard
	PaymentMethodOther
)

type Payment struct {
	gorm.Model
	ID            int         `json:"id"`
	ReservationID int         `json:"reservation_id" form:"reservation_id" gorm:"index"`
	Reservation   Reservation `json:"reservation" gorm:"references:ID"`
	Amount        float64     `json:"amount" form:"amount"`
	Currency      int         `json:"currency" form:"currency"`
	PaymentMethod int         `json:"payment_method" form:"payment_method"`
	Note          string      `json:"note" form:"note"`
	UserID        int         `json:"user_id" reference:"ID" gorm:"index"`
	User          User        `json:"user" gorm:"references:ID"`
	CreatedAt     time.Time   `json:"created_at"`
}

func GetPaymentHistory(reservationID []int, db *gorm.DB) []Payment {
	var payments []Payment
	if len(reservationID) == 1 {
		db.Preload("User").Where("reservation_id = ?", reservationID[0]).Find(&payments)
	} else {
		db.Preload("User").Where("reservation_id IN (?)", reservationID).Find(&payments)
	}

	return payments
}

func (p *Payment) GetCurrency() string {
	switch p.Currency {
	case CurrencyUSD:
		return "USD"
	case CurrencyVND:
		return "VND"
	default:
		return "Unknown"
	}
}

func (p Payment) GetPaymentMethod() string {
	switch p.PaymentMethod {
	case PaymentMethodCash:
		return "Cash"
	case PaymentMethodBankTransfer:
		return "Bank Transfer"
	case PaymentMethodCreditCard:
		return "Credit Card"
	case PaymentMethodOther:
		return "Other"
	default:
		return "Unknown"
	}
}
