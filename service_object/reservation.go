package service_object

import (
	"slices"
	"time"

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
	TotalPrice    float64
	TotalPaidUSD  float64
	TotalPaidVND  float64
}

func LoadDropDownReservations(ctx *gin.Context, server internal.Server) []DropDownReservations {

	var dropDownReservations []DropDownReservations
	var reservations []model.Reservation
	var tourItems []model.TourItem
	var tours []model.Tour

	tx := server.DB
	searchConditions(ctx, tx).Order("departure_date DESC").Find(&tourItems)
	reservationIDs := []int{}
	tourIDs := []int{}

	for _, item := range tourItems {
		if !slices.Contains(reservationIDs, item.ID) {
			reservationIDs = append(reservationIDs, item.ID)
		}

		if !slices.Contains(tourIDs, item.TourID) {
			tourIDs = append(tourIDs, item.TourID)
		}
	}

	if len(tourIDs) > 0 {
		tx.Where("id IN (?)", tourIDs).Find(&tours)
	}

	if len(reservationIDs) > 0 {
		tx.Preload(
			"Hotel").Preload(
			"TourItem").Preload(
			"Payments").Preload(
			"User").Where(
			"tour_item_id IN (?)", reservationIDs).Find(&reservations)
	}

	for _, tourItem := range tourItems {
		found := false
		for _, dropDownReservation := range dropDownReservations {
			if dropDownReservation.DepartureDate == tourItem.DepartureDate {
				found = true
				break
			}
		}

		if !found {
			dropDownReservations = append(dropDownReservations, DropDownReservations{
				TourItem: tourItem,
				Tour:     getTour(tours, tourItem),
			})
		}
	}

	for _, reservation := range reservations {
		mappingDropDownData(dropDownReservations, reservation)
	}

	return dropDownReservations
}

func searchConditions(ctx *gin.Context, tx *gorm.DB) *gorm.DB {
	if ctx.Query("from_date") != "" {
		tx = tx.Where("departure_date >= ?", ctx.Query("from_date"))
	} else if ctx.Query("departure_date") != "" {
		tx = tx.Where("departure_date = ?", ctx.Query("departure_date"))
	} else {
		tx = tx.Where("departure_date >= ?", time.Now().AddDate(0, 0, -7).Format("02/01/2006"))
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
