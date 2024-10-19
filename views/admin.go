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
