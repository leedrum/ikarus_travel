package ssr

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/leedrum/ikarus_travel/internal"
	"github.com/leedrum/ikarus_travel/locales"
	"github.com/leedrum/ikarus_travel/model"
	"github.com/leedrum/ikarus_travel/service_object"
	"github.com/leedrum/ikarus_travel/views"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm/clause"
)

func NewReservationHandler(server internal.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		hotels := getHotels(server)
		tours := getTours(server)
		internal.Render(ctx, http.StatusOK, views.NewReservation(hotels, tours))
	}
}

func CreateReservationHandler(server internal.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tourItem := model.TourItem{}
		server.DB.Where("departure_date = ?", ctx.PostForm("departure_date")).First(&tourItem)

		if tourItem.ID == 0 {

			if err := ctx.ShouldBind(&tourItem); err != nil {
				log.Error().Err(err).Msg("Error binding data")
				hotels := getHotels(server)
				tours := getTours(server)
				internal.Render(ctx, http.StatusBadRequest, views.NewReservation(hotels, tours))
			}

			response := server.DB.Create(&tourItem)
			if tourItem.ID == 0 {
				log.Fatal().Err(response.Error).Msg(response.Error.Error())
				internal.Render(ctx, http.StatusBadRequest, views.Error(locales.Translate(ctx, "something_went_wrong")))
				return
			}
		}

		var reservation model.Reservation
		reservation.TourItemID = tourItem.ID
		if err := ctx.ShouldBind(&reservation); err != nil {
			log.Error().Err(err).Msg("Error binding data")
			hotels := getHotels(server)
			tours := getTours(server)
			internal.Render(ctx, http.StatusBadRequest, views.NewReservation(hotels, tours))
		}
		reservation.TourItemID = tourItem.ID
		reservation.UserID = ctx.MustGet("user").(model.User).ID

		result := server.DB.Omit(clause.Associations).Create(&reservation)
		if result.Error != nil {
			internal.Render(ctx, http.StatusBadRequest, views.Error(result.Error.Error()))
			return
		}

		infoRes := locales.Translate(ctx, "success") + "! " + locales.Translate(ctx, "go_to_detail")
		internal.Render(ctx, http.StatusOK,
			views.SuccessWithLink(
				"/admin/reservations/"+strconv.Itoa(reservation.ID)+"/edit",
				infoRes,
			),
		)
	}
}

func ListReservationsHandler(server internal.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var reservations []model.Reservation
		tx := service_object.SearchConditions(ctx, server.DB)
		tx.Preload("Hotel").Preload("TourItem").Preload("User").Find(&reservations)
		internal.Render(ctx, http.StatusOK, views.ListReservations(reservations))
	}
}

func MineReservationsHandler(server internal.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		dropDownReservations := service_object.LoadDropDownReservations(ctx, server)
		internal.Render(ctx, http.StatusOK, views.ListGroupReservations(dropDownReservations))
	}
}

func EditReservationHandler(server internal.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		hotels := getHotels(server)
		tours := getTours(server)

		if err != nil {
			internal.Render(ctx, http.StatusBadRequest, views.Error("Invalid reservation ID"))
			return
		}

		var reservation model.Reservation
		server.DB.Preload("TourItem").First(&reservation, id)
		internal.Render(ctx, http.StatusOK, views.EditReservation(reservation, hotels, tours))
	}
}

func DeleteReservationHandler(server internal.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		var reservation model.Reservation
		server.DB.First(&reservation, id)
		server.DB.Delete(&reservation)

		ctx.Data(http.StatusOK, "text/html", []byte(""))
	}
}

func UpdateReservationHandler(server internal.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		hotels := getHotels(server)
		tours := getTours(server)

		if err != nil {
			internal.Render(ctx, http.StatusBadRequest, views.Error("Invalid reservation ID"))
			return
		}
		var reservation model.Reservation
		server.DB.Where("id = ?", id).First(&reservation)

		previousTourItem := model.TourItem{}
		server.DB.Where("id = ?", reservation.TourItemID).First(&previousTourItem)

		tourItem := model.TourItem{}
		server.DB.Where("tour_id = ?", ctx.PostForm("tour_id")).Where("departure_date = ?", ctx.PostForm("departure_date")).First(&tourItem)

		if tourItem.ID == 0 {
			if err := ctx.ShouldBind(&tourItem); err != nil {
				log.Error().Err(err).Msg("Error binding data")
				hotels := getHotels(server)
				tours := getTours(server)
				internal.Render(ctx, http.StatusBadRequest, views.NewReservation(hotels, tours))
			}

			tourItem.TourID = previousTourItem.TourID
			if ctx.PostForm("tour_id") != "" {
				tourItem.TourID, _ = strconv.Atoi(ctx.PostForm("tour_id"))
			}
			response := server.DB.Omit(clause.Associations).Create(&tourItem)
			if tourItem.ID == 0 {
				log.Fatal().Err(response.Error).Msg(response.Error.Error())
				internal.Render(ctx, http.StatusBadRequest, views.Error(locales.Translate(ctx, "something_went_wrong")))
				return
			}
		}

		err = ctx.ShouldBind(&reservation)
		if err != nil {
			internal.Render(ctx, http.StatusBadRequest, views.EditReservation(reservation, hotels, tours))
			return
		}
		reservation.TourItemID = tourItem.ID
		server.DB.Omit(clause.Associations).Save(&reservation)

		internal.Render(ctx, http.StatusOK,
			views.SuccessWithLink(
				"/admin/reservations/"+strconv.Itoa(reservation.ID)+"/edit",
				locales.Translate(ctx, "success")+"! "+locales.Translate(ctx, "go_to_detail"),
			),
		)
	}
}

func GenerateQRCodeHandler(server internal.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			internal.Render(ctx, http.StatusBadRequest, views.Error("Invalid reservation ID"))
			return
		}

		var reservation model.Reservation
		server.DB.First(&reservation, id)

		domain := server.Config.DomainAddress
		qrCode := internal.IkarusQRCode{
			Content: "https://" + domain + "/reservations/preview/" + reservation.Code,
			Size:    256,
		}

		qrCodeData, err := qrCode.Generate()
		if err != nil {
			msg_str := locales.Translate(ctx, "errors.something_went_wrong")
			message := fmt.Sprintf("%s: %v", msg_str, err)
			log.Error().Err(err).Msg(message)
			internal.Render(ctx, http.StatusInternalServerError, views.Error(msg_str))
			return
		}
		imgBase64 := base64.StdEncoding.EncodeToString(qrCodeData)
		internal.Render(ctx, http.StatusOK, views.QRCode("data:image/png;base64, "+imgBase64))
	}
}

func PreviewQRCodeHandler(server internal.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		code := ctx.Param("code")
		var reservation model.Reservation
		server.DB.Preload("Payments").Preload("Hotel").Preload("TourItem").Where(
			"code = ?", code).First(&reservation)
		server.DB.Model(&reservation.TourItem).Association("Tour").Find(&reservation.TourItem.Tour)
		internal.Render(ctx, http.StatusOK, views.PreviewReservation(reservation))
	}
}

func ExportReservationsHandler(server internal.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fileExcel, err := service_object.ExportExcelFile(ctx, server)
		if err != nil {
			internal.Render(ctx, http.StatusInternalServerError, views.Error("Error exporting file"))
			return
		} else {
			ctx.Header("Content-Description", "File Transfer")
			ctx.Header("Content-Disposition", "attachment; filename=report.xlsx")
			fileExcel.Write(ctx.Writer)
			ctx.Data(http.StatusOK, "application/octet-stream", []byte(""))
		}
	}
}

func getHotels(server internal.Server) []model.Hotel {
	hotels := []model.Hotel{}
	server.DB.Table("hotels").Select("id, name").Find(&hotels)
	return hotels
}

func getTours(server internal.Server) []model.Tour {
	tours := []model.Tour{}
	server.DB.Table("tours").Select("id, name").Find(&tours)
	return tours
}
