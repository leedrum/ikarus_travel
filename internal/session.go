package internal

import (
	"github.com/gin-contrib/sessions/cookie"
)

func (server *Server) InitSessionStore() {
	store := cookie.NewStore([]byte("secret"))
	server.SessionStore = &store
}
