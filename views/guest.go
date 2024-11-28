package views

import (
	"github.com/a-h/templ"
	"github.com/leedrum/ikarus_travel/model"

	guest_pages "github.com/leedrum/ikarus_travel/views/guest"
)

func PreviewReservation(reservation model.Reservation) templ.Component {
	return guest_pages.PreviewReservation(reservation)
}
