package ssr

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/leedrum/ikarus_travel/internal"
	"github.com/leedrum/ikarus_travel/locales"
	"github.com/leedrum/ikarus_travel/model"
	"github.com/leedrum/ikarus_travel/views"
	"github.com/rs/zerolog/log"
)

func NewHotelHandler(server internal.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		internal.Render(ctx, http.StatusOK, views.NewHotel())
	}
}

func CreateHotelHandler(server internal.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		var hotel model.Hotel
		err := ctx.ShouldBind(&hotel)
		if err != nil {
			log.Fatal().Err(err).Msg("Error binding hotel data")
			internal.Render(ctx, http.StatusBadRequest, views.NewHotel())
			return
		}
		server.DB.Create(&hotel)
		internal.Render(ctx, http.StatusOK,
			views.SuccessWithLink(
				"/admin/hotels/"+strconv.Itoa(hotel.ID)+"/edit",
				locales.Translate(ctx, "success")+locales.Translate(ctx, "go_to_detail"),
			),
		)
	}
}

func ListHotelsHandler(server internal.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		var hotels []model.Hotel
		server.DB.Find(&hotels)

		internal.Render(ctx, http.StatusOK, views.ListHotels(hotels))
	}
}

func DeleteHotelHandler(server internal.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		id := ctx.Param("id")
		var hotel model.Hotel
		server.DB.First(&hotel, id)
		server.DB.Delete(&hotel)

		ctx.Data(http.StatusOK, "text/html", []byte(""))
	}
}

func EditHotelHandler(server internal.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		id := ctx.Param("id")
		var hotel model.Hotel
		server.DB.First(&hotel, id)

		internal.Render(ctx, http.StatusOK, views.EditHotel(hotel))
	}
}
