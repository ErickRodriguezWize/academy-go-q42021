package service

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/ErickRodriguezWize/academy-go-q42021/config"
	spotierr "github.com/ErickRodriguezWize/academy-go-q42021/errors"
	"github.com/ErickRodriguezWize/academy-go-q42021/domain/model"

	"github.com/stretchr/testify/assert"
)

// TestSpotifyService_SearchArtist: Unit Testing
func TestSpotifyService_SearchArtist(t *testing.T) {
	//Disabled Logs from the server.
	log.SetOutput(ioutil.Discard)

		//testcases
		testCases:= []struct{
			name string
			artist string
			responseArtist model.Artist
			errorArtist error
			hasError bool 
		}{
			{
				name:"Found Queens",
				artist:"Queen",
				responseArtist: model.Artist{ID:"6QWuYtzBkQ2Re44gRxaB2e", Name: "Queen", SpotifyURL:"https://open.spotify.com/artist/6QWuYtzBkQ2Re44gRxaB2e", Genres: []string{"classic rock", "glam rock", "rock"}},
				errorArtist: nil,
				hasError: false,
			},
			{
				name:"Found Linkin Park",
				artist:"linkin+park",
				responseArtist: model.Artist{ID:"6XyY86QOPPrYVGvF9ch6wz", Name: "Linkin Park", SpotifyURL:"https://open.spotify.com/artist/6XyY86QOPPrYVGvF9ch6wz", Genres: []string{"alternative metal", "nu metal", "post-grunge", "rap metal"} },
				errorArtist: nil, 
				hasError: false,
			},
			{
				name:"Not Found Artist",
				artist:"quenxious",
				responseArtist: model.Artist{},
				errorArtist:  spotierr.ErrArtistNotFound,
				hasError: true,
			},
			{
				name:"Coma in name",
				artist:"papa,roach",
				responseArtist: model.Artist{},
				errorArtist: spotierr.ErrArtistNotFound, 
				hasError: true,
			},
		}
	
	// Table test cases.
	for _, tsc := range testCases {
		t.Run(tsc.name, func(t *testing.T) {
			// Initialize config struct with environment variables.
			config, _ := config.LoadConfig()
			
			// Init service struct. 
			service := NewSpotifyService(config)
			
			// Execute method. 
			artist, err := service.SearchArtist(tsc.artist)
			
			// Assert
			assert.EqualValues(t, artist, tsc.responseArtist)
			if tsc.hasError{
				assert.EqualError(t, err, tsc.errorArtist.Error())
			}else{
				assert.Nil(t, err)
			}
		})
	}
}
