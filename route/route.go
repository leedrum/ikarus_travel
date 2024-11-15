package route

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/leedrum/ikarus_travel/internal"
	"github.com/leedrum/ikarus_travel/middlewares"
	"github.com/leedrum/ikarus_travel/ssr"
)

func InitRoutes(server *internal.Server) {
	router := gin.Default()
	router.Use(middlewares.CORSMiddleware())
	router.Use(middlewares.TimeoutMiddleware())
	router.Use(middlewares.I18nMiddleware())
	router.Use(sessions.Sessions(server.Config.SessionName, *server.SessionStore))

	router.GET("/login", ssr.LoginFormHandler(*server))
	router.POST("/login", ssr.LoginHandler(*server))
	router.GET("/logout", ssr.LogoutHandler)
	router.Use(middlewares.AuthRequired(*server))
	router.GET("/", ssr.IndexHandler())

	routerResUser := router.Group("/reservations")
	{
		routerResUser.GET("/mine", ssr.MineReservationsHandler(*server))
		routerResUser.GET("/new", ssr.NewReservationHandler(*server))
		routerResPayment := routerResUser.Group("/:res_id/payments")
		{
			routerResPayment.GET("/", ssr.ListResPaymentHandler(*server))
			routerResPayment.POST("/", ssr.CreateResPaymentHandler(*server))
			routerResPayment.DELETE("/:id", ssr.DeleteResPaymentHandler(*server))
		}
	}

	adminRoute := router.Group("/admin")
	{
		adminRoute.Use(sessions.Sessions(server.Config.SessionName, *server.SessionStore))
		adminRoute.Use(middlewares.AuthRequired(*server))
		adminRoute.Use(middlewares.AdminRequired())
		adminRoute.GET("/", ssr.AdminIndexHandler())
		adminUserRoutes := adminRoute.Group("/users")
		{
			adminUserRoutes.GET("/", ssr.ListUsersHandler(*server))
			adminUserRoutes.GET("/new", ssr.NewUserHandler(*server))
			adminUserRoutes.POST("/", ssr.CreateUserHandler(*server))
			adminUserRoutes.DELETE("/:id", ssr.DeleteUserHandler(*server))
		}

		tourRoutes := adminRoute.Group("/tours")
		{
			tourRoutes.GET("/", ssr.ListToursHandler(*server))
			tourRoutes.GET("/new", ssr.NewTourHandler(*server))
			tourRoutes.POST("/", ssr.CreateTourHandler(*server))
			tourRoutes.GET("/:id/edit", ssr.EditTourHandler(*server))
			tourRoutes.PUT("/:id", ssr.UpdateTourHandler(*server))
			tourRoutes.DELETE("/:id", ssr.DeleteTourHandler(*server))
		}

		hotelRoutes := adminRoute.Group("/hotels")
		{
			hotelRoutes.GET("/", ssr.ListHotelsHandler(*server))
			hotelRoutes.GET("/new", ssr.NewHotelHandler(*server))
			hotelRoutes.POST("/", ssr.CreateHotelHandler(*server))
			hotelRoutes.GET("/:id/edit", ssr.EditHotelHandler(*server))
			hotelRoutes.DELETE("/:id", ssr.DeleteHotelHandler(*server))
		}

		resRoutes := adminRoute.Group("/reservations")
		{
			resRoutes.GET("/", ssr.ListReservationsHandler(*server))
			resRoutes.POST("/", ssr.CreateReservationHandler(*server))
			resRoutes.GET("/:id/edit", ssr.EditReservationHandler(*server))
			resRoutes.PUT("/:id", ssr.UpdateReservationHandler(*server))
			resRoutes.DELETE("/:id", ssr.DeleteReservationHandler(*server))
		}
	}
	server.Router = router
}
