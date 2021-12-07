package router

import (
	"log"
	"net/http"
	"time"

	"github.com/ErickRodriguezWize/academy-go-q42021/domain/model"
	"github.com/ErickRodriguezWize/academy-go-q42021/interface/controller"

	"github.com/gorilla/mux"
)

// Router: Struct that use to Initialized server and add Routes to project.
type Router struct {
	mux    *mux.Router
	config model.Config
	app    *controller.AppController
}

// NewRouter: Constructor for Router struct.
func NewRouter(config model.Config, app *controller.AppController) *Router {
	return &Router{
		mux.NewRouter(),
		config,
		app,
	}
}

// CreateRoutes: Manage all Handlers(controller methods) for the project.
func (r *Router) CreateRoutes() {
	handlers := r.app

	//Pokemon Endpoints
	r.mux.HandleFunc("/pokemons", handlers.PokemonHandler.GetAllPokemons).Methods("GET")
	r.mux.HandleFunc("/pokemons/worker", handlers.PokemonHandler.GetPokemonsWorker).Methods("GET")
	r.mux.HandleFunc("/pokemons/{id}", handlers.PokemonHandler.GetPokemon).Methods("GET")

	// Movie Endpoints
	r.mux.HandleFunc("/artists/{artist}", handlers.ArtistHandler.SearchArtist).Methods("GET")
}

// InitServer: Setup configuration for Go server (IP,host, Timeouts).
func (r *Router) InitServer() {
	handler := r.mux

	// Server setup
	srv := &http.Server{
		Handler:      handler,
		Addr:         r.config.Ip + ":" + r.config.Port,
		WriteTimeout: time.Duration(r.config.ReadTimeout) * time.Second,  // Define seconds of WriteTimeout
		ReadTimeout:  time.Duration(r.config.WriteTimeout) * time.Second, // Define seconds of ReadTimeout
	}

	log.Println("Server Started at", srv.Addr)

	// top Server if something goes wrong with server setup.
	log.Fatal(srv.ListenAndServe())
}