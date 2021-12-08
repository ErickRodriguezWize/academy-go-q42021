package interactor

import (
	"testing"

	pokerror "github.com/ErickRodriguezWize/academy-go-q42021/errors"
	"github.com/ErickRodriguezWize/academy-go-q42021/domain/model"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/assert"
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
	{ID:31, Name:"pokemon#31"},
	{ID:32, Name:"pokemon#32"},
	{ID:33, Name:"pokemon#33"},
	{ID:34, Name:"pokemon#34"},
	{ID:35, Name:"pokemon#35"},
	{ID:36, Name:"pokemon#36"},
	{ID:37, Name:"pokemon#37"},
	{ID:38, Name:"pokemon#38"},
	{ID:39, Name:"pokemon#39"},
	{ID:40, Name:"pokemon#40"},
	{ID:41, Name:"pokemon#41"},
	{ID:42, Name:"pokemon#42"},
	{ID:43, Name:"pokemon#43"},
	{ID:44, Name:"pokemon#44"},
	{ID:45, Name:"pokemon#45"},
	{ID:46, Name:"pokemon#46"},
	{ID:47, Name:"pokemon#47"},
	{ID:48, Name:"pokemon#48"},
	{ID:49, Name:"pokemon#49"},
	{ID:50, Name:"pokemon#50"},
	{ID:51, Name:"pokemon#51"},
	{ID:52, Name:"pokemon#52"},
	{ID:53, Name:"pokemon#53"},
	{ID:54, Name:"pokemon#54"},
	{ID:55, Name:"pokemon#55"},
	{ID:56, Name:"pokemon#56"},
	{ID:57, Name:"pokemon#57"},
	{ID:58, Name:"pokemon#58"},
	{ID:59, Name:"pokemon#59"},
	{ID:60, Name:"pokemon#60"},
}

// Create mock for PokemonService 
type mockPokemonService struct {
	mock.Mock
}

// GetPokemonById: mock method of GetPokemonByID
func (mp mockPokemonService) GetPokemonByID(pokemons []model.Pokemon, ID int) (model.Pokemon, error){
	arg:= mp.Called()
	return arg.Get(0).(model.Pokemon), arg.Error(1)
}

// GetPokemonById: mock method of GetPokemonWorker
func (mp mockPokemonService) GetPokemonWorker(pokemons []model.Pokemon, t string, items int, itemsPerWorker int) ([]model.Pokemon, error){
	arg:= mp.Called()
	return arg.Get(0).([]model.Pokemon), arg.Error(1)
}

// Mock struct of FileService
type mockReadService struct {
	mock.Mock
}

// GetPokemonById: mock method of ReadAll
func (mf mockReadService) ReadAll() ([]model.Pokemon, error) {
	arg:= mf.Called()
	return arg.Get(0).([]model.Pokemon), arg.Error(1)
}

// TestPokemonInteractor_GetAllPokemons Unit testing. 
func TestPokemonInteractor_GetAllPokemons(t *testing.T){
	// Test Cases.
	testCases:= []struct{
		name string
		expectedLength int 
		response []model.Pokemon
		error error
		hasError bool 
	}{
		{name: "GetAllPokemon: OK", 
		expectedLength: 60, 
		response: Pokemons, 
		error: nil, 
		hasError: false},
	}

	for _,tc := range testCases {
		// Run Test scenario. 
		t.Run(tc.name, func(t *testing.T){
			// Init Mocks. 
			mockPokemon := mockPokemonService{}
			mockFile:= mockReadService{}
			mockFile.On("ReadAll").Return(tc.response, tc.error)

			// Implement mocks. 
			service := NewPokemonInteractor(mockPokemon, mockFile)

			// Execute target method. 
			result, err := service.GetAllPokemons()

			// Asserts results. 
			assert.EqualValues(t, tc.expectedLength, len(result))
			if tc.hasError{
				assert.EqualError(t, err, tc.error.Error())
			}else{
				assert.Nil(t, err)
			}

		})
	}
}

// TestPokemonInteractor_GetPokemon: Unit Testing. 
func TestPokemonInteractor_GetPokemon(t *testing.T){
	// test scenarios. 
	testCases:= []struct{
		name string
		expectedLength int 
		response []model.Pokemon
		id int
		responsePokemon model.Pokemon
		error error
		hasError bool 
	}{
		{
			name: "Found pokemon with ID 5", 
			expectedLength: 30, 
			response: Pokemons, 
			id: 5,
			responsePokemon: model.Pokemon{ID:5, Name:"pokemon#5"},
			error: nil, 
			hasError: false,
		},
		{
			name: "Found pokemon with ID 12", 
			expectedLength: 30, 
			response: Pokemons, 
			id: 12,
			responsePokemon: model.Pokemon{ID:5, Name:"pokemon#5"},
			error: nil, 
			hasError: false,
		},
		{
			name: "Found pokemon with ID 23", 
			expectedLength: 30, 
			response: Pokemons, 
			id: 23,
			responsePokemon: model.Pokemon{ID:5, Name:"pokemon#5"},
			error: nil, 
			hasError: false,
		},
		{
			name: "Pokemon Not Found", 
			expectedLength: 30, 
			response: Pokemons, 
			id: 80,
			responsePokemon: model.Pokemon{},
			error: pokerror.ErrPokemonNotFound, 
			hasError: true,
		},
	}

	for _,tc := range testCases {
		// Run tests. 
		t.Run(tc.name, func(t *testing.T){
			// Init Mocks. 
			mockPokemon := mockPokemonService{}
			mockPokemon.On("GetPokemonByID").Return(tc.responsePokemon, tc.error)
			mockFile:= mockReadService{}
			mockFile.On("ReadAll").Return(tc.response, tc.error)

			// Implements Mocks. 
			service := NewPokemonInteractor(mockPokemon, mockFile)

			// Execute method. 
			pokemon, err := service.GetPokemon(tc.id)
			
			// Asserts. 
			assert.EqualValues(t, tc.responsePokemon, pokemon)
			if tc.hasError{
				assert.EqualError(t, err, tc.error.Error())
			}else{
				assert.Nil(t, err)
			}

		})
	}
}

// TestPokemonInteractor_GetPokemonWorker: Unit testing for the method GetPokemonWorker
func TestPokemonInteractor_GetPokemonWorker(t *testing.T){
	// create TestCases.
	testCases:= []struct{
		name string
		expectedLength int
		responseRead []model.Pokemon
		errorRead error
		t string
		items int 
		itemsWorker int 
		responseWorker []model.Pokemon
		errorWorker error
		hasError bool 
	}{
		{
			name: "PokemonWorker: Type odd", 
			expectedLength: 20,
			responseRead: Pokemons, 
			errorRead: nil,
			t: "odd",
			items : 20,
			itemsWorker: 5,
			responseWorker:  Pokemons[:20],
			errorWorker: nil, 
			hasError: false,
		},
		{
			name: "PokemonWorker: Type even", 
			expectedLength: 20,
			responseRead: Pokemons, 
			errorRead: nil,
			t: "even",
			items : 20,
			itemsWorker: 5,
			responseWorker:  Pokemons[1:21],
			errorWorker: nil, 
			hasError: false,
		},
	}

	// For loop to test each testCase. 
	for _,tc := range testCases {
		t.Run(tc.name, func(t *testing.T){
			// Init Mocks.
			mockPokemon := mockPokemonService{}
			mockPokemon.On("GetPokemonWorker").Return(tc.responseWorker, tc.errorWorker)
			mockFile:= mockReadService{}
			mockFile.On("ReadAll").Return(tc.responseRead, tc.errorRead)

			// Implement interfaces using mocks. 
			service := NewPokemonInteractor(mockPokemon, mockFile)

			// Test GetPokemon method. 
			pokemon, _ := service.GetPokemonWorker(tc.t, tc.items, tc.itemsWorker)
			assert.EqualValues(t, tc.expectedLength, len(pokemon))

			if tc.t == "even"{
				assert.EqualValues(t, (pokemon[0].ID % 2) , 0)
			}else{
				assert.EqualValues(t, (pokemon[0].ID % 2) , 1)
			}

		})
	}
}