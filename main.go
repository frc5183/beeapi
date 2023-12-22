package main

import (
	"beeapi/config"
	"beeapi/database"
	"beeapi/routes"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Starting BeeAPI...")

	config.Init()

	database.Init()

	routes.Init()

	routes.RegisterRoutes()

	err := routes.GetRouter().Run()
	if err != nil {
		panic(err)
	}
}
