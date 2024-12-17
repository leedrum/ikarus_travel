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
	DepartureDate            string
	TourItemID               int
	ItemDropDownReservations []ItemDropDownReservations
	TotalAdults              int
	TotalChildren            int
}

type ItemDropDownReservations struct {
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
	SearchConditions(ctx, tx).Order("departure_date DESC").Find(&tourItems)
	tourItemIds := []int{}
	tourIDs := []int{}

	for _, item := range tourItems {
		if !slices.Contains(tourItemIds, item.ID) {
			tourItemIds = append(tourItemIds, item.ID)
		}

		if !slices.Contains(tourIDs, item.TourID) {
			tourIDs = append(tourIDs, item.TourID)
		}
	}

	if len(tourIDs) > 0 {
		tx.Where("id IN (?)", tourIDs).Find(&tours)
	}

	if len(tourItemIds) > 0 {
		queryTourItems := tx.Preload(
			"Hotel").Preload(
			"TourItem").Preload(
			"Payments").Preload(
			"User").Where(
			"tour_item_id IN (?)", tourItemIds)

		if ctx.Request.URL.Query().Get("status") != "" {
			queryTourItems.Where("status = ?", ctx.Request.URL.Query().Get("status"))
		}
		queryTourItems.Find(&reservations)
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
				DepartureDate: tourItem.DepartureDate,
				TourItemID:    tourItem.ID,
			})
		}
	}

	for i, dropDownReservation := range dropDownReservations {
		for _, tourItem := range tourItems {
			if dropDownReservation.DepartureDate == tourItem.DepartureDate {
				dropDownReservations[i].ItemDropDownReservations = append(dropDownReservations[i].ItemDropDownReservations, ItemDropDownReservations{
					TourItem: tourItem,
					Tour:     getTour(tours, tourItem),
				})
			}
		}
	}

	for _, reservation := range reservations {
		mappingDropDownData(dropDownReservations, reservation)
	}

	return dropDownReservations
}

func SearchConditions(ctx *gin.Context, tx *gorm.DB) *gorm.DB {
	from_date := ctx.Request.URL.Query().Get("from_date")
	to_date := ctx.Request.URL.Query().Get("to_date")
	departure_date := ctx.Request.URL.Query().Get("departure_date")

	if from_date != "" && to_date != "" {
		tx = tx.Where("departure_date >= ?", from_date).Where("departure_date <= ?", to_date)
	} else if departure_date != "" {
		tx = tx.Where("departure_date = ?", departure_date)
	} else {
		tx = tx.Where("departure_date > ?", time.Now().AddDate(0, 0, -8).Format("02/01/2006"))
	}

	if ctx.Query("customer_name") != "" {
		tx = tx.Where("LOWER(customer_name) LIKE LOWER(?)", "%"+ctx.Query("customer_name")+"%")
	}
	return tx
}

func mappingDropDownData(dropDownReservations []DropDownReservations, reservation model.Reservation) {
	for i, dropDownReservation := range dropDownReservations {
		for j, itemDropDownReservation := range dropDownReservation.ItemDropDownReservations {
			if itemDropDownReservation.TourItem.ID == reservation.TourItemID {
				dropDownReservations[i].ItemDropDownReservations[j].Reservations = append(itemDropDownReservation.Reservations, reservation)
				dropDownReservations[i].ItemDropDownReservations[j].TotalAdults += reservation.Adults
				dropDownReservations[i].ItemDropDownReservations[j].TotalChildren += reservation.Children
				dropDownReservations[i].ItemDropDownReservations[j].TotalPrice += reservation.TotalPrice()
				dropDownReservations[i].ItemDropDownReservations[j].TotalPaidUSD += reservation.GetPaidUSD()
				dropDownReservations[i].ItemDropDownReservations[j].TotalPaidVND += reservation.GetPaidVND()
				dropDownReservations[i].TotalAdults += reservation.Adults
				dropDownReservations[i].TotalChildren += reservation.Children
				break
			}
		}
	}
}

func getTour(tours []model.Tour, tourItem model.TourItem) model.Tour {
	for _, tour := range tours {
		if tour.ID == tourItem.TourID {
			return tour
		}
	}

	return model.Tour{}
}
