package middlewares

import (
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/leedrum/ikarus_travel/internal"
	"github.com/leedrum/ikarus_travel/model"
	"github.com/leedrum/ikarus_travel/views"
	"github.com/rs/zerolog/log"
)

func AuthRequired(server internal.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("userID")
		if userID == nil {
			log.Info().Msg("userID is nil")
			c.Redirect(http.StatusTemporaryRedirect, "/login")
			c.Abort()
			return
		} else {
			user := model.User{}
			server.DB.Where("id = ?", userID).First(&user)
			if user.ID == 0 {
				log.Info().Msg("user.ID is 0")
				c.Redirect(http.StatusTemporaryRedirect, "/login")
				c.Abort()
				return
			} else {
				c.Set("user", user)
			}
		}

		expired_time := session.Get("userID_expired_time")
		if expired_time == nil {
			log.Info().Msg("expired_time is nil")
			c.Redirect(http.StatusTemporaryRedirect, "/login")
			c.Abort()
			return
		} else if expired_time.(int64) < (time.Now().Unix() - 84600) {
			session.Clear()
			session.Save()
			log.Info().Msg("eexpired_time is less than 84600")
			c.Redirect(http.StatusTemporaryRedirect, "/login")
			c.Abort()
			return
		}

		c.Set("userID", session.Get("userID"))
		c.Next()
	}
}

func AdminRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("user").(model.User)
		if user.Role != "admin" {
			views.NoPermission().Render(c, c.Writer)
			c.Abort()
			return
		}

		c.Next()
	}
}
