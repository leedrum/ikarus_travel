package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/leedrum/ikarus_travel/api"
	"github.com/leedrum/ikarus_travel/middlewares"
	"github.com/leedrum/ikarus_travel/ssr"
	"gorm.io/gorm"
)

type Server struct {
	Router *gin.Engine
	DB     *gorm.DB
	Config Config
}

func (s *Server) initRoutes() {
	router := gin.Default()
	router.Use(middlewares.I18nMiddleware())
	s.Router = router
	api.InitAPIRoutes(s.Router)
	ssr.InitSsrRoutes(s.Router)
}

func (s *Server) Run() {
	InitI18n()
	s.DB = initDB(s.Config)
	s.initRoutes()
	s.Router.Run(s.Config.HTTPServerAddress)
}
