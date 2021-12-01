package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ErickRodriguezWize/academy-go-q42021/domain/model"
	"github.com/ErickRodriguezWize/academy-go-q42021/service"

	"github.com/gorilla/mux"
)

type artistController struct {
	Config model.Config
}

// SearchArtist: Search an Artist using the artist name.
func (ac *artistController) SearchArtist(res http.ResponseWriter, req *http.Request) {
	// Get param of artist from gorilla/mux package.
	artist := mux.Vars(req)["artist"]
	log.Printf("HTTP GET /artists/%v \n", artist)

	CsvPath := ac.Config.ArtistCsvPath

	// Search for the artist on the service: spotify.go
	targetArtist := model.Artist{}
	err := service.SearchArtist(artist, &targetArtist, ac.Config)
	if err != nil {
		log.Println("Error: " + err.Error())
		http.Error(res, err.Error(), http.StatusBadRequest)

		return
	}

	// Write the response from External API into a csv file.
	if err := service.WriteArtistIntoCSV(CsvPath, targetArtist); err != nil {
		log.Println("Error: " + err.Error())
		http.Error(res, err.Error(), http.StatusBadRequest)

		return
	}

	// Setup response (headers, http Status)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(res).Encode(targetArtist); err != nil {
		log.Println("Error: " + err.Error())
		http.Error(res, err.Error(), http.StatusBadRequest)

		return
	}

}

// NewArtistController: Returns an empty Struct of artistController.
func NewArtistController(config model.Config) *artistController {
	return &artistController{
		Config: config,
	}
}