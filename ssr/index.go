package ssr

import (
	"context"
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"github.com/leedrum/ikarus_travel/views"
)

const appTimeout = time.Second * 10

func render(ctx *gin.Context, status int, template templ.Component) error {
	ctx.Status(status)
	return template.Render(ctx, ctx.Writer)
}

func IndexHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		defer cancel()
		render(ctx, http.StatusOK, views.Index())
	}
}
