package ssr

import (
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

			server.DB.Create(&tourItem)
			if tourItem.ID == 0 {
				internal.Render(ctx, http.StatusBadRequest, views.Error("Error creating tour item"))
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
		reservation.UserID = ctx.MustGet("user").(model.User).ID

		result := server.DB.Create(&reservation)
		if result.Error != nil {
			internal.Render(ctx, http.StatusBadRequest, views.Error(result.Error.Error()))
			return
		}

		infoRes := fmt.Sprintf(
			"Reservation created with ID: %d \n code: %s, go to detail",
			reservation.ID, reservation.Code,
		)
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
		server.DB.Find(&reservations)
		hotels := getHotels(server)
		users := getUers(server)

		reservations = mappingData(reservations, users, hotels)
		internal.Render(ctx, http.StatusOK, views.ListReservations(reservations))
	}
}

func MineReservationsHandler(server internal.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var reservations []model.Reservation
		tx := server.DB
		tx.Find(&reservations)
		hotels := getHotels(server)
		users := getUers(server)
		reservations = mappingData(reservations, users, hotels)

		dropDownReservations := service_object.LoadDropDownReservations(ctx, server)
		reservations = mappingData(reservations, users, hotels)
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
		hotels := getHotels(server)
		tours := getTours(server)

		if err != nil {
			internal.Render(ctx, http.StatusBadRequest, views.Error("Invalid reservation ID"))
			return
		}

		tourItem := model.TourItem{}
		server.DB.Where("departure_date = ?", ctx.PostForm("departure_date")).First(&tourItem)

		if tourItem.ID == 0 {

			if err := ctx.ShouldBind(&tourItem); err != nil {
				log.Error().Err(err).Msg("Error binding data")
				hotels := getHotels(server)
				tours := getTours(server)
				internal.Render(ctx, http.StatusBadRequest, views.NewReservation(hotels, tours))
			}

			server.DB.Create(&tourItem)
			if tourItem.ID == 0 {
				internal.Render(ctx, http.StatusBadRequest, views.Error("Error creating tour item"))
				return
			}
		}

		var reservation model.Reservation
		server.DB.First(&reservation, id)

		err = ctx.ShouldBind(&reservation)
		if err != nil {
			internal.Render(ctx, http.StatusBadRequest, views.EditReservation(reservation, hotels, tours))
			return
		}
		reservation.TourItemID = tourItem.ID

		server.DB.Save(&reservation)

		internal.Render(ctx, http.StatusOK,
			views.SuccessWithLink(
				"/admin/reservations/"+strconv.Itoa(reservation.ID)+"/edit",
				locales.Translate(ctx, "success")+locales.Translate(ctx, "go_to_detail"),
			),
		)
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

func getUers(server internal.Server) []model.User {
	users := []model.User{}
	server.DB.Table("users").Select("id, username, full_name").Find(&users)
	return users
}

func mappingData(reservations []model.Reservation, users []model.User, hotels []model.Hotel) []model.Reservation {
	for i := range reservations {
		reservations[i] = mappingUser(reservations[i], users)
		reservations[i] = mappingHotel(reservations[i], hotels)
	}
	return reservations
}

func mappingUser(reservation model.Reservation, users []model.User) model.Reservation {
	for i := range users {
		if reservation.UserID == users[i].ID {
			reservation.User = users[i]
		}
	}
	return reservation
}

func mappingHotel(reservation model.Reservation, hotels []model.Hotel) model.Reservation {
	for i := range hotels {
		if reservation.HotelID == hotels[i].ID {
			reservation.Hotel = hotels[i]
		}
	}
	return reservation
}
