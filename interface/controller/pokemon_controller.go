package controller

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"encoding/json"
	"WizelineApi/domain/model"
	"strconv"
)


type pokemonController struct{}

type PokemonController interface{
	GetAllPokemon()
	GetPokemon()
}

pokemon_controller = pokemonController{}

func (pc *pokemonController) GetAllPokemons(res http.ResponseWriter, req *http.Request)  {
	
	//leer csv
	// generar array del csv

	pokemons := []model.Pokemon{
		{1, "Bulbasaur"},
		{2, "Ivysaur"},
		{3, "Venasaur"},
	}

	jsonR, jsonErr := json.Marshal(pokemons)
	
	if jsonErr != nil {
		fmt.Println("Json unable to encode")
	}

	res.Write(jsonR)
}

func (pc *pokemonController) GetPokemon(res http.ResponseWriter, req *http.Request) {
	pokemons := []model.Pokemon{
		{1, "Bulbasaur"},
		{2, "Ivysaur"},
		{3, "Venasaur"},
	}

	var t_pkm model.Pokemon
	params := mux.Vars(req)
	id,err := strconv.Atoi(params["id"])

	if err != nil {
		fmt.Println("Couldn't parse to Integer")
	}

	for _,pk := range pokemons{
		if pk.ID == id {
			t_pkm = pk
			break
		}
	}

	jsonR, _ := json.Marshal(t_pkm)

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(jsonR)

}

func renderResponse(*res http.ResponseWriter, err Error, json interface{}){
	res.Header.Set("Content-Type", "application/json")
	if err != nil {
		res.WriteHedear(http.BadRequest)
		res.Write(err)
	}else{
		res.WriteHeader(http.StatusOK)
		res.Writer(json)
	}

}

func NewPokemonController() *pokemonController{
	return &pokemonController{}
}
