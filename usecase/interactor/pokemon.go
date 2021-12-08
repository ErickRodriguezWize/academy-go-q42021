package interactor

import (
	"github.com/ErickRodriguezWize/academy-go-q42021/domain/model"
)

// interface that will handle all the  methods need it for the implementation.
type iPokemonService interface {
	GetPokemonByID(pokemons []model.Pokemon, ID int) (model.Pokemon, error)
	GetPokemonWorker(pokemons []model.Pokemon, t string, items int, itemsPerWorker int) ([]model.Pokemon, error)
}

// Pokemon Interator struct that will contain all the interfaces implementations.
type PokemonInteractor struct {
	pokemonService iPokemonService
	fileService    iReadService
}

// NewPokemonInteractor: Construcot for PokemonInteractor struct and implement interfaces.
func NewPokemonInteractor(ips iPokemonService, irs iReadService) *PokemonInteractor {
	return &PokemonInteractor{ips, irs}
}

// GetAllPokemon: Interactor that handle GetAllPokemon logic.
func (pi *PokemonInteractor) GetAllPokemons() ([]model.Pokemon, error) {
	// Read and handle errors from file service: read
	pokemons, err :=  pi.fileService.ReadAll()
	if err != nil {
		return pokemons, err
	}
	return pokemons, nil
}

// GetPokemon: Interactor method that handles Get Pokemon with ID.
func (pi *PokemonInteractor) GetPokemon(ID int) (model.Pokemon, error) {
	var pokemons []model.Pokemon
	// Read and handle errors from CsvService.ReadCSV
	pokemons, err := pi.fileService.ReadAll()
	if  err != nil {
		return model.Pokemon{}, err
	}

	// Get the results and posible errors from GetPokemonByID
	pkm, err := pi.pokemonService.GetPokemonByID(pokemons, ID)
	if err != nil {
		return model.Pokemon{}, err
	}

	return pkm, nil
}

// GetPokemonWorker: Interactor method that handles the Retrieve of pokemons using worker pool.
func (pi *PokemonInteractor) GetPokemonWorker(t string, items int, itemsPerWorker int) ([]model.Pokemon, error) {
	// Read and handle errors from CsvService.ReadCSV
	pokemons, err := pi.fileService.ReadAll()
	if err != nil {
		return pokemons, err
	}

	// Get the results and posible errors from GetPokemonWorker (worker pool method. )
	results, err := pi.pokemonService.GetPokemonWorker(pokemons, t, items, itemsPerWorker)
	if err != nil {
		return results, err
	}

	return results, nil
}