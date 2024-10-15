package main

import (
	"log"

	"github.com/invopop/ctxi18n"
	"github.com/leedrum/ikarus_travel/locales"
	"github.com/leedrum/ikarus_travel/routes"
)

func main() {
	initI18n()
	initRouter()
}

func initI18n() {
	if err := ctxi18n.Load(locales.Content); err != nil {
		log.Fatalf("error loading locales: %v", err)
	}
}

func initRouter() {
	router := routes.InitRoutes()
	router.Run(":8080")
}
