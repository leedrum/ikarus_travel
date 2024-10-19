package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/leedrum/ikarus_travel/internal"
	"github.com/leedrum/ikarus_travel/middlewares"
	"github.com/leedrum/ikarus_travel/ssr"
)

const AppTimeout = time.Second * 10

func InitRoutes(server *internal.Server) {
	router := gin.Default()
	router.Use(middlewares.I18nMiddleware())
	router.GET("/", ssr.IndexHandler())
	admin := router.Group("/admin")
	{
		admin.GET("/users/new", ssr.NewUserHandler(*server))
		admin.POST("/users", ssr.CreateUserHandler(*server))
	}
	server.Router = router
}
