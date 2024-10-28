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
	"gorm.io/gorm"
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
		var reservation model.Reservation
		if err := ctx.ShouldBind(&reservation); err != nil {
			log.Error().Err(err).Msg("Error binding data")
			hotels := getHotels(server)
			tours := getTours(server)
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
		hotels := getHotels(server)
		tours := getTours(server)
		users := getUers(server)

		reservations = mappingData(reservations, users, hotels, tours)
		internal.Render(ctx, http.StatusOK, views.ListReservations(reservations))
	}
}

type MineReservationsResponse struct {
	Reservations  []model.Reservation
	TotalAdults   int
	TotalChildren int
	Tour          model.Tour
	Hotel         model.Hotel
	DepartureDate string
}

func MineReservationsHandler(server internal.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var reservations []model.Reservation
		user := ctx.MustGet("user").(model.User)
		tx := server.DB.Where("user_id = ?", user.ID)
		tx = searchConditions(ctx, tx)
		tx.Find(&reservations)
		hotels := getHotels(server)
		tours := getTours(server)
		users := getUers(server)

		reservations = mappingData(reservations, users, hotels, tours)
		internal.Render(ctx, http.StatusOK, views.ListGroupReservations())
	}
}

func EditReservationHandler(server internal.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		hotels := getHotels(server)
		tours := getTours(server)

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
		hotels := getHotels(server)
		tours := getTours(server)

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
	server.DB.Table("users").Select("id, username").Find(&users)
	return users
}

func mappingData(reservations []model.Reservation, users []model.User, hotels []model.Hotel, tours []model.Tour) []model.Reservation {
	for i := range reservations {
		reservations[i] = mappingUser(reservations[i], users)
		reservations[i] = mappingHotel(reservations[i], hotels)
		reservations[i] = mappingTour(reservations[i], tours)
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

func mappingTour(reservation model.Reservation, tours []model.Tour) model.Reservation {
	for i := range tours {
		if reservation.TourID == tours[i].ID {
			reservation.Tour = tours[i]
		}
	}
	return reservation
}

func searchConditions(ctx *gin.Context, tx *gorm.DB) *gorm.DB {
	if ctx.Query("start_date") != "" {
		tx = tx.Where("start_date >= ?", ctx.Query("start_date"))
	}
	if ctx.Query("departure_date") != "" {
		tx = tx.Where("departure_date <= ?", ctx.Query("departure_date"))
	}
	return tx
}
