package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/leedrum/ikarus_travel/api"
	"github.com/leedrum/ikarus_travel/ssr"
)

func InitRoutes() *gin.Engine {
	router := gin.Default()
	api.InitAPIRoutes(router)
	ssr.InitSsrRoutes(router)

	return router
}
