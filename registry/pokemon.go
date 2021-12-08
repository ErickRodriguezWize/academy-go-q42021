package registry

import (
	"github.com/ErickRodriguezWize/academy-go-q42021/interface/controller"
	"github.com/ErickRodriguezWize/academy-go-q42021/usecase/interactor"
)

// NewPokemonController: Register PokemonController
func (r *Registry) NewPokemonController() *controller.PokemonController {
	return controller.NewPokemonController(r.NewPokemonInteractor())
}

// NewPokemonInteractor: Register PokemonInteractor
func (r *Registry) NewPokemonInteractor() *interactor.PokemonInteractor {
	return interactor.NewPokemonInteractor(r.NewPokemonService(), r.NewFileService())
}