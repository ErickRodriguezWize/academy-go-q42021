// File use to register services.

package registry

import "github.com/ErickRodriguezWize/academy-go-q42021/service"

// NewPokemonService: Register PokemonService.
func (r *Registry) NewPokemonService() service.PokemonService {
	return service.NewPokemonService()
}

// NewCsvService: Register CsvService
func (r *Registry) NewCsvService() service.CsvService {
	return service.NewCsvService(r.config)
}

// NewSpotifyService: Register Spotifyservice.
func (r *Registry) NewSpotifyService() service.SpotifyService {
	return service.NewSpotifyService(r.config)
}