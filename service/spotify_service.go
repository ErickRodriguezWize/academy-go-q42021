/*
	Service focus on Searching information about Music Artist using Spotify API.
	DOCs: https://developer.spotify.com/documentation/general/guides/authorization/client-credentials/
*/

package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/ErickRodriguezWize/academy-go-q42021/domain/model"
	spotiferr "github.com/ErickRodriguezWize/academy-go-q42021/errors"
)

var (
	// Environment Variables.
	/* CONST */ SpotifyEndpoint = GetEnvVariable("SPOTIFY_API_ENDPOINT")
	/* CONST */ SpotifyRefreshEndpoint = GetEnvVariable("SPOTIFY_REFRESH_ENDPOINT")
	/* CONST */ SpotifyAuhtorizationToken = GetEnvVariable("SPOTIFY_AUTHORIZATION_TOKEN")
	/* CONST */ SpotifyRefreshToken = GetEnvVariable("SPOTIFY_REFRESH_TOKEN")
	/* CONST */ LimitArtist = GetEnvVariable("LIMIT_ARTIST")
)

// Struct made using JSON to GO Structure Tool: https://mholt.github.io/json-to-go/
// ArtistResponseJSON is a struct that handle the "unmarshall" information from HTTP Response Body.
type ArtistResponseJSON struct {
	Artists struct {
		Items []struct {
			Genres       []string `json:"genres"`
			SpotifyID    string   `json:"id"`
			Name         string   `json:"name"`
			ExternalUrls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
		} `json:"items"`
	} `json:"artists"`
}

//SearcArtist: Makes a HTTP GET call to Spotify API to search for an Artist Information.
// This func will return an error (in case that one is triggered).
func SearchArtist(artist string, targetArtist *model.Artist) error {
	//Get token from Spotify API (Token experies every 30 minutes.)
	accessToken, err := RefreshToken()
	if err != nil {
		return spotiferr.TokenMissing
	}

	artist = strings.ReplaceAll(artist, "+", " ")

	//Encoded spaces with "%20". Spotify ask for this escaping for correctly HTTP Calls.
	encoded_artist := url.QueryEscape(artist)
	endpoint := SpotifyEndpoint + "search?q=artist:" + encoded_artist + "&type=artist&limit=" + LimitArtist
	log.Println("*** HTTP GET to", endpoint)

	//Create request with URL endpoint, Method and Headers.
	request, err := http.NewRequest("GET", endpoint, nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+accessToken)
	if err != nil {
		return spotiferr.BadRequestFormat
	}

	//HTTP Client makes the request to Spotify API Endpoint.
	response, _ := http.DefaultClient.Do(request)
	if response.StatusCode >= 400 {
		log.Println("Error: ", response)

		return spotiferr.InvalidToken
	}

	//Close http.Client Response
	defer response.Body.Close()

	//Read Response Body from HTTP Call
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return spotiferr.ResponseError
	}

	//Unmarshall JSON is use to MAP the values from body json to the Struct ArtistResponseJSON
	jsonResponse := ArtistResponseJSON{}
	if err := json.Unmarshal([]byte(string(body)), &jsonResponse); err != nil {
		return spotiferr.UnmarshallError

	}

	// Search through ITEM array (from Response body) and find the Band/artist Name.
	for _, itemArtist := range jsonResponse.Artists.Items {
		if strings.ToLower(itemArtist.Name) == strings.ToLower(artist) {
			targetArtist.ID = itemArtist.SpotifyID
			targetArtist.Name = itemArtist.Name
			targetArtist.SpotifyURL = itemArtist.ExternalUrls.Spotify

			//  Genres is a []string
			// []string... Let me pass multiple values to the func append.
			targetArtist.Genres = append(targetArtist.Genres, itemArtist.Genres...)
		}
	}

	if targetArtist.Name == "" {
		return spotiferr.ArtistNotFound
	}

	return nil

}

// RepreshResponseJSON is a struct that handle the "unmarshall" information from HTTP Response Body
// of RefreshToken.
type RefreshTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

// RefreshToken: Make a HTTP POST call to  Spotify API to get a new Token.
// this func will return an string (refresh token) and an error.
func RefreshToken() (string, error) {
	//Data for the HTTP POST CALL
	postData := url.Values{
		"grant_type":    {"refresh_token"},
		"refresh_token": {SpotifyRefreshToken},
	}

	//Create Request HTTP POST CALL structure with data and headers.
	request, err := http.NewRequest("POST", SpotifyRefreshEndpoint, strings.NewReader(postData.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Authorization", SpotifyAuhtorizationToken)
	if err != nil {
		return "", spotiferr.BadRequestFormat
	}

	//HTTP Client makes the request to the Spotify API Endpoint.
	response, _ := http.DefaultClient.Do(request)
	if response.StatusCode >= 400 {
		log.Println("Error: ", response)
		return "", spotiferr.InvalidToken

	}

	//Close http.Client
	defer response.Body.Close()

	//Read Response Body from HTTP Call Response
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", spotiferr.ResponseError

	}

	//Unmarshall JSON is use to MAP the values from body json to the Struct RefreshTokenResponse
	jsonResponse := RefreshTokenResponse{}
	if err := json.Unmarshal([]byte(string(body)), &jsonResponse); err != nil {
		return "", spotiferr.UnmarshallError

	}

	return jsonResponse.AccessToken, nil

}
