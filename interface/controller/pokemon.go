package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/ErickRodriguezWize/academy-go-q42021/domain/model"
	"github.com/ErickRodriguezWize/academy-go-q42021/service"

	"github.com/gorilla/mux"
)

type pokemonController struct{}

// GetAllPokemon: returns all values inside the csv File.
func (pc *pokemonController) GetAllPokemons(res http.ResponseWriter, req *http.Request) {
	log.Println("HTTP GET /pokemons")
	CsvPath := service.GetEnvVariable("CSV_PATH")

	//Get an array of model.Pokemon from service "ReadCSV".
	pokemons := []model.Pokemon{}
	if err := service.ReadCSV(CsvPath, &pokemons); err != nil {
		log.Println("Error: " + err.Error())
		http.Error(res, err.Error(), http.StatusBadRequest)

		return
	}

	//Setup response (headers, http Status)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)

	//Handling Response Json.
	if err := json.NewEncoder(res).Encode(pokemons); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)

		return
	}

}

//GetPokemon: Returns a Pokemon using the ID as a filter.
func (pc *pokemonController) GetPokemon(res http.ResponseWriter, req *http.Request) {
	pkm := model.Pokemon{}
	params := mux.Vars(req)
	CsvPath := service.GetEnvVariable("CSV_PATH")
	log.Printf("HTTP GET /pokemons/%v \n", params["id"])

	//Parsing the 'id' from string into int.
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Println("ERROR: Couldn't parse Properly to Integer.")
		http.Error(res, err.Error(), http.StatusBadRequest)

		return
	}

	//Get an array of model.Pokemon from service "ReadCSV".
	pokemons := []model.Pokemon{}
	if err := service.ReadCSV(CsvPath, &pokemons); err != nil {
		log.Println("Error:" + err.Error())
		http.Error(res, err.Error(), http.StatusBadRequest)

		return
	}

	//Find a pokemon using ID as a filter.
	pkm, err = service.GetPokemonByID(pokemons, id)
	if err != nil {
		log.Printf("ERROR: Couldn't find pokemon with id: %v \n", id)
		http.Error(res, err.Error(), http.StatusBadRequest)

		return
	}

	//Setup response (headers, http Status)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)

	//Handling Response Json.
	if err := json.NewEncoder(res).Encode(pkm); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)

		return
	}

}

//NewPokemonController: Returns an empty Struct of pokemonController.
func NewPokemonController() *pokemonController {
	return &pokemonController{}

}
