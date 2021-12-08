package service

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/ErickRodriguezWize/academy-go-q42021/config"
	filerror "github.com/ErickRodriguezWize/academy-go-q42021/errors"
	"github.com/ErickRodriguezWize/academy-go-q42021/domain/model"

	"github.com/stretchr/testify/assert"
)

func TestFileService_Read(t *testing.T){
	//Disabled log ouputs.
	log.SetOutput(ioutil.Discard)

	//testcases
	testsCases:= []struct{
		name string
		expectedLength int
		dummyPath string
		error error
		hasError bool
	}{
		{
			name:"Retrieve all Pokemons", 
			expectedLength: 150,
			dummyPath: "",
			error: nil,
			hasError: false,
		},
		{
			name:"Couldn't Open File", 
			expectedLength: 0,
			dummyPath: "dummy_path_to_file.csv",
			error: filerror.ErrFileError,
			hasError: true,
		},
	}

	for _,tsc := range testsCases {
		// Start test. 
		t.Run(tsc.name, func(t *testing.T){
			// Initialize config struct with environment variables.
			config, _ := config.LoadConfig()
			log.Println("Path", config.ArtistCsvPath)
			// Modify csv path to trigger specific error. 
			if tsc.dummyPath != ""{
				config.PokemonCsvPath = tsc.dummyPath
			}else{
				config.PokemonCsvPath = "../" + config.PokemonCsvPath
			}

			// Init FileService
			service := NewFileService(config)
			
			// Executed method. 
			results, err := service.ReadAll()

			//assert
			assert.EqualValues(t, tsc.expectedLength, len(results))
			if tsc.hasError{
				assert.EqualValues(t, tsc.error.Error(), err.Error())
			}else{
				assert.Nil(t, err)
			}
		})
	}
}

// TestFileService_Write: Unit testing. 
func TestFileService_Write(t *testing.T){
	//Disabled log ouputs.
	log.SetOutput(ioutil.Discard)

	//testcases
	testsCases:= []struct{
		name string
		input model.Artist
	}{
		{
			name:"Writing on CSV Queen", 
			input: model.Artist{ID:"6QWuYtzBkQ2Re44gRxaB2e", Name: "Queen", SpotifyURL:"https://open.spotify.com/artist/6QWuYtzBkQ2Re44gRxaB2e", Genres: []string{"classic rock", "glam rock", "rock"}},
		},
		{
			name:"Empty Write on CSV LinkinPark", 
			input: model.Artist{ID:"6XyY86QOPPrYVGvF9ch6wz", Name: "Linkin Park", SpotifyURL:"https://open.spotify.com/artist/6XyY86QOPPrYVGvF9ch6wz", Genres: []string{"alternative metal", "nu metal", "post-grunge", "rap metal"} },
		},
	}

	for _,tsc := range testsCases {
		// Start test. 
		t.Run(tsc.name, func(t *testing.T){
			// Initialize config struct with environment variables.
			config, _ := config.LoadConfig()
			log.Println("Path", config.ArtistCsvPath)
			// Modify csv path to trigger specific error. 
			config.ArtistCsvPath = "../" + config.ArtistCsvPath

			// Init FileService
			service := NewFileService(config)
			
			// Executed method. 
			err := service.Write(tsc.input)
			
			//assert
			assert.Nil(t, err)
		})
	}
}