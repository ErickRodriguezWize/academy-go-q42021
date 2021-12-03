package router

import (
	"log"
	"net/http"
	"time"

	"github.com/ErickRodriguezWize/academy-go-q42021/domain/model"
	"github.com/ErickRodriguezWize/academy-go-q42021/interface/controller"

	"github.com/gorilla/mux"
)

type router struct{}

// NewRouter: Returns an emptu router Struct.
func NewRouter() router {
	return router{}
}

// NewRouter: Make the Routing Setup for the App with viper(package).
func (router) Init(config *model.Config) {
	r := mux.NewRouter()

	AddRoutes(r, config)
	initServer(r, *config)
}

func AddRoutes(r *mux.Router, config *model.Config){
	// Pokemon Endpoints
	pokemonHandler := controller.NewPokemonController(*config)
	r.HandleFunc("/pokemons", pokemonHandler.GetAllPokemons).Methods("GET")
	r.HandleFunc("/pokemons/{id}", pokemonHandler.GetPokemon).Methods("GET")

	// Movie Endpoints
	artistHandler := controller.NewArtistController(*config)
	r.HandleFunc("/artists/{artist}", artistHandler.SearchArtist).Methods("GET")
}

// initServer: Setup configuration for Go server (IP,host, Timeouts).
func initServer(r http.Handler, config model.Config) {
	// Server setup
	srv := &http.Server{
		Handler:      r,
		Addr:         config.Ip + ":" + config.Port,
		WriteTimeout: time.Duration(config.ReadTimeout) * time.Second,  // Define seconds of WriteTimeout
		ReadTimeout:  time.Duration(config.WriteTimeout) * time.Second, // Define seconds of ReadTimeout
	}

	log.Println("Server Started at", srv.Addr)

	// top Server if something goes wrong with server setup.
	log.Fatal(srv.ListenAndServe())
}