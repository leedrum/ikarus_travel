package ssr

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func InitSsrRoutes(router *gin.Engine) {
	fmt.Println("SSR routes initialized")
	router.GET("/", IndexHandler())
}
