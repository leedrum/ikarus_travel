package ssr

import (
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/leedrum/ikarus_travel/internal"
	"github.com/leedrum/ikarus_travel/locales"
	"github.com/leedrum/ikarus_travel/model"
	"github.com/leedrum/ikarus_travel/views"
	"github.com/rs/zerolog/log"
)

func LoginFormHandler(server internal.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		internal.Render(ctx, 200, views.LoginForm())
	}
}

func LogoutHandler(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Writer.Header().Set("HX-Redirect", "/login")
	c.Redirect(http.StatusTemporaryRedirect, "/login")
}

func LoginHandler(server internal.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var loginUser struct {
			Username string `form:"username" binding:"required" json:"username"`
			Password string `form:"password" binding:"required" json:"password"`
		}

		if err := ctx.ShouldBind(&loginUser); err != nil {
			log.Err(err).Msg("Invalid form data")
			internal.Render(ctx, http.StatusBadRequest, views.Error("Invalid form data"))
			return
		}

		user := model.User{}
		server.DB.Where("username = ?", loginUser.Username).First(&user)
		if !internal.CheckPasswordHash(loginUser.Password, user.HashPassword) {
			internal.Render(ctx, http.StatusBadRequest, views.Error(
				locales.Translate(ctx, "errors.invalid_username_or_password")),
			)
			return
		}

		session := sessions.Default(ctx)
		session.Set("userID", user.ID)
		session.Set("userID_expired_time", time.Now().Unix()+server.Config.SessionsExpiration)
		session.Save()
		if user.IsAdmin() {
			ctx.Writer.Header().Set("HX-Redirect", "/admin")
		} else {
			ctx.Writer.Header().Set("HX-Redirect", "/")
		}
	}
}
