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

func NewTourHandler(_ internal.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		internal.Render(ctx, http.StatusOK, views.NewTour())
	}
}

func CreateTourHandler(server internal.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tour := model.Tour{}
		err := ctx.ShouldBind(&tour)
		if err != nil {
			log.Fatal().Err(err).Msg("Error binding tour data")
			internal.Render(ctx, http.StatusBadRequest, views.NewTour())
			return
		}

		result := server.DB.Create(&tour)
		if result.Error != nil {
			internal.Render(ctx, http.StatusBadRequest, views.NewTour())
			return
		}

		internal.Render(ctx, http.StatusOK,
			views.SuccessWithLink(
				"/admin/tours/"+strconv.Itoa(tour.ID)+"/edit",
				locales.Translate(ctx, "success")+"! "+locales.Translate(ctx, "go_to_detail"),
			),
		)
	}
}

func ListToursHandler(server internal.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tours []model.Tour
		server.DB.Find(&tours)

		internal.Render(ctx, http.StatusOK, views.ListTours(tours))
	}
}

func DeleteTourHandler(server internal.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		var tour model.Tour
		server.DB.First(&tour, id)
		server.DB.Delete(&tour)

		ctx.Data(http.StatusOK, "text/html", []byte(""))
	}
}

func EditTourHandler(server internal.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		var tour model.Tour
		server.DB.First(&tour, id)

		internal.Render(ctx, http.StatusOK, views.EditTour(tour))
	}
}

func UpdateTourHandler(server internal.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		var tour model.Tour
		server.DB.First(&tour, id)

		err := ctx.ShouldBind(&tour)
		if err != nil {
			log.Fatal().Err(err).Msg("Error binding tour data")
			internal.Render(ctx, http.StatusBadRequest, views.EditTour(tour))
			return
		}

		server.DB.Save(&tour)

		internal.Render(ctx, http.StatusOK,
			views.SuccessWithLink(
				"/admin/tours/",
				locales.Translate(ctx, "success")+"! "+locales.Translate(ctx, "go_to_list"),
			),
		)
	}
}

func ImportToursHandler(server internal.Server) gin.HandlerFunc {
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

		if file.Size > services.TourImportLimit {
			log.Error().Msg(locales.Translate(ctx, "errors.import.common.file_too_large"))
			internal.Render(ctx, http.StatusBadRequest, views.Error(locales.Translate(ctx, "errors.import.common.file_too_large")))
			return
		}

		err = services.ImportTours(server.DB, file)
		if err != nil {
			log.Error().Err(err).Msg(locales.Translate(ctx, "errors.import.common.upload_file_failed"))
			internal.Render(ctx, http.StatusBadRequest, views.Error(locales.Translate(ctx, "errors.import.common.importing_file")))
			return
		}
		internal.Render(ctx, http.StatusOK, views.Success(locales.Translate(ctx, "success")))
	}
}
