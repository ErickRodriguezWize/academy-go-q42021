package router

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ErickRodriguezWize/academy-go-q42021/config"
	"github.com/ErickRodriguezWize/academy-go-q42021/registry"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *mux.Router {
	// Init Config
	config, _ := config.LoadConfig()
	config.PokemonCsvPath = "./../../" + config.PokemonCsvPath
	config.ArtistCsvPath = "./../../" + config.ArtistCsvPath

	// Dependency Injection using Register local package.
	reg := registry.NewRegistry(config)
	app := reg.NewAppController() // app that will contain all controllers with their implementations.

	// Init Router
	router := NewRouter(config, app)
	router.CreateRoutes()

	return router.mux
}

// Global Variable with the Routing configuration for testing on all Handlers Testers.
var m = setupRouter()

// TestRouter_PokeemonHandler: Unit testing for PokemonHandler.
func TestRouter_PokemonHandler(t *testing.T) {
	//Disabled log ouputs.
	log.SetOutput(ioutil.Discard)

	//testcases
	tests := []struct {
		name      string
		handler   string
		method    string
		endpoint  string
		error     error
		hasError  bool
		want_code int
	}{
		{
			name:      "GetAllPokemons OK",
			method:    "GET",
			endpoint:  "/pokemons",
			want_code: 200,
		},
		{
			name:      "GetPokemon OK",
			method:    "GET",
			endpoint:  "/pokemons/60",
			want_code: 200,
		},
		{
			name:      "GetPokemon Not Found",
			method:    "GET",
			endpoint:  "/pokemons/600",
			want_code: 400,
		},
		{
			name:      "Parse Error",
			method:    "GET",
			endpoint:  "/pokemons/titulo",
			want_code: 400,
		},
		{
			name:      "PokemonWorker Even OK",
			method:    "GET",
			endpoint:  "/pokemons/worker?type=even&items=25&items_per_worker=5",
			want_code: 200,
		},
		{
			name:      "PokemonWorker Odd OK",
			method:    "GET",
			endpoint:  "/pokemons/worker?type=odd&items=25&items_per_worker=5",
			want_code: 200,
		},
		{
			name:      "PokemonWorker Invalid Type",
			method:    "GET",
			endpoint:  "/pokemons/worker?type=odding&items=25&items_per_worker=5",
			want_code: 400,
		},
		{
			name:      "PokemonWorker Empty Type",
			method:    "GET",
			endpoint:  "/pokemons/worker?type=&items=25&items_per_worker=5",
			want_code: 400,
		},
		{
			name:      "PokemonWorker Invalid item",
			method:    "GET",
			endpoint:  "/pokemons/worker?type=&items=six&items_per_worker=5",
			want_code: 400,
		},
	}

	for _, tsc := range tests {
		// Run Test.
		t.Run(tsc.name, func(t *testing.T) {
			// Dummy res http.ResponseWriter and req *http.Request
			req, _ := http.NewRequest(tsc.method, tsc.endpoint, nil)
			w := httptest.NewRecorder()

			// Serve router and run request.
			m.ServeHTTP(w, req)

			// Assert results
			assert.EqualValues(t, w.Code, tsc.want_code)
		})
	}

}

func TestRouter_ArtistHandler(t *testing.T) {
	//Disabled log ouputs.
	log.SetOutput(ioutil.Discard)

	//testcases
	tests := []struct {
		name      string
		handler   string
		method    string
		endpoint  string
		error     error
		hasError  bool
		want_code int
	}{
		{
			name:      "SearchArtist: Linkin Park",
			method:    "GET",
			endpoint:  "/artists/linkin+park",
			want_code: 200,
		},
		{
			name:      "Search: Queens",
			method:    "GET",
			endpoint:  "/artists/queens",
			want_code: 200,
		},
		{
			name:      "Not found Artist",
			method:    "GET",
			endpoint:  "/artists/queenssichu",
			want_code: 400,
		},
		{
			name:      "Not found with coma",
			method:    "GET",
			endpoint:  "/artists/papa,roach",
			want_code: 400,
		},
	}

	for _, tsc := range tests {
		// Run Test.
		t.Run(tsc.name, func(t *testing.T) {
			// Dummy res http.ResponseWriter and req *http.Request
			req, _ := http.NewRequest(tsc.method, tsc.endpoint, nil)
			w := httptest.NewRecorder()

			// Serve router and run request.
			m.ServeHTTP(w, req)

			// Assert results
			assert.EqualValues(t, w.Code, tsc.want_code)
		})
	}

}