/* 
	Service focus on Searching information about Music Artist using Spotify API. 
	DOCs: https://developer.spotify.com/documentation/general/guides/authorization/client-credentials/
*/

package service

import (
	"net/http"
	"net/url"
	"errors"
	"log"
	"io/ioutil"
	"encoding/json"
	"strings"

	"github.com/ErickRodriguezWize/academy-go-q42021/domain/model"
)

//Environment Variables. 
var  (
	/* CONST */ SPOTIFY_API_KEY = GetEnvVariable("SPOTIFY_ACCESS_TOKEN")
	/* CONST */ SPOTIFY_API_DOMAIN = GetEnvVariable("SPOTIFY_API_DOMAIN")
	/* CONST */ LIMIT_ARTIST = GetEnvVariable("LIMIT_ARTIST")

	// Errors
	InvalidToken = errors.New("Spotify Token is invalid")
	BadRequestFormat = errors.New("Bad Request HTTP Format.")
	MissingArtist = errors.New("Couldn't Find the Artist.")
)


// Struct made using JSON to GO Structure Tool: https://mholt.github.io/json-to-go/
// ArtistResponseJSON is a struct that handle the "unmarshall" information from HTTP Response Body. 
type ArtistResponseJSON struct {
	Artists struct{
		Items [] struct {
			Genres []string `json:"genres"`
			SpotifyID string `json:"id"`
			Name string `json:"name"`
			ExternalUrls struct{
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
		} `json:"items"`
	} `json:"artists"`
}

//SearcArtist - Make a HTTP GET call to Spotify API to search for an Artist Information. , targetArtist *model.Artist
func SearchArtist(artist string, targetArtist *model.Artist) (error){
	artist = strings.ReplaceAll(artist, "+", " ")

	//Encoded spaces with "%20". Spotify ask for this escaping for correctly HTTP Call.
	encoded_artist := url.QueryEscape(artist)
	endpoint := SPOTIFY_API_DOMAIN + "search?q=artist:" +encoded_artist + "&type=artist&limit=" + LIMIT_ARTIST
	log.Println("*** HTTP GET to",endpoint)

	//Create request with URL endpoint, Method and Headers. 
	request, err := http.NewRequest("GET", endpoint, nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", SPOTIFY_API_KEY)
	if err != nil{
		return  BadRequestFormat
	}

	//HTTP Client makes the request to the API Endpoint. 
	response, _ := http.DefaultClient.Do(request)
	if response.StatusCode >= 400{
		log.Println("Error: ",response)
		return InvalidToken

	}

	log.Println("response",response)
	
	//Close http.Client Response
	defer response.Body.Close() 

	//Read Response Body from HTTP Call
	body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        return errors.New("Couldn't find anything")
    
    }

    //Unmarshall JSON is use to MAP the values from body json to the Struct ArtistResponseJSON
    //NOTE. Spotify API returns "similar names" artist on the response. 
    jsonResponse := ArtistResponseJSON{}
    if err := json.Unmarshal([]byte(string(body)), &jsonResponse) ; err !=nil {
    	return errors.New("Couldnt process Api Response")
    
    }

    // Search through ITEM array (from Response body) and find the Band Name. 
    for _,itemArtist := range jsonResponse.Artists.Items {
    	log.Println("Artist", itemArtist)
    	if strings.ToLower(itemArtist.Name) ==  strings.ToLower(artist){
    		targetArtist.ID = itemArtist.SpotifyID
    		targetArtist.Name = itemArtist.Name
    		targetArtist.SpotifyURL = itemArtist.ExternalUrls.Spotify
    		//  Genres is a []string
    		// []string... Let me pass multiple values to the func append.
    		targetArtist.Genres = append(targetArtist.Genres, itemArtist.Genres...)
    	}
    }

    if targetArtist.Name == ""{
    	return MissingArtist
    }

	return nil

}