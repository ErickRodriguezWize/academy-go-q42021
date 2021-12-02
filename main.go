package main

import (
	"log"

	"github.com/ErickRodriguezWize/academy-go-q42021/config"
	"github.com/ErickRodriguezWize/academy-go-q42021/infrastructure/router"
)

func main() {
	// Initialize config struct with environment variables.
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize Routing Handling
	router := router.NewRouter()
	router.Init(config)
}