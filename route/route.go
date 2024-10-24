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
		admin.GET("/users", ssr.ListUsersHandler(*server))
		admin.DELETE("/users/:id", ssr.DeleteUserHandler(*server))

		admin.GET("/tours/new", ssr.NewTourHandler(*server))
		admin.POST("/tours", ssr.CreateTourHandler(*server))
		admin.GET("/tours", ssr.ListToursHandler(*server))
		admin.GET("/tours/:id/edit", ssr.EditTourHandler(*server))
		admin.PUT("/tours/:id", ssr.UpdateTourHandler(*server))
		admin.DELETE("/tours/:id", ssr.DeleteTourHandler(*server))

		admin.GET("/hotels/new", ssr.NewHotelHandler(*server))
		admin.POST("/hotels", ssr.CreateHotelHandler(*server))
		admin.GET("/hotels", ssr.ListHotelsHandler(*server))
		admin.GET("/hotels/:id/edit", ssr.EditHotelHandler(*server))
		admin.DELETE("/hotels/:id", ssr.DeleteHotelHandler(*server))

		admin.GET("/reservations/new", ssr.NewReservationHandler(*server))
		admin.POST("/reservations", ssr.CreateReservationHandler(*server))
		admin.GET("/reservations", ssr.ListReservationsHandler(*server))
		admin.GET("/reservations/:id/edit", ssr.EditReservationHandler(*server))
		admin.PUT("/reservations/:id", ssr.UpdateReservationHandler(*server))
		admin.DELETE("/reservations/:id", ssr.DeleteReservationHandler(*server))
	}
	server.Router = router
}
