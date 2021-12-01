package main

import (
	"github.com/ErickRodriguezWize/academy-go-q42021/config"
	"github.com/ErickRodriguezWize/academy-go-q42021/infrastructure/router"
)

func main() {
	// Initialize config struct with environment variables.
	config := config.LoadConfig()

	// Initialize Routing Handling
	router := router.NewRouter()
	router.Init(config)
}