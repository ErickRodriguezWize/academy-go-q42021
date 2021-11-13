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

//Structura vacia del controlador. 
type pokemonController struct{}

type PokemonController interface{
	GetAllPokemon()
	GetPokemon()
}



func (pc *pokemonController) GetAllPokemons(res http.ResponseWriter, req *http.Request)  {
	
	pokemons := []model.Pokemon{}
	err := service.ReadCSV("./test/bateria_csv/pokemon.csv", &pokemons)

	if err != nil {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(err.Error()))
	}

	//leer csv
	// generar array del csv

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(pokemons)
}

func (pc *pokemonController) GetPokemon(res http.ResponseWriter, req *http.Request) {
	pokemons := []model.Pokemon{}
	err := service.ReadCSV("./test/bateria_csv/pokemon.csv", &pokemons)

	if err != nil {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(err.Error()))
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
}

//Funcion para obtener una structura vacia del Controlador a la hora de llamarla por paquete. 
func NewPokemonController() *pokemonController{
	return &pokemonController{}
}

/*

Codigo a reutilziar 

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