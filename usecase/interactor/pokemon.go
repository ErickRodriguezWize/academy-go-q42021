package interactor

import (
	"github.com/ErickRodriguezWize/academy-go-q42021/domain/model"
	"github.com/ErickRodriguezWize/academy-go-q42021/service"
)

type PokemonInteractor struct {
	PokemonService service.PokemonService
	CsvService     service.CsvService
}

// NewPokemonInteractor: Construcot for PokemonInteractor struct.
func NewPokemonInteractor(ps service.PokemonService, cs service.CsvService) PokemonInteractor {
	return PokemonInteractor{ps, cs}
}

// GetAllPokemon: Interactor Method that handles CsvService.ReadCSV and retrieve pokemons.
func (pi *PokemonInteractor) GetAllPokemons(pokemons *[]model.Pokemon) error {
	if err := pi.CsvService.ReadCSV(pokemons); err != nil {
		return err
	}
	return nil

}

// GetPokemon: Interactor method that handles Get Pokemon with ID.
func (pi *PokemonInteractor) GetPokemon(ID int) (model.Pokemon, error) {
	var pokemons []model.Pokemon
	if err := pi.CsvService.ReadCSV(&pokemons); err != nil {
		return model.Pokemon{}, err
	}

	pkm, err := pi.PokemonService.GetPokemonByID(pokemons, ID)
	if err != nil {
		return model.Pokemon{}, err
	}

	return pkm, nil
}

func(pi *PokemonInteractor) GetPokemonWorker(t string, items int, itemsPerWorker int)([]model.Pokemon){
	var pokemons []model.Pokemon
	if err := pi.CsvService.ReadCSV(&pokemons); err != nil {
		return pokemons
	}
	results := pi.PokemonService.GetPokemonWorker(pokemons, t, items, itemsPerWorker)

	return results
}