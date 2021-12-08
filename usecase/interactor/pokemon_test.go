package interactor

import (
	"testing"

	"github.com/ErickRodriguezWize/academy-go-q42021/domain/model"

	"github.com/stretchr/testify/mock"
)

type mockPokemonInteractor struct {
	mock.Mock
}

var Pokemons = []model.Pokemon{
	{ID:1, Name:"pokemon#1"},
	{ID:2, Name:"pokemon#2"},
	{ID:3, Name:"pokemon#3"},
	{ID:4, Name:"pokemon#4"},
	{ID:5, Name:"pokemon#5"},
	{ID:6, Name:"pokemon#6"},
	{ID:7, Name:"pokemon#7"},
	{ID:8, Name:"pokemon#8"},
	{ID:9, Name:"pokemon#9"},
	{ID:10, Name:"pokemon#10"},
	{ID:11, Name:"pokemon#11"},
	{ID:12, Name:"pokemon#12"},
	{ID:13, Name:"pokemon#13"},
	{ID:14, Name:"pokemon#14"},
	{ID:15, Name:"pokemon#15"},
	{ID:16, Name:"pokemon#16"},
	{ID:17, Name:"pokemon#17"},
	{ID:18, Name:"pokemon#18"},
	{ID:19, Name:"pokemon#19"},
	{ID:20, Name:"pokemon#20"},
	{ID:21, Name:"pokemon#21"},
	{ID:22, Name:"pokemon#22"},
	{ID:23, Name:"pokemon#23"},
	{ID:24, Name:"pokemon#24"},
	{ID:25, Name:"pokemon#25"},
	{ID:26, Name:"pokemon#26"},
	{ID:27, Name:"pokemon#27"},
	{ID:28, Name:"pokemon#28"},
	{ID:29, Name:"pokemon#29"},
	{ID:30, Name:"pokemon#30"},
}

func (mr mockPokemonInteractor) GetAllPokemons()([]model.Pokemon, error){
	arg:= mr.Called()
	return arg.Get(0).([]model.Pokemon), arg.Error(1)
}

func (mr mockPokemonInteractor)GetPokemon(ID int) (model.Pokemon, error){
	arg:= mr.Called()
	return arg.Get(0).([]model.Pokemon), arg.Error(1)
}



func TestPokemonInteractor_GetAllPokemons(t *testing.T){
	testCases:= []struct{
		name string
		expectedLength int 
		response []model.Pokemon,
		error error
	}{
		{"Retrieve Correctly Pokemons", 30, characters, nil},
		{"Wrong lenght of pokemons", 30, characters, nil},
	}

	for _,tc := range testCases {
		t.Run(tc.name, func(t *testing.T){
			mock:= mockPokemonInteractor{}
			mock.On("GetAllPokemons").Return(tc.response, tc.error)

			iteractor
		})
	}
}

func TestPokemonInteractor_GetPokemon(t *testing.T){

}

func TestPokemonInteractor_GetPokemonsWorker(t *testing.T){

}