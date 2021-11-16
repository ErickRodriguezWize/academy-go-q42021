package controller

import (
	"net/http"
	"encoding/json"
	"strconv"

	"github.com/ErickRodriguezWize/academy-go-q42021/domain/model"
	"github.com/ErickRodriguezWize/academy-go-q42021/service"

	"github.com/gorilla/mux"
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

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(pokemons)
}

func (pc *pokemonController) GetPokemon(res http.ResponseWriter, req *http.Request)  {
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
		res.Write([]byte("Couldn't Parse properly to Integer."))
		return
	}

	pkm, err = service.GetPokemonById(pokemons, id)
	if err != nil {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		res.Write([]byte(err.Error()))
		return
	}
	
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(pkm)
}

//Funcion para obtener una structura vacia del Controlador a la hora de llamarla por paquete. 
func NewPokemonController() *pokemonController{
	return &pokemonController{}
}
