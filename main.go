package main

import (
	"beeapi/config"
	"beeapi/database"
	"beeapi/routes"
)

func main() {
	config.Init()

	database.Init()

	routes.Init()

	routes.RegisterRoutes()

	err := routes.GetRouter().Run()
	if err != nil {
		panic(err)
	}
}
