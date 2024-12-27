package ssr

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/leedrum/ikarus_travel/internal"
	"github.com/leedrum/ikarus_travel/locales"
	"github.com/leedrum/ikarus_travel/model"
	services "github.com/leedrum/ikarus_travel/services/import"
	"github.com/leedrum/ikarus_travel/views"
	"github.com/rs/zerolog/log"
)

func NewHotelHandler(server internal.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		internal.Render(ctx, http.StatusOK, views.NewHotel())
	}
}

func CreateHotelHandler(server internal.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
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
				locales.Translate(ctx, "success")+"! "+locales.Translate(ctx, "go_to_detail"),
			),
		)
	}
}

func ListHotelsHandler(server internal.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var hotels []model.Hotel
		server.DB.Find(&hotels)

		internal.Render(ctx, http.StatusOK, views.ListHotels(hotels))
	}
}

func DeleteHotelHandler(server internal.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		var hotel model.Hotel
		server.DB.First(&hotel, id)
		server.DB.Delete(&hotel)

		ctx.Data(http.StatusOK, "text/html", []byte(""))
	}
}

func EditHotelHandler(server internal.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		var hotel model.Hotel
		server.DB.First(&hotel, id)

		internal.Render(ctx, http.StatusOK, views.EditHotel(hotel))
	}
}

func ImportHotelsHandler(server internal.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		file, err := ctx.FormFile("file")
		if err != nil {
			log.Error().Err(err).Msg(locales.Translate(ctx, "errors.import.common.upload_file_failed"))
			internal.Render(ctx, http.StatusBadRequest, views.Error(locales.Translate(ctx, "errors.import.common.upload_file_failed")))
			return
		}

		if file.Header.Get("Content-Type") != "text/csv" {
			log.Error().Msg(locales.Translate(ctx, "errors.import.common.invalid_file_format"))
			internal.Render(ctx, http.StatusBadRequest, views.Error(locales.Translate(ctx, "errors.import.common.invalid_file_format")))
			return
		}

		if file.Size == 0 {
			log.Error().Msg(locales.Translate(ctx, "errors.import.common.empty_file"))
			internal.Render(ctx, http.StatusBadRequest, views.Error(locales.Translate(ctx, "errors.import.common.empty_file")))
			return
		}

		if file.Size > services.HotelImportLimitFileSize {
			log.Error().Msg(locales.Translate(ctx, "errors.import.common.file_too_large"))
			internal.Render(ctx, http.StatusBadRequest, views.Error(locales.Translate(ctx, "errors.import.common.file_too_large")))
			return
		}

		err = services.ImportHotels(server.DB, file)
		if err != nil {
			log.Error().Err(err).Msg(locales.Translate(ctx, "errors.import.common.upload_file_failed"))
			internal.Render(ctx, http.StatusBadRequest, views.Error(locales.Translate(ctx, "errors.import.common.importing_file")))
			return
		}
		internal.Render(ctx, http.StatusOK, views.Success(locales.Translate(ctx, "success")))
	}
}
