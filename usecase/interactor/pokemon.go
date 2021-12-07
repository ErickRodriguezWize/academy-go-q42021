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
	// Read and handle errors from CsvService.ReadCSV
	if err := pi.CsvService.ReadCSV(pokemons); err != nil {
		return err
	}
	return nil

}

// GetPokemon: Interactor method that handles Get Pokemon with ID.
func (pi *PokemonInteractor) GetPokemon(ID int) (model.Pokemon, error) {
	var pokemons []model.Pokemon
	// Read and handle errors from CsvService.ReadCSV
	if err := pi.CsvService.ReadCSV(&pokemons); err != nil {
		return model.Pokemon{}, err
	}

	// Get the results and posible errors from GetPokemonByID
	pkm, err := pi.PokemonService.GetPokemonByID(pokemons, ID)
	if err != nil {
		return model.Pokemon{}, err
	}

	return pkm, nil
}

// GetPokemonWorker: Interactor method that handles the Retrieve of pokemons using worker pool.
func (pi *PokemonInteractor) GetPokemonWorker(t string, items int, itemsPerWorker int) ([]model.Pokemon, error) {
	var pokemons []model.Pokemon
	// Read and handle errors from CsvService.ReadCSV
	if err := pi.CsvService.ReadCSV(&pokemons); err != nil {
		return pokemons, err
	}

	// Get the results and posible errors from GetPokemonWorker (worker pool method. )
	results, err := pi.PokemonService.GetPokemonWorker(pokemons, t, items, itemsPerWorker)
	if err != nil {
		return results, err
	}

	return results, nil
}