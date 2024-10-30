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
	router.Use(sessions.Sessions("auth-token", *server.SessionStore))

	// Auth
	router.GET("/login", ssr.LoginFormHandler(*server))
	router.POST("/login", ssr.LoginHandler(*server))
	router.GET("/logout", ssr.LogoutHandler)
	router.Use(middlewares.AuthRequired(*server))
	router.GET("/", ssr.IndexHandler())

	reservationRoutes := router.Group("/reservations")
	{
		reservationRoutes.GET("/", ssr.ListReservationsHandler(*server))
		reservationRoutes.GET("/new", ssr.NewReservationHandler(*server))
		reservationRoutes.POST("/", ssr.CreateReservationHandler(*server))
	}

	admin := router.Group("/admin")
	{
		admin.Use(middlewares.AdminRequired())
		admin.GET("/", ssr.AdminIndexHandler())
		adminUserRoutes := admin.Group("/users")
		{
			adminUserRoutes.GET("/", ssr.ListUsersHandler(*server))
			adminUserRoutes.GET("/new", ssr.NewUserHandler(*server))
			adminUserRoutes.POST("/", ssr.CreateUserHandler(*server))
			adminUserRoutes.DELETE("/:id", ssr.DeleteUserHandler(*server))
		}

		tourRoutes := admin.Group("/tours")
		{
			tourRoutes.GET("/", ssr.ListToursHandler(*server))
			tourRoutes.GET("/new", ssr.NewTourHandler(*server))
			tourRoutes.POST("/", ssr.CreateTourHandler(*server))
			tourRoutes.GET("/:id/edit", ssr.EditTourHandler(*server))
			tourRoutes.PUT("/:id", ssr.UpdateTourHandler(*server))
			tourRoutes.DELETE("/:id", ssr.DeleteTourHandler(*server))
		}

		hotelRoutes := admin.Group("/hotels")
		{
			hotelRoutes.GET("/", ssr.ListHotelsHandler(*server))
			hotelRoutes.GET("/new", ssr.NewHotelHandler(*server))
			hotelRoutes.POST("/", ssr.CreateHotelHandler(*server))
			hotelRoutes.GET("/:id/edit", ssr.EditHotelHandler(*server))
			hotelRoutes.DELETE("/:id", ssr.DeleteHotelHandler(*server))
		}

		reservationAdminRoutes := admin.Group("/reservations")
		{
			reservationAdminRoutes.GET("/", ssr.ListReservationsHandler(*server))
			reservationAdminRoutes.GET("/new", ssr.NewReservationHandler(*server))
			reservationAdminRoutes.POST("/", ssr.CreateReservationHandler(*server))
			reservationAdminRoutes.GET("/:id/edit", ssr.EditReservationHandler(*server))
			reservationAdminRoutes.PUT("/:id", ssr.UpdateReservationHandler(*server))
			reservationAdminRoutes.DELETE("/:id", ssr.DeleteReservationHandler(*server))
			reservationAdminRoutes.GET("/mine", ssr.MineReservationsHandler(*server))
		}
	}
	server.Router = router
}
