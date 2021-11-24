package service

import (
	"github.com/ErickRodriguezWize/academy-go-q42021/domain/model"
	pokerr "github.com/ErickRodriguezWize/academy-go-q42021/errors"
)

//GetPokemonById: finds a pokemon inside pokemons array using the ID as filter.
func GetPokemonByID(pokemons []model.Pokemon, ID int) (model.Pokemon, error) {
	targetPkm := model.Pokemon{}

	for _, pkm := range pokemons {
		if pkm.ID == ID {
			return pkm, nil
		}
	}

	return targetPkm, pokerr.PokemonNotFound

}
