package interactor

import (
	"github.com/ErickRodriguezWize/academy-go-q42021/domain/model"
)

// interface that will handle all file methods for implementacion.
type iFileService interface {
	ReadAll(pkms *[]model.Pokemon) error
	Write(artist model.Artist) error
}

// interface that will handle all spotify methods for implementacion.
type iSpotifyService interface {
	SearchArtist(artist string) (model.Artist, error)
}

// Artist Interator struct that will contain all the interfaces.
type ArtistInteractor struct {
	fileService    iFileService
	spotifyService iSpotifyService
}

// NewArtistInteractor: Construct for ArtistInteractor struct and implement interfaces.
func NewArtistInteractor(ifs iFileService, isp iSpotifyService) *ArtistInteractor {
	return &ArtistInteractor{ifs, isp}
}

// SearchArtist: Handle use of SearchArtist service
func (ai *ArtistInteractor) SearchArtist(name string) (model.Artist, error) {
	// Search artis using Spotify Service
	foundArtist, err := ai.spotifyService.SearchArtist(name)
	if err != nil {
		return model.Artist{}, err
	}

	// Store it into the CSV.
	if err := ai.StoreArtist(foundArtist); err != nil {
		return model.Artist{}, err
	}

	return foundArtist, nil
}

// StoreArtist: Interactor method to handle storeartist service.
func (ai *ArtistInteractor) StoreArtist(artist model.Artist) error {
	// Write the response from External API into a file.
	if err := ai.fileService.Write(artist); err != nil {
		return err
	}

	return nil
}