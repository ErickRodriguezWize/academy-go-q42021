package main

import (
	"log"

	"github.com/ErickRodriguezWize/academy-go-q42021/config"
	"github.com/ErickRodriguezWize/academy-go-q42021/infrastructure/router"
	"github.com/ErickRodriguezWize/academy-go-q42021/registry"
)

func main() {
	// Initialize config struct with environment variables.
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Intialize controllers and dependencies.
	reg := registry.NewRegistry(config)
	app := reg.NewAppController()

	// Initialize Routing Handling
	router := router.NewRouter(config, app)
	router.CreateRoutes()

	// Init and serve Go Server.
	router.InitServer()
}