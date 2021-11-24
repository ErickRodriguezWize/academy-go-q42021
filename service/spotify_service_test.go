package service

import (
	"fmt"
	"io/ioutil"
	"log"
	"testing"

	"github.com/ErickRodriguezWize/academy-go-q42021/domain/model"
)

// TestSearch: Testing Function that will simulated scenarios for SearchArtist
func TestSearchArtist(t *testing.T) {
	//Disabled Logs from the server.
	log.SetOutput(ioutil.Discard)

	//Define cases for testing Scenarios
	artistCases := []string{"eminem", "daft+punk", "papa+roach"}

	//t.Run for test scenarios for SearchArtist Func.
	for _, artist := range artistCases {
		t.Run(fmt.Sprintf("[TEST] Search Artist: %v ", artist), func(t *testing.T) {
			emptyArtist := model.Artist{}
			err := SearchArtist(artist, &emptyArtist)
			if err != nil {
				t.Errorf("Error: %v", err.Error())
			}
		})
	}
}
