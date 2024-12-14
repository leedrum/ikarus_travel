package middlewares

import (
	"net/http"

	"strings"

	"github.com/gin-gonic/gin"
	"github.com/leedrum/ikarus_travel/internal"
)

func CORSMiddleware(server internal.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if Gin is not running in production mode
		if gin.Mode() != gin.ReleaseMode {
			// Allow all origins in non-production modes
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
		} else {
			// In production mode, apply stricter rules

			// Allow all origins for GET requests
			if c.Request.Method == http.MethodGet {
				c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
				c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
				c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
			} else {
				// For non-GET requests, check the referer
				referer := c.GetHeader("Referer")
				if !strings.Contains(referer, server.Config.DomainAddress) {
					c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Access denied"})
					return
				}

				// Set CORS headers for the allowed referer
				c.Writer.Header().Set("Access-Control-Allow-Origin", server.Config.DomainAddress)
				c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, PUT, DELETE, OPTIONS")
				c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
			}
		}

		// Handle preflight OPTIONS request
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
