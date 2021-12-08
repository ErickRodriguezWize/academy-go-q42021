package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/ErickRodriguezWize/academy-go-q42021/domain/model"
	querror "github.com/ErickRodriguezWize/academy-go-q42021/errors"

	"github.com/gorilla/mux"
)

// Interface to handle usecase/interactor methods for PokemonController.
type interactorPokemon interface {
	GetAllPokemons() ([]model.Pokemon, error)
	GetPokemon(ID int) (model.Pokemon, error)
	GetPokemonWorker(t string, items int, itemsPerWorker int) ([]model.Pokemon, error)
}

// PokemonController struct with the interfaces to implement
type PokemonController struct {
	service interactorPokemon
}

// GetAllPokemon: returns all values inside the csv File.
func (pc *PokemonController) GetAllPokemons(res http.ResponseWriter, req *http.Request) {
	log.Println("HTTP GET /pokemons")

	// Get an array of model.Pokemon from service "ReadCSV".
	pokemons, err := pc.service.GetAllPokemons()
	if err != nil {
		log.Println("Error: " + err.Error())
		http.Error(res, err.Error(), http.StatusBadRequest)

		return
	}

	// Setup response (headers, http Status)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)

	// Handling Response Json.
	if err := json.NewEncoder(res).Encode(pokemons); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)

		return
	}

}

// GetPokemon: Returns a Pokemon using the ID as a filter.
func (pc *PokemonController) GetPokemon(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	log.Printf("HTTP GET /pokemons/%v \n", params["id"])

	// Parsing the 'id' from string into int.
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Println("ERROR: Couldn't parse Properly to Integer.")
		http.Error(res, querror.ErrParseError.Error(), http.StatusBadRequest)

		return
	}

	// Search for pokemon using his ID.
	pokemon, err := pc.service.GetPokemon(id)
	if err != nil {
		log.Println("Error:" + err.Error())
		http.Error(res, err.Error(), http.StatusBadRequest)

		return
	}

	// Setup response (headers, http Status)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)

	// Handling Response Json.
	if err := json.NewEncoder(res).Encode(pokemon); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)

		return
	}

}

// GetPokemonWorker: Get a n items of Pokemon using Workers.
func (pc *PokemonController) GetPokemonsWorker(res http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	t := params.Get("type")

	if t == "" {
		log.Println("Query param type is missing.")
		http.Error(res, "error: query param type is missing.", http.StatusBadRequest)

		return

	}

	// Validate if the type query params has the value odd or even.
	if t != "odd" && t != "even" {
		log.Println("ERROR: Invalid type value in query params.", t)
		http.Error(res, querror.ErrInvalidTypeParams.Error(), http.StatusBadRequest)

		return
	}

	// Parsing the 'item' from string into int.
	items, err := strconv.Atoi(params.Get("items"))
	if err != nil {
		log.Println("ERROR: Couldn't parse items Properly to Integer.")
		http.Error(res, querror.ErrParseError.Error(), http.StatusBadRequest)

		return
	}

	// Parsing the 'itemsWorker' from string into int.
	itemsWorker, err := strconv.Atoi(params.Get("items_per_worker"))
	if err != nil {
		log.Println("ERROR: Couldn't parse itemsWorker Properly to Integer.")
		http.Error(res, querror.ErrParseError.Error(), http.StatusBadRequest)

		return
	}

	log.Printf("HTTP GET /pokemons/worker?type=%v&item=%v&item_per_worker=%v \n", t, items, itemsWorker)
	// Get pokemons using the WorkerPool
	results, err := pc.service.GetPokemonWorker(t, items, itemsWorker)
	if err != nil {
		log.Println(err)
		http.Error(res, err.Error(), http.StatusBadRequest)

		return
	}

	// Setup response (headers, http Status)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)

	// Handling Response Json.
	if err := json.NewEncoder(res).Encode(results); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)

		return
	}

}

// NewPokemonController: Initialiazed PokemonController struct and implement interactorPokemon interface.
func NewPokemonController(ip interactorPokemon) *PokemonController {
	return &PokemonController{ip}
}