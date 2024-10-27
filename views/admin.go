package views

import (
	"github.com/a-h/templ"
	"github.com/leedrum/ikarus_travel/model"
	admin_pages "github.com/leedrum/ikarus_travel/views/admin"
)

func NewUser(user model.User) templ.Component {
	return admin_pages.NewUser(user)
}

func CreateSuccessUser(user model.User) templ.Component {
	return admin_pages.CreateSuccessUser(user)
}

func ListUsers(users []model.User) templ.Component {
	return admin_pages.ListUsers(users)
}

func NewTour() templ.Component {
	return admin_pages.NewTour()
}

func ListTours(tours []model.Tour) templ.Component {
	return admin_pages.ListTours(tours)
}

func EditTour(tour model.Tour) templ.Component {
	return admin_pages.EditTour(tour)
}

func NewHotel() templ.Component {
	return admin_pages.NewHotel()
}

func ListHotels(hotels []model.Hotel) templ.Component {
	return admin_pages.ListHotels(hotels)
}

func EditHotel(hotel model.Hotel) templ.Component {
	return admin_pages.EditHotel(hotel)
}

func NewReservation(hotels []model.Hotel, tours []model.Tour) templ.Component {
	return admin_pages.NewReservation(hotels, tours)
}

func ListReservations(reservations []model.Reservation) templ.Component {
	return admin_pages.ListReservations(reservations)
}

func EditReservation(reservation model.Reservation, hotels []model.Hotel, tours []model.Tour) templ.Component {
	return admin_pages.EditReservation(reservation, hotels, tours)
}

func AdminDashboard() templ.Component {
	return admin_pages.AdminDashboard()
}

func LoginForm() templ.Component {
	return admin_pages.LoginForm()
}
