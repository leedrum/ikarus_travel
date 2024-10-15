package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/leedrum/ikarus_travel/api"
	"github.com/leedrum/ikarus_travel/middlewares"
	"github.com/leedrum/ikarus_travel/ssr"
)

func InitRoutes() *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.I18nMiddleware())
	api.InitAPIRoutes(router)
	ssr.InitSsrRoutes(router)

	return router
}
