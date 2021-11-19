package model

type Artist struct{
	ID string `json:"spotify_id"`
	Name string `json:"artist_name"`
	SpotifyURL string `json:"spotify_url"`
	Genres []string `json:"genres"`
}