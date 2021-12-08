package service

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/ErickRodriguezWize/academy-go-q42021/domain/model"
	pokerror "github.com/ErickRodriguezWize/academy-go-q42021/errors"

	"github.com/stretchr/testify/assert"
)

var pokemons = []model.Pokemon{
	{ID: 1, Name: "pokemon#1"},
	{ID: 2, Name: "pokemon#2"},
	{ID: 3, Name: "pokemon#3"},
	{ID: 4, Name: "pokemon#4"},
	{ID: 5, Name: "pokemon#5"},
	{ID: 6, Name: "pokemon#6"},
	{ID: 7, Name: "pokemon#7"},
	{ID: 8, Name: "pokemon#8"},
	{ID: 9, Name: "pokemon#9"},
	{ID: 10, Name: "pokemon#10"},
	{ID: 11, Name: "pokemon#11"},
	{ID: 12, Name: "pokemon#12"},
	{ID: 13, Name: "pokemon#13"},
	{ID: 14, Name: "pokemon#14"},
	{ID: 15, Name: "pokemon#15"},
	{ID: 16, Name: "pokemon#16"},
	{ID: 17, Name: "pokemon#17"},
	{ID: 18, Name: "pokemon#18"},
	{ID: 19, Name: "pokemon#19"},
	{ID: 20, Name: "pokemon#20"},
	{ID: 21, Name: "pokemon#21"},
	{ID: 22, Name: "pokemon#22"},
	{ID: 23, Name: "pokemon#23"},
	{ID: 24, Name: "pokemon#24"},
	{ID: 25, Name: "pokemon#25"},
	{ID: 26, Name: "pokemon#26"},
	{ID: 27, Name: "pokemon#27"},
	{ID: 28, Name: "pokemon#28"},
	{ID: 29, Name: "pokemon#29"},
	{ID: 30, Name: "pokemon#30"},
	{ID: 31, Name: "pokemon#31"},
	{ID: 32, Name: "pokemon#32"},
	{ID: 33, Name: "pokemon#33"},
	{ID: 34, Name: "pokemon#34"},
	{ID: 35, Name: "pokemon#35"},
	{ID: 36, Name: "pokemon#36"},
	{ID: 37, Name: "pokemon#37"},
	{ID: 38, Name: "pokemon#38"},
	{ID: 39, Name: "pokemon#39"},
	{ID: 40, Name: "pokemon#40"},
	{ID: 41, Name: "pokemon#41"},
	{ID: 42, Name: "pokemon#42"},
	{ID: 43, Name: "pokemon#43"},
	{ID: 44, Name: "pokemon#44"},
	{ID: 45, Name: "pokemon#45"},
	{ID: 46, Name: "pokemon#46"},
	{ID: 47, Name: "pokemon#47"},
	{ID: 48, Name: "pokemon#48"},
	{ID: 49, Name: "pokemon#49"},
	{ID: 50, Name: "pokemon#50"},
	{ID: 51, Name: "pokemon#51"},
	{ID: 52, Name: "pokemon#52"},
	{ID: 53, Name: "pokemon#53"},
	{ID: 54, Name: "pokemon#54"},
	{ID: 55, Name: "pokemon#55"},
	{ID: 56, Name: "pokemon#56"},
	{ID: 57, Name: "pokemon#57"},
	{ID: 58, Name: "pokemon#58"},
	{ID: 59, Name: "pokemon#59"},
	{ID: 60, Name: "pokemon#60"},
}

// TestPokemonService_GetPokemonByID: Unit Testing.
func TestPokemonService_GetPokemonByID(t *testing.T) {
	//Disabled log ouputs.
	log.SetOutput(ioutil.Discard)

	// test scenarios.
	testCases := []struct {
		name            string
		id              int
		responsePokemon model.Pokemon
		errorPokemon    error
		hasError        bool
	}{
		{
			name:            "Found pokemon with ID 5",
			id:              5,
			responsePokemon: model.Pokemon{ID: 5, Name: "pokemon#5"},
			errorPokemon:    nil,
			hasError:        false,
		},
		{
			name:            "Found pokemon with ID 12",
			id:              12,
			responsePokemon: model.Pokemon{ID: 12, Name: "pokemon#12"},
			errorPokemon:    nil,
			hasError:        false,
		},
		{
			name:            "Found pokemon with ID 23",
			id:              23,
			responsePokemon: model.Pokemon{ID: 23, Name: "pokemon#23"},
			errorPokemon:    nil,
			hasError:        false,
		},
		{
			name:            "Pokemon Not Found",
			id:              80,
			responsePokemon: model.Pokemon{},
			errorPokemon:    pokerror.ErrPokemonNotFound,
			hasError:        true,
		},
	}

	// Table test cases.
	for _, tsc := range testCases {
		t.Run(tsc.name, func(t *testing.T) {
			// Init service struct.
			service := NewPokemonService()

			// Execute method.
			result, err := service.GetPokemonByID(pokemons, tsc.id)

			// Assert
			assert.EqualValues(t, result, tsc.responsePokemon)
			if tsc.hasError {
				assert.EqualError(t, err, tsc.errorPokemon.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

// TestPokemonService_GetPokemonWorker: Unit Test.
func TestPokemonService_GetPokemonWorker(t *testing.T) {
	//Disabled log ouputs.
	log.SetOutput(ioutil.Discard)

	// create TestCases.
	testCases := []struct {
		name           string
		expectedLength int
		t              string
		items          int
		itemsWorker    int
	}{
		{
			name:           "PokemonWorker: Type odd",
			expectedLength: 25,
			t:              "odd",
			items:          25,
			itemsWorker:    5,
		},
		{
			name:           "PokemonWorker: Type even",
			expectedLength: 20,
			t:              "even",
			items:          20,
			itemsWorker:    5,
		},
	}

	// Table cases.
	for _, tsc := range testCases {
		//Start testing.
		t.Run(tsc.name, func(t *testing.T) {
			// Init Service struct.
			service := NewPokemonService()

			// Execute method.
			result, err := service.GetPokemonWorker(pokemons, tsc.t, tsc.items, tsc.itemsWorker)

			// Asserts
			assert.EqualValues(t, tsc.expectedLength, len(result))
			assert.Nil(t, err)
			if tsc.t == "even" {
				assert.EqualValues(t, (result[0].ID % 2), 0)
			} else {
				assert.EqualValues(t, (result[0].ID % 2), 1)
			}
		})
	}

}