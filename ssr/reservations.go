package ssr

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/leedrum/ikarus_travel/internal"
	"github.com/leedrum/ikarus_travel/locales"
	"github.com/leedrum/ikarus_travel/model"
	"github.com/leedrum/ikarus_travel/views"
	"github.com/rs/zerolog/log"
)

func NewReservationHandler(server internal.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		hotels := []model.Hotel{}
		tours := []model.Tour{}
		server.DB.Table("hotels").Select("id, name").Find(&hotels)
		server.DB.Table("tours").Select("id, name").Find(&tours)
		internal.Render(ctx, http.StatusOK, views.NewReservation(hotels, tours))
	}
}

func CreateReservationHandler(server internal.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var reservation model.Reservation
		if err := ctx.ShouldBind(&reservation); err != nil {
			log.Error().Err(err).Msg("Error binding data")
			hotels := []model.Hotel{}
			tours := []model.Tour{}
			server.DB.Table("hotels").Select("id, name").Find(&hotels)
			server.DB.Table("tours").Select("id, name").Find(&tours)
			internal.Render(ctx, http.StatusBadRequest, views.NewReservation(hotels, tours))
		}

		server.DB.Create(&reservation)

		internal.Render(ctx, http.StatusOK,
			views.SuccessWithLink(
				"/admin/reservations/"+strconv.Itoa(reservation.ID)+"/edit",
				locales.Translate(ctx, "success")+locales.Translate(ctx, "go_to_detail"),
			),
		)
	}
}

func ListReservationsHandler(server internal.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var reservations []model.Reservation
		server.DB.Find(&reservations)
		internal.Render(ctx, http.StatusOK, views.ListReservations(reservations))
	}
}

func EditReservationHandler(server internal.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		hotels := []model.Hotel{}
		tours := []model.Tour{}
		server.DB.Table("hotels").Select("id, name").Find(&hotels)
		server.DB.Table("tours").Select("id, name").Find(&tours)

		if err != nil {
			internal.Render(ctx, http.StatusBadRequest, views.NewReservation(hotels, tours))
			return
		}

		var reservation model.Reservation
		server.DB.First(&reservation, id)
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
		hotels := []model.Hotel{}
		tours := []model.Tour{}
		server.DB.Table("hotels").Select("id, name").Find(&hotels)
		server.DB.Table("tours").Select("id, name").Find(&tours)

		if err != nil {
			internal.Render(ctx, http.StatusBadRequest, views.NewReservation(hotels, tours))
			return
		}

		var reservation model.Reservation
		server.DB.First(&reservation, id)

		err = ctx.ShouldBind(&reservation)
		if err != nil {
			internal.Render(ctx, http.StatusBadRequest, views.EditReservation(reservation, hotels, tours))
			return
		}

		server.DB.Save(&reservation)

		internal.Render(ctx, http.StatusOK,
			views.SuccessWithLink(
				"/admin/reservations/"+strconv.Itoa(reservation.ID)+"/edit",
				locales.Translate(ctx, "success")+locales.Translate(ctx, "go_to_detail"),
			),
		)
	}
}
