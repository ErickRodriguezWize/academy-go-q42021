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

	// Dependency Injection using Register local package.
	reg := registry.NewRegistry(config)
	app := reg.NewAppController() // app that will contain all controllers with their implementations.

	// Initialize Routing Handling
	router := router.NewRouter(config, app)
	router.CreateRoutes()

	// Init and serve Go Server.
	router.InitServer()
}