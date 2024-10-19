package ssr

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/leedrum/ikarus_travel/internal"
	"github.com/leedrum/ikarus_travel/locales"
	"github.com/leedrum/ikarus_travel/model"
	"github.com/leedrum/ikarus_travel/views"
)

func NewUserHandler(_ internal.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		internal.Render(ctx, http.StatusOK, views.NewUser(model.User{}))
	}
}

func CreateUserHandler(server internal.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		user := model.User{}
		err := ctx.ShouldBind(&user)
		if err != nil {
			log.Fatal(err)
			internal.Render(ctx, http.StatusBadRequest, views.NewUser(user))
			return
		}

		password := ctx.PostForm("password")
		if password == "" {
			internal.Render(ctx, http.StatusBadRequest, views.Error(locales.Translate(ctx, "password_cannot_be_empty")))
			return
		}

		// TODO: Hash password

		result := server.DB.Create(&user)
		if result.Error != nil {
			internal.Render(ctx, http.StatusBadRequest, views.NewUser(user))
			return
		}

		internal.Render(ctx, http.StatusBadRequest, views.CreateSuccessUser(user))
	}
}

func ListUsersHandler(server internal.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		var users []model.User
		server.DB.Find(&users)

		internal.Render(ctx, http.StatusOK, views.ListUsers(users))
	}
}
