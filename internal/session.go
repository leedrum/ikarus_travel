package internal

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/leedrum/ikarus_travel/model"
)

func (server *Server) InitSessionStore() {
	store := cookie.NewStore([]byte(server.Config.SessionName))
	secureSession := false
	if server.Config.Environment == "production" {
		secureSession = true
	}
	store.Options(sessions.Options{
		HttpOnly: true,
		Secure:   secureSession, // Set to true if your website is served over HTTPS
		SameSite: http.SameSiteLaxMode,
		MaxAge:   int(server.Config.SessionsExpiration),
	})

	server.SessionStore = &store
}

func UserFromContext(ctx *gin.Context) model.User {
	return ctx.MustGet("user").(model.User)
}
