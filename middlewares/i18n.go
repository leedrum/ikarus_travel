package middlewares

import (
	"github.com/gin-gonic/gin"
)

func I18nMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		lang, exist := c.Get("Accept-Language")
		if !exist {
			lang = "en"
		}
		c.Set("lang", lang)
		c.Next()
	}
}
