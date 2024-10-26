package middlewares

import (
	"net/http"
	"time"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
	"github.com/leedrum/ikarus_travel/locales"
)

func TimeoutMiddleware() gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(500*time.Millisecond),
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		timeout.WithResponse(timeoutResponse),
	)
}

func timeoutResponse(c *gin.Context) {
	c.JSON(http.StatusRequestTimeout, gin.H{"error": locales.Translate(c, "errors.request_timeout")})
}
