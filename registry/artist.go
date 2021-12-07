package registry

import (
	"github.com/ErickRodriguezWize/academy-go-q42021/interface/controller"
	"github.com/ErickRodriguezWize/academy-go-q42021/usecase/interactor"
)

// NewArtistController: Register - ArtistController 
func (r *Registry) NewArtistController() *controller.ArtistController {
	return controller.NewArtistController(r.NewArtistInteractor())
}

// NewArtistController: Register - NewArtistInteractor
func (r *Registry) NewArtistInteractor() interactor.ArtistInteractor {
	return interactor.NewArtistInteractor(r.NewCsvService(), r.NewSpotifyService())
}