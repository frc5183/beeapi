package main

import (
	"beeapi/config"
	"beeapi/database"
	"beeapi/routes"
	"github.com/rs/zerolog/log"
)

func main() {
	config.GenerateConfig()

	log.Info().Msg("wsg")

	database.Init()

	routes.RegisterRoutes()

	err := routes.GetRouter().Run()
	if err != nil {
		panic(err)
	}
}
