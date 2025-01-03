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
	router.Use(sessions.Sessions(server.Config.SessionName, *server.SessionStore))
	router.Use(middlewares.CORSMiddleware(*server))
	router.Use(middlewares.TimeoutMiddleware())
	router.Use(middlewares.I18nMiddleware())

	router.GET("/login", ssr.LoginFormHandler(*server))
	router.POST("/login", ssr.LoginHandler(*server))
	router.GET("/logout", ssr.LogoutHandler)
	router.GET("reservations/preview/:code", ssr.PreviewQRCodeHandler(*server))
	router.Use(middlewares.AuthRequired(*server))
	router.GET("/", ssr.IndexHandler())
	router.GET("/password", ssr.PasswordHandler())
	router.PUT("/password/change", ssr.ChangePasswordHandler(*server))

	routerResUser := router.Group("/reservations")
	{
		routerResUser.GET("/mine", ssr.MineReservationsHandler(*server))
		routerResUser.GET("/new", ssr.NewReservationHandler(*server))
		routerResUser.POST("/", ssr.CreateReservationHandler(*server))
		routerResUser.GET("/qr_code/:id", ssr.GenerateQRCodeHandler(*server))
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
			adminUserRoutes.GET("/:id", ssr.EditUserHandler(*server))
			adminUserRoutes.PUT("/:id", ssr.UpdateUserHandler(*server))
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
			tourRoutes.POST("/import", ssr.ImportToursHandler(*server))
		}

		hotelRoutes := adminRoute.Group("/hotels")
		{
			hotelRoutes.GET("/", ssr.ListHotelsHandler(*server))
			hotelRoutes.GET("/new", ssr.NewHotelHandler(*server))
			hotelRoutes.POST("/", ssr.CreateHotelHandler(*server))
			hotelRoutes.GET("/:id/edit", ssr.EditHotelHandler(*server))
			hotelRoutes.DELETE("/:id", ssr.DeleteHotelHandler(*server))
			hotelRoutes.POST("/import", ssr.ImportHotelsHandler(*server))
		}

		resRoutes := adminRoute.Group("/reservations")
		{
			resRoutes.GET("/", ssr.ListReservationsHandler(*server))
			resRoutes.GET("/:id/edit", ssr.EditReservationHandler(*server))
			resRoutes.PUT("/:id", ssr.UpdateReservationHandler(*server))
			resRoutes.DELETE("/:id", ssr.DeleteReservationHandler(*server))
			resRoutes.GET("/export", ssr.ExportReservationsHandler(*server))
		}
	}
	server.Router = router
}
