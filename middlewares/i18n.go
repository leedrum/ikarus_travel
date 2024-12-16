package middlewares

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func I18nMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var lang string

		session := sessions.Default(c)
		langUserSet := session.Get("lang")
		if c.Request.URL.Query().Get("lang") != "" {
			lang = c.Request.URL.Query().Get("lang")
			session.Set("lang", lang)
			session.Save()
		} else if langUserSet != nil {
			lang = langUserSet.(string)
		} else {
			lang = "en"
		}

		c.Set("lang", lang)
		c.Next()
	}
}
