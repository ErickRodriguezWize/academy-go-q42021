package service

import (
	"log"
	"sync"
	"time"

	"github.com/ErickRodriguezWize/academy-go-q42021/domain/model"
	pokerr "github.com/ErickRodriguezWize/academy-go-q42021/errors"
)

type PokemonService struct{}

// GetPokemonById: finds a pokemon inside pokemons array using the ID as filter.
func (sp PokemonService) GetPokemonByID(pokemons []model.Pokemon, ID int) (model.Pokemon, error) {
	targetPkm := model.Pokemon{}

	// Loop through the pokemosn to find the pokemon with the correct ID
	for _, pkm := range pokemons {
		if pkm.ID == ID {
			return pkm, nil
		}
	}

	return targetPkm, pokerr.ErrPokemonNotFound

}

// GetPokemonWorker: Get Pokemons based on ID (odd or even) using Workers pools. 
func (sp PokemonService)GetPokemonWorker(pokemons []model.Pokemon, t string, items int, itemsPerWorker int) ([]model.Pokemon){
	var wg sync.WaitGroup 


	var retrievePokemons []model.Pokemon
	workers := (items/itemsPerWorker) + 1
	numJobs:=len(pokemons)

	// Define channels for communication 
	jobs := make(chan model.Pokemon, numJobs)
	results := make(chan model.Pokemon, numJobs)
	done := make(chan bool)

	// Create n Workers to handle jobs. 
	wg.Add(workers)
	go func (){
		for w:=1 ; w <= workers; w++ {
			go createWorker(w, jobs, results, &wg, t, done)
		}

	}()


	// Send jobs to Workers
	for i:= 0 ; i < numJobs; i++{
		jobs <- pokemons[i]
	}
	close(jobs)

	// Receive results from workers.
	
	for j:=0; j< numJobs; j++{
		// Check if we retrieve the target item number. 
		if len(retrievePokemons) == items {
 			break
		}
		r:= <- results

		// Check if the results is not empty. 
		if (r != model.Pokemon{}){
			retrievePokemons =  append(retrievePokemons, r)
		}
	}

	wg.Wait()
	log.Printf("Results #: %v\n", len(retrievePokemons))
	return retrievePokemons
}

// createWorker: Handle Worker logic and values to result channel. 
func createWorker(id int, jobs <-chan model.Pokemon, results chan<- model.Pokemon, wg *sync.WaitGroup, t string, done chan bool) {
	wg.Done()
	for j := range jobs {
		log.Println("worker", id, "started  job: ", j)
		select{
			case <-done:
				return
			default:
		}

		time.Sleep(time.Second)
		switch t {
		case "odd":
			if (j.ID % 2) == 1 {
				results <- j
				continue
			}else{
				results <- model.Pokemon{}
			}
		case "even":
			if (j.ID % 2) == 0 {
				results <- j
			}else{
				results <- model.Pokemon{}
			}
		default:
		}

		log.Println("worker", id, "finishs  job:", j)
	}
}

// NewPokemonService: Constructor for PokemonService struct. 
func NewPokemonService() PokemonService {
	return PokemonService{}
}