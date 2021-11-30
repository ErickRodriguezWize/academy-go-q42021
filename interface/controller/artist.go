package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ErickRodriguezWize/academy-go-q42021/domain/model"
	"github.com/ErickRodriguezWize/academy-go-q42021/service"

	"github.com/gorilla/mux"
)

var path = service.GetEnvVariable("CSV_ARTIST_PATH")

type artistController struct{}

type ArtistController interface {
	SearchArtist(res http.ResponseWriter, req *http.Request)
}

// SearchArtist: Search an Artist using the artist name.
func (mv *artistController) SearchArtist(res http.ResponseWriter, req *http.Request) {
	//Get param of artist from gorilla/mux package.
	artist := mux.Vars(req)["artist"]
	log.Printf("HTTP GET /artists/%v \n", artist)

	//Search for the artist on the service: spotify.go
	targetArtist := model.Artist{}
	err := service.SearchArtist(artist, &targetArtist)
	if err != nil {
		log.Println("Error: " + err.Error())
		http.Error(res, err.Error(), http.StatusBadRequest)

		return
	}

	//Write the response from External API into a csv file.
	if err := service.WriteArtistIntoCSV(path, targetArtist); err != nil {
		log.Println("Error: " + err.Error())
		http.Error(res, err.Error(), http.StatusBadRequest)

		return
	}

	//Setup response (headers, http Status)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(res).Encode(targetArtist); err != nil {
		log.Println("Error: " + err.Error())
		http.Error(res, err.Error(), http.StatusBadRequest)

		return
	}

}

//NewArtistController: Returns an empty Struct of artistController.
func NewArtistController() ArtistController {
	return &artistController{}
}
