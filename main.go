package main

import "github.com/leedrum/ikarus_travel/routes"

func main() {
	router := routes.InitRoutes()
	router.Run(":8080")
}
