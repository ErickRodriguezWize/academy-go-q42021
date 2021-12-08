package registry

import (
	"github.com/ErickRodriguezWize/academy-go-q42021/domain/model"
	"github.com/ErickRodriguezWize/academy-go-q42021/interface/controller"
)

type Registry struct {
	config model.Config
}

// NewRegistry: Construct for Registry struct.
func NewRegistry(c model.Config) *Registry {
	return &Registry{c}
}

// NewAppController: Construct for controller.AppController struct. Also make the dependency injectio for all dependencies (controller, interactors, services)
func (r *Registry) NewAppController() *controller.AppController {
	return &controller.AppController{
		r.NewPokemonController(),
		r.NewArtistController(),
	}
}