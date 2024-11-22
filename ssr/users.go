package ssr

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leedrum/ikarus_travel/internal"
	"github.com/leedrum/ikarus_travel/locales"
	"github.com/leedrum/ikarus_travel/model"
	"github.com/leedrum/ikarus_travel/views"
)

func NewUserHandler(_ internal.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		internal.Render(ctx, http.StatusOK, views.NewUser(model.User{}))
	}
}

func CreateUserHandler(server internal.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := model.User{}
		err := ctx.ShouldBind(&user)
		if err != nil {
			log.Fatal(err)
			internal.Render(ctx, http.StatusBadRequest, views.NewUser(user))
			return
		}

		password := ctx.PostForm("password")
		if password == "" {
			internal.Render(
				ctx,
				http.StatusBadRequest,
				views.Error(locales.Translate(ctx, "errors.password_cannot_be_empty")),
			)
			return
		}

		user.HashPassword, err = internal.HashPassword(password)
		if err != nil {
			internal.Render(
				ctx,
				http.StatusBadRequest,
				views.Error(locales.Translate(ctx, "errors.password_hash_error")),
			)
			return
		}

		result := server.DB.Create(&user)
		if result.Error != nil {
			internal.Render(ctx, http.StatusBadRequest, views.NewUser(user))
			return
		}

		internal.Render(
			ctx,
			http.StatusOK,
			views.SuccessWithLink(
				"/admin/users",
				"User created",
			),
		)
	}
}

func ListUsersHandler(server internal.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var users []model.User
		server.DB.Find(&users)

		internal.Render(ctx, http.StatusOK, views.ListUsers(users))
	}
}

func DeleteUserHandler(server internal.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		var user model.User
		server.DB.First(&user, id)
		server.DB.Delete(&user)

		ctx.Data(http.StatusOK, "text/html", []byte(""))
	}
}

func EditUserHandler(server internal.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		var user model.User
		server.DB.First(&user, id)

		internal.Render(ctx, http.StatusOK, views.EditUser(user))
	}
}

func UpdateUserHandler(server internal.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		var user model.User
		server.DB.First(&user, id)

		err := ctx.ShouldBind(&user)
		if err != nil {
			log.Fatal(err)
			internal.Render(ctx, http.StatusBadRequest, views.EditUser(user))
			return
		}

		password := ctx.PostForm("password")
		if password != "" {
			user.HashPassword, err = internal.HashPassword(password)
			if err != nil {
				internal.Render(
					ctx,
					http.StatusBadRequest,
					views.Error(locales.Translate(ctx, "errors.password_hash_error")),
				)
				return
			}
		}

		server.DB.Save(&user)

		internal.Render(
			ctx,
			http.StatusOK,
			views.SuccessWithLink(
				"/admin/users",
				"User updated",
			),
		)
	}
}
