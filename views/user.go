package views

import (
	"github.com/a-h/templ"
	user_pages "github.com/leedrum/ikarus_travel/views/user"
)

func QRCode(content string) templ.Component {
	return user_pages.QRCode(content)
}
