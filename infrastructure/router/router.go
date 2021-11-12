package router

import (
	"WizelineApi/interface/controller"
	//"academy-go-q42021/interface/controller"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
)



func NewRouter() {
	r := mux.NewRouter()

	handler_pkm := controller.NewPokemonController()
	//Pokemon Endpoints
	r.HandleFunc("/pokemons", handler_pkm.GetAllPokemons).Methods("GET")
	r.HandleFunc("/pokemon/{id}", handler_pkm.GetPokemon).Methods("GET")

	//Server setup
	srv := &http.Server{
		Handler: r,
		Addr: "127.0.0.1:8000",
	}

	fmt.Printf("Server Started at %v", srv.Addr)

	log.Fatal(srv.ListenAndServe())
}
