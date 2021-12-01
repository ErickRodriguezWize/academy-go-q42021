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

// SearcArtist: Makes a HTTP GET call to Spotify API to search for an Artist Information.
// This func will return an error (in case that one is triggered).
func SearchArtist(artist string, targetArtist *model.Artist, config model.Config) error {
	// Get token from Spotify API (Token experies every 30 minutes.)
	accessToken, err := RefreshToken(config.RefreshEndpoint, config.RefreshToken, config.AuthorizationToken)
	if err != nil {
		return spotiferr.ErrTokenMissing
	}

	spotifyEndpoint := config.SpotifyEndpoint
	limitArtist := config.LimitArtist
	artist = strings.ReplaceAll(artist, "+", " ")

	// Encoded spaces with "%20". Spotify ask for this escaping for correctly HTTP Calls.
	encodedArtist := url.QueryEscape(artist)
	endpoint := spotifyEndpoint + "search?q=artist:" + encodedArtist + "&type=artist&limit=" + limitArtist
	log.Println("*** HTTP GET to", endpoint)

	// Create request with URL endpoint, Method and Headers.
	request, err := http.NewRequest("GET", endpoint, nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+accessToken)
	if err != nil {
		return spotiferr.ErrBadRequestFormat
	}

	// HTTP Client makes the request to Spotify API Endpoint.
	response, _ := http.DefaultClient.Do(request)
	if response.StatusCode >= 400 {
		log.Println("Error: ", response)

		return spotiferr.ErrInvalidToken
	}

	// Close http.Client Response
	defer response.Body.Close()

	// Read Response Body from HTTP Call
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return spotiferr.ErrResponseError
	}

	// Unmarshall JSON is use to MAP the values from body json to the Struct ArtistResponseJSON
	jsonResponse := model.ArtistResponse{}
	if err := json.Unmarshal([]byte(string(body)), &jsonResponse); err != nil {
		return spotiferr.ErrUnmarshallError

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

	// Check if the Artist was found or not. 
	if targetArtist.Name == "" {
		return spotiferr.ErrArtistNotFound
	}

	return nil

}

// RefreshToken: Make a HTTP POST call to  Spotify API to get a new Token.
// this func will return an string (refresh token) and an error.
func RefreshToken(endpoint string, refreshToken string, authToken string) (string, error) {
	// Data for the HTTP POST CALL
	postData := url.Values{
		"grant_type":    {"refresh_token"},
		"refresh_token": {refreshToken},
	}

	// Create Request HTTP POST CALL structure with data and headers.
	request, err := http.NewRequest("POST", endpoint, strings.NewReader(postData.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Authorization", authToken)
	if err != nil {
		return "", spotiferr.ErrBadRequestFormat
	}

	//HTTP Client makes the request to the Spotify API Endpoint.
	response, _ := http.DefaultClient.Do(request)
	if response.StatusCode >= 400 {
		log.Println("Error: ", response)
		return "", spotiferr.ErrInvalidToken

	}

	// Close http.Client.
	defer response.Body.Close()

	// Read Response Body from HTTP Call Response.
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", spotiferr.ErrResponseError

	}

	// Unmarshall JSON is use to MAP the values from body json to the Struct RefreshTokenResponse.
	jsonResponse := model.TokenResponse{}
	if err := json.Unmarshal([]byte(string(body)), &jsonResponse); err != nil {
		return "", spotiferr.ErrUnmarshallError

	}

	return jsonResponse.AccessToken, nil

}