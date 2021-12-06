package interactor

import (
	"github.com/ErickRodriguezWize/academy-go-q42021/domain/model"
	"github.com/ErickRodriguezWize/academy-go-q42021/service"
)

type ArtistInteractor struct {
	CsvService     service.CsvService
	SpotifyService service.SpotifyService
}

// NewArtistInteractor: Constructo for artistInteractor.
func NewArtistInteractor(csv service.CsvService, sp service.SpotifyService) ArtistInteractor {
	return ArtistInteractor{csv, sp}
}

// SearchArtist: Handle use of SpotifyService/SearchArtist .
func (ai *ArtistInteractor) SearchArtist(name string) (model.Artist, error) {
	artist := model.Artist{}
	// Search artis using Spotify Service
	if err := ai.SpotifyService.SearchArtist(name, &artist); err != nil {
		return model.Artist{}, err
	}

	// Store it into the CSV.
	if err := ai.StoreArtist(artist); err != nil {
		return model.Artist{}, err
	}

	return artist, nil
}

// StoreArtist: Handle use of CsvService/StoreArtist
func (ai *ArtistInteractor) StoreArtist(artist model.Artist) error {
	// Write the response from External API into a csv file.
	if err := ai.CsvService.StoreArtist(artist); err != nil {
		return err
	}

	return nil
}