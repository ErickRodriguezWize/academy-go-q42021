package controller

import(
	"log"
	"net/http"
	"encoding/json"

	"github.com/ErickRodriguezWize/academy-go-q42021/service"
	"github.com/ErickRodriguezWize/academy-go-q42021/domain/model"

	"github.com/gorilla/mux"
)

type artistController struct{}

type ArtistController interface{
	SearchArtist(res http.ResponseWriter, req *http.Request)
}

func (mv *artistController) SearchArtist(res http.ResponseWriter, req *http.Request){
	artist := mux.Vars(req)["artist"]
	log.Printf("HTTP GET /artists/%v \n",artist)

	targetArtist := model.Artist{}
	err := service.SearchArtist(artist, &targetArtist)
	if err != nil {
		log.Println("Error: " + err.Error() )
		http.Error(res, err.Error(), http.StatusBadRequest)

		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(targetArtist)

}

func NewArtistController() (ArtistController){
	return &artistController{}
}