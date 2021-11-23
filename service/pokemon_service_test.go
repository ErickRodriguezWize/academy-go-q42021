package service

import (
	"testing"
	"fmt"
	"log"
	"io/ioutil"

	"github.com/ErickRodriguezWize/academy-go-q42021/domain/model"
)

// TestGetPokemonByID - Testing Function that will simulated scenarios for GetPokemonByID
func TestGetPokemonByID(t *testing.T){
	//Disabled log ouputs. 
	log.SetOutput(ioutil.Discard)

	//Defined cases for testing Scenarios. 
	pokemonCases := []int{1, 30, 55, 130}

	//Get pokemons for the search.
	pokemons := []model.Pokemon{}
	if err := ReadCSV("./../test/bateria_csv/pokemon.csv", &pokemons); err != nil {
		log.Fatalf("Couldn't Read CSV")
	}

	//Create the Test for each Cas of GetPokemonById
	for _,ID := range pokemonCases{
		t.Run(fmt.Sprintf("[TEST] GetPokemonById: %v", ID), func(t *testing.T){
			_, errPokemon := GetPokemonByID(pokemons, ID)
			if errPokemon != nil {
				t.Errorf("Error: %v", errPokemon.Error())
			}
		})
	}

}