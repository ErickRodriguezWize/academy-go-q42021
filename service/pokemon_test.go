package service

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/ErickRodriguezWize/academy-go-q42021/domain/model"
)

// TestGetPokemonByID - Testing Function that will simulated scenarios for GetPokemonByID
func TestGetPokemonByID(t *testing.T) {
	//Disabled log ouputs.
	log.SetOutput(ioutil.Discard)

	// test cases.
	tests := map[string]struct {
		input int
		want  string
	}{
		"pokemon found: Charmeleon": {input: 5, want: ""},
		"pokemon found: Charizard":  {input: 6, want: ""},
		"pokemon not found":         {input: 10, want: "Pokemon not Found"},
	}

	mockPokemons := []model.Pokemon{
		{1, "Bulbasaur"},
		{2, "Ivysaur"},
		{3, "Venusaur"},
		{4, "Charmander"},
		{5, "Charmeleon"},
		{6, "Charizard"},
		{8, "Wartortle"},
		{9, "Blastoise"},
	}

	// Table test cases.
	for name, tsc := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := GetPokemonByID(mockPokemons, tsc.input)
			if err != nil {
				if err.Error() != tsc.want {
					t.Fatalf("Pokemon not found with the id: %d", tsc.input)
				}
			}
		})
	}

}