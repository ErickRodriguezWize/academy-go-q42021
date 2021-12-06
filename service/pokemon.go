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
	retrievePokemons:=  make([]model.Pokemon,0)

	// Workers
	workers :=  8 //(items/itemsPerWorker) + 1
	// Number of jobs. 
	numJobs:=len(pokemons)

	// Define channels for communication 
	jobs := make(chan model.Pokemon, numJobs)
	results := make(chan model.Pokemon, numJobs)
	done := make(chan bool, workers) // channel to end working pool. 

	// Create and run Workers to handle jobs. 
	log.Println(" ******* Started Worker pool with", workers, "Workers *********")
	wg.Add(workers)
	go func (){
		for w:=1 ; w <= workers; w++ {
			go createWorker(w, jobs, results, &wg, t, done, itemsPerWorker)
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
			// Send signals to end to every worker. 
			done <- true
			break
		}
		// Obtain results from Worker pool. 
		r:= <- results
		if( r != model.Pokemon{}) {
			// Append result to slice (return value of method).
			retrievePokemons =  append(retrievePokemons, r)
		}
	}


	// wait for all Workers to finish execution to log and return results. 
	wg.Wait() 
	close(done)
	time.Sleep(time.Second)
	log.Printf("Results #: %v\n", len(retrievePokemons))
	return retrievePokemons
}

// createWorker: Handle Worker logic and values to result channel. 
func createWorker(id int, jobs <-chan model.Pokemon, results chan<- model.Pokemon, wg *sync.WaitGroup, t string, done chan bool, items int) {
	count := 0
	// range between channel job.
	for j := range jobs {
		log.Println("worker", id, "started  job: ", j)
		
		select{
			case <-done: // Signal Worker to stop. 
				wg.Done() 
				log.Println("****** SHUTDOWN WORKER #", id, "With", count, " Jobs Done")
				return
			default:
						//Worker processing the Job
		time.Sleep(time.Second)
		switch t {
			// in case type is equal to odd. 
			case "odd":
				if (j.ID % 2) == 1 {
					count = count + 1
					results <- j
					continue
				}else{
					results <- model.Pokemon{}
				}
			case "even": // In case type is equal to even. 
				
				if (j.ID % 2) == 0 {
					count = count + 1
					results <- j
				}else{
					results <- model.Pokemon{}
				}
			default:
		}

		log.Println("worker", id, "finishs  job:", j)

		// If worker reach its limits jobs done, it will close. 
		if count == items {
			wg.Done() 
			log.Println("****** Worker #", id, "Reach limit of ", items, " Jobs Done")
			return
		}
		}


	}
	// Close worker in case the range of jobs finish. 
	wg.Done()
	log.Println("****** FINISH WORKER #", id, "With", count, " Jobs Done")
}

// NewPokemonService: Constructor for PokemonService struct. 
func NewPokemonService() PokemonService {
	return PokemonService{}
}