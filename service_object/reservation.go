package service_object

import (
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/leedrum/ikarus_travel/internal"
	"github.com/leedrum/ikarus_travel/model"
	"github.com/rs/zerolog/log"
	"github.com/xuri/excelize/v2"
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

func ExportExcelFile(ctx *gin.Context, server internal.Server) (*excelize.File, error) {
	dropDownReservations := LoadDropDownReservations(ctx, server)
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			log.Err(err).Msg("Failed to close file")
		}
	}()

	for _, dropDownReservation := range dropDownReservations {
		nextRow := 0
		if len(dropDownReservation.ItemDropDownReservations) == 0 {
			continue
		}
		sheetName := strings.ReplaceAll(dropDownReservation.DepartureDate, "/", "-")
		index, err := f.NewSheet(sheetName)
		if err != nil {
			log.Err(err).Msg("Failed to create new sheet")
			return &excelize.File{}, err
		}

		// Header
		f.SetCellValue(sheetName, "A1", strings.ReplaceAll(sheetName, "/", "-"))
		f.MergeCell(sheetName, "A1", "A2")

		f.SetCellValue(sheetName, "B1", "Hotel")
		f.MergeCell(sheetName, "B1", "B2")

		f.SetCellValue(sheetName, "C1", "Room note")
		f.MergeCell(sheetName, "C1", "C2")

		f.SetCellValue(sheetName, "D1", "Name")
		f.MergeCell(sheetName, "D1", "D2")

		f.SetCellValue(sheetName, "E1", "Guest")
		f.MergeCell(sheetName, "E1", "F1")
		f.SetCellValue(sheetName, "E2", "Adults")
		f.SetCellValue(sheetName, "F2", "Children")

		f.SetCellValue(sheetName, "G1", "Tel")
		f.MergeCell(sheetName, "G1", "G2")

		f.SetCellValue(sheetName, "H1", "Pickup")
		f.MergeCell(sheetName, "H1", "H2")

		f.SetCellValue(sheetName, "I1", "Total")
		f.MergeCell(sheetName, "I1", "I2")

		f.SetCellValue(sheetName, "J1", "Payment")
		f.MergeCell(sheetName, "J1", "K1")
		f.SetCellValue(sheetName, "J2", "Deposit (USD)")

		f.SetCellValue(sheetName, "K2", "Balance")

		f.SetCellValue(sheetName, "L1", "Change")
		f.SetCellValue(sheetName, "L2", "Paid VND")

		f.SetCellValue(sheetName, "M1", "Address")
		f.MergeCell(sheetName, "M1", "M2")

		f.SetCellValue(sheetName, "N1", "Sale")
		f.MergeCell(sheetName, "N1", "P1")
		f.SetCellValue(sheetName, "N2", "Name")
		f.SetCellValue(sheetName, "O2", "No.Ticket")
		f.SetCellValue(sheetName, "P2", "Date")
		// end header
		nextRow += 2

		for _, itemDropDownReservation := range dropDownReservation.ItemDropDownReservations {
			nextRow++
			f.SetCellValue(sheetName, "A"+strconv.Itoa(nextRow), itemDropDownReservation.Tour.Name)
			styleTourName, err := f.NewStyle(&excelize.Style{
				Font: &excelize.Font{
					Bold: true,
					Size: 15,
				},
				Alignment: &excelize.Alignment{
					Horizontal: "center",
					Vertical:   "center",
				},
			})
			if err != nil {
				log.Err(err).Msg("Failed to create new style")
				return &excelize.File{}, err
			}
			f.SetCellStyle(sheetName, "A"+strconv.Itoa(nextRow), "A"+strconv.Itoa(nextRow), styleTourName)
			f.MergeCell(sheetName, "A"+strconv.Itoa(nextRow), "P"+strconv.Itoa(nextRow))

			for k, reservation := range itemDropDownReservation.Reservations {
				f.SetCellValue(sheetName, "A"+strconv.Itoa(nextRow+1), k+1)
				f.SetCellValue(sheetName, "B"+strconv.Itoa(nextRow+1), reservation.Hotel.Name)
				f.SetCellValue(sheetName, "C"+strconv.Itoa(nextRow+1), reservation.RoomNote)
				f.SetCellValue(sheetName, "D"+strconv.Itoa(nextRow+1), reservation.CustomerName)
				f.SetCellValue(sheetName, "E"+strconv.Itoa(nextRow+1), reservation.Adults)
				f.SetCellValue(sheetName, "F"+strconv.Itoa(nextRow+1), reservation.Children)
				f.SetCellValue(sheetName, "G"+strconv.Itoa(nextRow+1), reservation.Phone)
				f.SetCellValue(sheetName, "H"+strconv.Itoa(nextRow+1), reservation.PickupTime)
				f.SetCellValue(sheetName, "I"+strconv.Itoa(nextRow+1), reservation.TotalPrice())
				f.SetCellValue(sheetName, "J"+strconv.Itoa(nextRow+1), reservation.GetPaidUSD())
				f.SetCellValue(sheetName, "K"+strconv.Itoa(nextRow+1), "")
				f.SetCellValue(sheetName, "L"+strconv.Itoa(nextRow+1), reservation.GetPaidVND())
				f.SetCellValue(sheetName, "M"+strconv.Itoa(nextRow+1), reservation.Address)
				f.SetCellValue(sheetName, "N"+strconv.Itoa(nextRow+1), reservation.User.FullName)
				f.SetCellValue(sheetName, "O"+strconv.Itoa(nextRow+1), reservation.ID)
				f.SetCellValue(sheetName, "P"+strconv.Itoa(nextRow+1), reservation.CreatedAt.Format("02/01/2006"))

				// set total USD and VND
				f.SetCellValue(sheetName, "I"+strconv.Itoa(nextRow+2), itemDropDownReservation.TotalPrice)
				f.SetCellValue(sheetName, "J"+strconv.Itoa(nextRow+2), itemDropDownReservation.TotalPaidUSD)
				f.SetCellValue(sheetName, "L"+strconv.Itoa(nextRow+2), itemDropDownReservation.TotalPaidVND)

				nextRow += 2
			}
			nextRow++
		}

		f.SetActiveSheet(index)
	}
	f.DeleteSheet("Sheet1") // delete default sheet

	return f, nil
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
