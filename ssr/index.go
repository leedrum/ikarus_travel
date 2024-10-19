package ssr

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/leedrum/ikarus_travel/internal"
	"github.com/leedrum/ikarus_travel/views"
)

func IndexHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		internal.Render(ctx, http.StatusOK, views.Index())
	}
}
