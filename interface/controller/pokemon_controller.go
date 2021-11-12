package controller

import (
	"net/http"
	//"fmt"
	//"errors"
	"github.com/gorilla/mux"
	"encoding/json"
	"WizelineApi/domain/model"
	"WizelineApi/service"
	"strconv"
)


type pokemonController struct{}

type PokemonController interface{
	GetAllPokemon()
	GetPokemon()
}



func (pc *pokemonController) GetAllPokemons(res http.ResponseWriter, req *http.Request)  {
	pokemons := []model.Pokemon{
		{1, "Bulbasaur"},
		{2, "Ivysaur"},
		{3, "Venasaur"},
	}	

	//leer csv
	// generar array del csv

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(pokemons)
}

func (pc *pokemonController) GetPokemon(res http.ResponseWriter, req *http.Request) {
	pokemons := []model.Pokemon{
		{1, "Bulbasaur"},
		{2, "Ivysaur"},
		{3, "Venasaur"},
	}	

	pkm := model.Pokemon{}

	params := mux.Vars(req)
	id,err := strconv.Atoi(params["id"])
	if err != nil {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(err.Error()))
		return
	}

	pkm, err = service.GetPokemonById(pokemons, id)
	if err != nil {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		res.Write([]byte(err.Error()))
		return
	}

	//jsonR, _ := json.Marshal(pkm)
	
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	//res.Write(jsonR)
	json.NewEncoder(res).Encode(pkm)
	return
}

func NewPokemonController() *pokemonController{
	return &pokemonController{}
}

/*
func renderResponse(res *http.ResponseWriter, err error, json interface{}){
	res.Header().Set("Content-Type", "application/json")
	if err != nil{
		res.WriteHeader(http.StatusBadRequest)
		res.Writer([]byte("Something Went Wrong!!"))
	}else{
		res.Writeheader(http.StatusOK)
		res.Writer(json)
	}
}
*/