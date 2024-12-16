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

func ListResPaymentHandler(server internal.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resID := ctx.Param("res_id")
		reservationID, err := strconv.Atoi(resID)
		if err != nil {
			internal.Render(ctx, http.StatusBadRequest, views.Error("Invalid reservation ID"))
			return
		}
		payments := model.GetPaymentHistory([]int{reservationID}, server.DB)
		internal.Render(ctx, http.StatusOK, views.ListPayments(payments))
	}
}

func CreateResPaymentHandler(server internal.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resID := ctx.Param("res_id")
		reservationID, err := strconv.Atoi(resID)
		if err != nil {
			internal.Render(ctx, http.StatusBadRequest, views.Error("Invalid reservation ID"))
			return
		}
		payment := model.Payment{}
		if err := ctx.ShouldBind(&payment); err != nil {
			internal.Render(ctx, http.StatusBadRequest, views.Error("Invalid payment data"))
			return
		}
		payment.ReservationID = reservationID
		payment.UserID = ctx.MustGet("user").(model.User).ID
		if err := server.DB.Create(&payment).Error; err != nil {
			internal.Render(ctx, http.StatusBadRequest, views.Error("Error creating payment"))
			return
		}
		payments := model.GetPaymentHistory([]int{reservationID}, server.DB)
		internal.Render(ctx, http.StatusOK, views.ListPayments(payments))
	}
}

func DeleteResPaymentHandler(server internal.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resID := ctx.Param("res_id")
		reservationID, err := strconv.Atoi(resID)
		if err != nil {
			internal.Render(ctx, http.StatusBadRequest, views.Error("Invalid reservation ID"))
			return
		}
		paymentID := ctx.Param("id")
		id, err := strconv.Atoi(paymentID)
		if err != nil {
			internal.Render(ctx, http.StatusBadRequest, views.Error("Invalid payment ID"))
			return
		}
		payment := model.Payment{}
		if err := server.DB.Preload("User").Where("id = ? AND reservation_id = ?", id, reservationID).First(&payment).Error; err != nil {
			internal.Render(ctx, http.StatusBadRequest, views.Error(locales.Translate(ctx, "errors.not_found")))
			return
		}
		currentUser := ctx.MustGet("user").(model.User)
		if payment.UserID != currentUser.ID || currentUser.Role != model.RoleAdmin {
			internal.Render(ctx, http.StatusForbidden, views.Error(locales.Translate(ctx, "errors.no_permission")))
			return
		}

		if err := server.DB.Where("id = ? AND reservation_id = ?", id, reservationID).Delete(&model.Payment{}).Error; err != nil {
			log.Fatal().Err(err).Msg(err.Error())
			internal.Render(ctx, http.StatusBadRequest, views.Error(locales.Translate(ctx, "errors.something_went_wrong")))
			return
		}
		ctx.Data(http.StatusOK, "text/html", []byte(""))
	}
}
