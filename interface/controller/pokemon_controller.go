package controller

import (
	"net/http"
	"encoding/json"
	"strconv"
	"log"

	"github.com/ErickRodriguezWize/academy-go-q42021/domain/model"
	"github.com/ErickRodriguezWize/academy-go-q42021/service"

	"github.com/gorilla/mux"
)

//Empty Structure
type pokemonController struct{}

type PokemonController interface{
	GetAllPokemon()
	GetPokemon()
}

// GetAllPokemon returns all values inside the csv File. 
func (pc *pokemonController) GetAllPokemons(res http.ResponseWriter, req *http.Request)  {
	log.Println("HTTP GET /pokemons")

	pokemons := []model.Pokemon{}
	if err := service.ReadCSV("./test/bateria_csv/pokemon.csv", &pokemons); err != nil {
		log.Println("Error: " + err.Error() )
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(err.Error()))
		
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(pokemons)

}

//GetPokemon: Returns a Pokemon using the ID as a filter.
func (pc *pokemonController) GetPokemon(res http.ResponseWriter, req *http.Request) {
	log.Println("HTTP GET /pokemons/")

	pokemons := []model.Pokemon{}
	if err := service.ReadCSV("./test/bateria_csv/pokemon.csv", &pokemons); err != nil {
		log.Println("Error:"+err.Error())
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(err.Error()))

		return
	}	

	pkm := model.Pokemon{}

	params := mux.Vars(req)

	log.Printf("HTTP GET /pokemons/%v \n",params["id"])

	id,err := strconv.Atoi(params["id"])
	if err != nil {
		log.Println("ERROR: Couldn't parse Properly to Integer.")
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Couldn't Parse properly to Integer."))
		
		return
	}

	pkm, err = service.GetPokemonByID(pokemons, id)
	if err != nil {
		log.Printf("ERROR: Couldn't find pokemon with id: %v \n",id)
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(err.Error()))
		
		return
	}
	
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(pkm)

}

//Func use to get an empty structure of pokemonController. 
func NewPokemonController() *pokemonController{
	return &pokemonController{}

}