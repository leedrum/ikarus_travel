package middlewares

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func TimeoutMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		timeout := 1 * time.Second
		// Create a context with a timeout
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()

		// Replace the request's context with the new timeout context
		c.Request = c.Request.WithContext(ctx)

		// Channel to signal when the request is done
		done := make(chan struct{})

		go func() {
			// Process the request
			c.Next()
			close(done)
		}()

		// Wait for the request to complete or timeout
		select {
		case <-ctx.Done():
			// If the context times out, respond with 504 status code
			c.AbortWithError(http.StatusGatewayTimeout, errors.New("request timed out"))
		case <-done:
			// If the request completes, proceed as usual
		}
	}
}
