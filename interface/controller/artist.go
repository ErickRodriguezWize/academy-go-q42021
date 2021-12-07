package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ErickRodriguezWize/academy-go-q42021/usecase/interactor"

	"github.com/gorilla/mux"
)

type ArtistController struct {
	ArtistInteractor interactor.ArtistInteractor
}

// SearchArtist: Search an Artist using the artist name.
func (ac *ArtistController) SearchArtist(res http.ResponseWriter, req *http.Request) {
	// Get param of artist from gorilla/mux package.
	artist := mux.Vars(req)["artist"]
	log.Printf("HTTP GET /artists/%v \n", artist)

	// Search for artist using his name.
	targetArtist, err := ac.ArtistInteractor.SearchArtist(artist)
	if err != nil {
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
func NewArtistController(ai interactor.ArtistInteractor) *ArtistController {
	return &ArtistController{ai}
}