package service_object

import (
	"github.com/gin-gonic/gin"
	"github.com/leedrum/ikarus_travel/internal"
	"github.com/leedrum/ikarus_travel/model"
	"gorm.io/gorm"
)

type DropDownReservations struct {
	model.TourItem
	Tour          model.Tour
	Reservations  []model.Reservation
	TotalAdults   int
	TotalChildren int
	TotalPrice    int
	TotalPaidUSD  int
	TotalPaidVND  int
}

func LoadDropDownReservations(ctx *gin.Context, server internal.Server) []DropDownReservations {

	var dropDownReservations []DropDownReservations
	var reservations []model.Reservation
	var tourItems []model.TourItem
	var tours []model.Tour

	tx := server.DB
	txReservation := searchConditions(ctx, tx)
	tx.Find(&tourItems)
	reservationIDs := []int{}
	tourIDs := []int{}

	for _, item := range tourItems {
		reservationIDs = append(reservationIDs, item.ID)
		tourIDs = append(tourIDs, item.TourID)
	}

	if len(tourIDs) > 0 {
		tx.Where("id IN (?)", tourIDs).Find(&tours)
	}

	if len(reservationIDs) > 0 {
		txReservation.Preload(
			"Hotel").Preload(
			"TourItem").Preload(
			"Payments").Where(
			"tour_item_id IN (?)", reservationIDs).Find(&reservations)
	}

	for _, reservation := range reservations {
		found := mappingDropDownData(dropDownReservations, reservation)
		reservationTour := getTour(tours, reservation.TourItem)

		if !found {
			dropDownReservations = append(dropDownReservations, LoadReservation(reservation, reservationTour))
		}
	}

	return dropDownReservations
}

func LoadReservation(reservation model.Reservation, tour model.Tour) DropDownReservations {
	return DropDownReservations{
		TourItem: model.TourItem{
			ID:            reservation.ID,
			TourID:        tour.ID,
			DepartureDate: reservation.DepartureDate,
		},
		Tour:          tour,
		Reservations:  []model.Reservation{reservation},
		TotalAdults:   reservation.Adults,
		TotalChildren: reservation.Children,
		TotalPrice:    reservation.TotalPrice(),
		TotalPaidUSD:  reservation.GetPaidUSD(),
		TotalPaidVND:  reservation.GetPaidVND(),
	}
}

func searchConditions(ctx *gin.Context, tx *gorm.DB) *gorm.DB {
	if ctx.Query("from_date") != "" {
		tx = tx.Where("departure_date >= ?", ctx.Query("from_date"))
	}
	if ctx.Query("departure_date") != "" {
		tx = tx.Where("departure_date = ?", ctx.Query("departure_date"))
	}
	return tx
}

func mappingDropDownData(dropDownReservations []DropDownReservations, reservation model.Reservation) bool {
	var found bool
	for i, dropDownReservation := range dropDownReservations {
		if dropDownReservation.DepartureDate == reservation.DepartureDate {
			dropDownReservations[i].Reservations = append(dropDownReservation.Reservations, reservation)
			dropDownReservations[i].TotalAdults += reservation.Adults
			dropDownReservations[i].TotalChildren += reservation.Children
			dropDownReservations[i].TotalPaidUSD += reservation.GetPaidUSD()
			dropDownReservations[i].TotalPaidVND += reservation.GetPaidVND()
			found = true
			break
		}
	}

	return found
}

func getTour(tours []model.Tour, tourItem model.TourItem) model.Tour {
	for _, tour := range tours {
		if tour.ID == tourItem.TourID {
			return tour
		}
	}

	return model.Tour{}
}
