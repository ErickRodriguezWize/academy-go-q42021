package router

import (
	"log"
	"net/http"

	"github.com/ErickRodriguezWize/academy-go-q42021/interface/controller"
	"github.com/ErickRodriguezWize/academy-go-q42021/service"

	"github.com/gorilla/mux"
)

func NewRouter() {
	r := mux.NewRouter()

	pokemonHandler := controller.NewPokemonController()

	//Pokemon Endpoints
	r.HandleFunc("/pokemons", pokemonHandler.GetAllPokemons).Methods("GET")
	r.HandleFunc("/pokemons/{id}", pokemonHandler.GetPokemon).Methods("GET")

	//Server setup
	srv := &http.Server{
		Handler: r,
		Addr: service.GetEnvVariable("IP_ADDRESS") + ":" + service.GetEnvVariable("PORT"),
	}

	log.Println("Server Started at", srv.Addr)

	log.Fatal(srv.ListenAndServe())

}