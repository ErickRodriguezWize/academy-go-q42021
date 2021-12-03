package service

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/ErickRodriguezWize/academy-go-q42021/config"
	"github.com/ErickRodriguezWize/academy-go-q42021/domain/model"
)

// TestSearch: Testing Function that will simulated scenarios for SearchArtist
func TestSearchArtist(t *testing.T) {
	//Disabled Logs from the server.
	log.SetOutput(ioutil.Discard)

	// test cases.
	tests := map[string]struct{
		input string
		want string
	}{
		"valid artist": 		{input: "papa+roach", want: ""},
		"artist with coma": 	{input: "linkin,park", want: "Artist not found"},
		"artist with space":	{input: "the beatles", want: ""},
		"artist not found":		{input: "queenshishiux", want: "Artist not found"},
	}

	conf, _ := config.LoadConfig()
	
	// Table test cases. 
	for name, tsc := range tests {
		t.Run(name, func(t *testing.T) {
			emptyArtist := model.Artist{}
			err := SearchArtist(tsc.input, &emptyArtist, *conf)

			if err != nil {
				if err.Error() != tsc.want {
					t.Fatalf("Error: %v", err.Error())
				}
				
			}
		})
	}
}
