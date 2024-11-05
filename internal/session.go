package internal

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
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
