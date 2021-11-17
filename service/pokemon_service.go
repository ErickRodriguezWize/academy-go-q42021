package service

import (
	"errors"

	"github.com/ErickRodriguezWize/academy-go-q42021/domain/model"
)

var (
	PokemonMissingError = errors.New("Couldn't find pokemon")
)

//GetPokemonById finds a pokemon inside pokemons array using the ID as filter. 
func GetPokemonByID(pokemons []model.Pokemon, ID int) (model.Pokemon, error){
	targetPkm := model.Pokemon{}

	for _,pkm := range pokemons{
		if pkm.ID == ID {
			targetPkm = pkm
			break
		}
	}

	if targetPkm.Name == "" {
		return targetPkm, PokemonMissingError
	}

	return targetPkm, nil

}