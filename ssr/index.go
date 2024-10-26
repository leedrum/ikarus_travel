package ssr

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leedrum/ikarus_travel/internal"
	"github.com/leedrum/ikarus_travel/views"
)

func IndexHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		internal.Render(ctx, http.StatusOK, views.AdminDashboard())
	}
}

func AdminIndexHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		internal.Render(ctx, http.StatusOK, views.AdminDashboard())
	}
}
