package router

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ErickRodriguezWize/academy-go-q42021/config"

	"github.com/gorilla/mux"
)

func TestPokewmonHandler(t *testing.T) {
	//Disabled log ouputs.
	log.SetOutput(ioutil.Discard)

	conf, _ := config.LoadConfig()
	conf.PokemonCsvPath = "./../../" + conf.PokemonCsvPath

	tests := map[string]struct {
		handler   string
		method    string
		endpoint  string
		want_code int
	}{
		"all pokemons":      {handler: "GetAllPokemons", method: "GET", endpoint: "/pokemons", want_code: 200},
		"found pokemon":     {handler: "GetPokemon", method: "GET", endpoint: "/pokemons/5", want_code: 200},
		"not found pokemon": {handler: "GetPokemon", method: "GET", endpoint: "/pokemons/167", want_code: 400},
	}

	m := mux.NewRouter()
	AddRoutes(m, conf)

	for name, tsc := range tests {
		t.Run(name, func(t *testing.T) {
			req, _ := http.NewRequest(tsc.method, tsc.endpoint, nil)
			w := httptest.NewRecorder()
			m.ServeHTTP(w, req)
			if w.Code != tsc.want_code {
				t.Fatalf("Expected HTTP Code:%d, Actual HTTP Code: %d ", tsc.want_code, w.Code)
			}
		})
	}

}

func TestArtistHandler(t *testing.T) {
	//Disabled log ouputs.
	log.SetOutput(ioutil.Discard)

	conf, _ := config.LoadConfig()
	conf.ArtistCsvPath = "./../../" + conf.ArtistCsvPath

	// test cases.
	tests := map[string]struct {
		handler   string
		method    string
		artist    string
		endpoint  string
		want_code int
	}{
		"simple artist":      {handler: "SearchArtist", method: "GET", artist: "linkin+park", endpoint: "/artists/linkin+park", want_code: 200},
		"artist with a coma": {handler: "SearchArtist", method: "GET", artist: "linki,park", endpoint: "/artist/linkin,park", want_code: 404},
		"not found artist":   {handler: "SearchArtist", method: "GET", artist: "linkin+bizkif", endpoint: "/artist/linkin+bizkif", want_code: 404},
	}

	m := mux.NewRouter()
	AddRoutes(m, conf)

	// Test table cases.
	for name, tsc := range tests {
		t.Run(name, func(t *testing.T) {

			req, _ := http.NewRequest(tsc.method, tsc.endpoint, nil)
			w := httptest.NewRecorder()
			m.ServeHTTP(w, req)
			if w.Code != tsc.want_code {
				t.Fatalf("Expected HTTP CODE:%d , actual HTTP CODE: %d", tsc.want_code, w.Code)
			}
		})
	}

}