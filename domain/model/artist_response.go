package model

// ArtistResponse is a struct that handle the "unmarshall" information from HTTP Response Body.
type ArtistResponse struct {
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