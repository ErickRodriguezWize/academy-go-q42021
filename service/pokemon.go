package service

import (
	"github.com/ErickRodriguezWize/academy-go-q42021/domain/model"
	pokerr "github.com/ErickRodriguezWize/academy-go-q42021/errors"
)

type PokemonService struct{}

// GetPokemonById: finds a pokemon inside pokemons array using the ID as filter.
func (sp PokemonService) GetPokemonByID(pokemons []model.Pokemon, ID int) (model.Pokemon, error) {
	targetPkm := model.Pokemon{}

	// Loop through the pokemosn to find the pokemon with the correct ID
	for _, pkm := range pokemons {
		if pkm.ID == ID {
			return pkm, nil
		}
	}

	return targetPkm, pokerr.ErrPokemonNotFound

}

// NewPokemonService: Constructor for PokemonService struct. 
func NewPokemonService() PokemonService {
	return PokemonService{}
}