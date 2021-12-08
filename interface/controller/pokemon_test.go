package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ErickRodriguezWize/academy-go-q42021/domain/model"
	pokerr "github.com/ErickRodriguezWize/academy-go-q42021/errors"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

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

type mockPokemonInteractor struct{
	mock.Mock
}

func (mp mockPokemonInteractor) GetAllPokemons() ([]model.Pokemon, error){
	arg:= mp.Called()
	return arg.Get(0).([]model.Pokemon), arg.Error(1)
}

func(mp mockPokemonInteractor) GetPokemon(ID int) (model.Pokemon, error){
	arg:= mp.Called()
	return arg.Get(0).(model.Pokemon), arg.Error(1)
}

func (mp mockPokemonInteractor)GetPokemonWorker(t string, items int, itemsPerWorker int) ([]model.Pokemon, error){
	arg:= mp.Called()
	return arg.Get(0).([]model.Pokemon), arg.Error(1)
}

// TestPokemonController_GetAllPokemons: unit
func TestPokemonController_GetAllPokemons(t *testing.T){
	//testcases
	testscases := []struct {
		name string 
		handler   string
		method    string
		endpoint  string
		response []model.Pokemon
		error error
		want_code int
	}{
		{
			name:"GetAllPokemons OK",
			method: "GET",
			endpoint: "/pokemons",
			response: Pokemons,
			error: nil,
			want_code: 200,
		},
		{
			name:"CSV Error",
			method: "GET",
			endpoint: "/pokemons",
			response: []model.Pokemon{},
			error: pokerr.ErrFileError ,
			want_code: 400,
		},
	}
	
	for _,tsc := range testscases{
		// Start Test.
		t.Run(tsc.name, func(t *testing.T){
			// Init mocks. 
			mockInteractor := mockPokemonInteractor{}
			mockInteractor.On("GetAllPokemons").Return(tsc.response, tsc.error)

			// Dummy res http.ResponseWriter and req *http.Request
			req, _ := http.NewRequest(tsc.method, tsc.endpoint, nil)
			res := httptest.NewRecorder()

			// Implemented mocks
			service := NewPokemonController(mockInteractor)

			// Executed methods
			service.GetAllPokemons(res,req)
			
			// Asserts
			assert.EqualValues(t, res.Code, tsc.want_code)
		})
	}
}

// TestPokemonController_GetPokemon: Unit tEST.
func TestPokemonController_GetPokemon(t *testing.T){
	//testcases
	testscases := []struct {
		name string 
		handler   string
		method    string
		endpoint  string
		response []model.Pokemon
		error error
		want_code int
	}{
		{
			name:"GetAllPokemons OK",
			method: "GET",
			endpoint: "/pokemons/30",
			response: Pokemons,
			error: nil,
			want_code: 200,
		},
		{
			name:"GetAllPokemons OK",
			method: "GET",
			endpoint: "/pokemons/10",
			response: Pokemons,
			error: nil,
			want_code: 200,
		},
		{
			name:"Parse Error",
			method: "GET",
			endpoint: "/pokemons/seis",
			response: []model.Pokemon{},
			error: pokerr.ErrParseError,
			want_code: 400,
		},
		{
			name:"PokemonNotFound",
			method: "GET",
			endpoint: "/pokemons/190",
			response: []model.Pokemon{},
			error: pokerr.ErrPokemonNotFound ,
			want_code: 400,
		},
	}
	
	for _,tsc := range testscases{
		// Start Test.
		t.Run(tsc.name, func(t *testing.T){
			// Init mocks. 
			mockInteractor := mockPokemonInteractor{}
			mockInteractor.On("GetAllPokemons").Return(tsc.response, tsc.error)
			mockInteractor.On("GetPokemon").Return(model.Pokemon{}, tsc.error)

			// Dummy res http.ResponseWriter and req *http.Request
			req, _ := http.NewRequest(tsc.method, tsc.endpoint, nil)
			res := httptest.NewRecorder()

			// Implemented mocks
			service := NewPokemonController(mockInteractor)

			// Executed method
			service.GetAllPokemons(res,req)
			
			// Asserts
			assert.EqualValues(t, res.Code, tsc.want_code)
		})
	}
}

// TestPokemonController_GetPokemonsWorker: Unit tEST.
func TestPokemonController_GetPokemonsWorker(t *testing.T){
	//testcases
	testscases := []struct {
		name string 
		method    string
		endpoint  string
		response []model.Pokemon
		error error
		want_code int
	}{
		{
			name: "PokemonWorker Even OK",
			method: "GET",
			endpoint: "/pokemons/worker?type=even&items=25&items_per_worker=5",
			response: Pokemons,
			error: nil,
			want_code: 200,
		},
		{
			name:"PokemonWorker Odd OK",
			method: "GET",
			endpoint: "/pokemons/worker?type=odd&items=25&items_per_worker=5",
			response: Pokemons,
			error: nil,
			want_code: 200,
		},
		{
			name:"PokemonWorker Invalid Type",
			method: "GET",
			endpoint: "/pokemons/worker?type=odding&items=25&items_per_worker=5",
			response: Pokemons,
			error: nil,
			want_code: 400,
		},
		{
			name:"PokemonWorker Empty Type",
			method: "GET",
			endpoint: "/pokemons/worker?type=&items=25&items_per_worker=5",
			response: Pokemons,
			error: nil,
			want_code: 400,
		},
		{
			name:"PokemonWorker Invalid item",
			method: "GET",
			endpoint: "/pokemons/worker?type=&items=six&items_per_worker=5",
			response: Pokemons,
			error: nil,
			want_code: 400,
		},
	}
	
	for _,tsc := range testscases{
		// Start Test.
		t.Run(tsc.name, func(t *testing.T){
			// Init mocks. 
			mockInteractor := mockPokemonInteractor{}
			mockInteractor.On("GetAllPokemons").Return(tsc.response, tsc.error)
			mockInteractor.On("GetPokemonWorker").Return(tsc.response, tsc.error)

			// Dummy res http.ResponseWriter and req *http.Request
			req, _ := http.NewRequest(tsc.method, tsc.endpoint, nil)
			res := httptest.NewRecorder()

			// Implemented mocks
			service := NewPokemonController(mockInteractor)

			// Executed method
			service.GetPokemonsWorker(res,req)
			
			// Asserts
			assert.EqualValues(t, res.Code, tsc.want_code)
		})
	}
}