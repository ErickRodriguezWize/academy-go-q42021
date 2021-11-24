package router

import (
	"log"
	"net/http"
	"time"

	"github.com/ErickRodriguezWize/academy-go-q42021/interface/controller"
	"github.com/ErickRodriguezWize/academy-go-q42021/service"

	"github.com/gorilla/mux"
)

// NewRouter: Make the Routing Setup for the App with viper(package).
func NewRouter() {
	r := mux.NewRouter()

	//Pokemon Endpoints
	pokemonHandler := controller.NewPokemonController()
	r.HandleFunc("/pokemons", pokemonHandler.GetAllPokemons).Methods("GET")
	r.HandleFunc("/pokemons/{id}", pokemonHandler.GetPokemon).Methods("GET")

	//Movie Endpoints
	artistHandler := controller.NewArtistController()
	r.HandleFunc("/artists/{artist}", artistHandler.SearchArtist).Methods("GET")

	readTime := service.GetIntEnvVariable("READ_TIMEOUT")
	writeTime := service.GetIntEnvVariable("WRITE_TIMEOUT")

	//Server setup
	srv := &http.Server{
		Handler:      r,
		Addr:         service.GetEnvVariable("IP_ADDRESS") + ":" + service.GetEnvVariable("PORT"),
		WriteTimeout: time.Duration(readTime) * time.Second,  //Define seconds of WriteTimeout
		ReadTimeout:  time.Duration(writeTime) * time.Second, //Define seconds of ReadTimeout
	}

	log.Println("Server Started at", srv.Addr)

	//Stop Server if something goes wrong with server setup.
	log.Fatal(srv.ListenAndServe())

}
