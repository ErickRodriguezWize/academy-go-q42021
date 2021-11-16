package router

import (
	"log"
	"net/http"
	"fmt"

	"github.com/ErickRodriguezWize/academy-go-q42021/interface/controller"

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
		Addr: "127.0.0.1:8000",
	}

	fmt.Printf("Server Started at %v", srv.Addr)

	log.Fatal(srv.ListenAndServe())
}
