package model

// Pokemon struct will handle the information of a pokemon.
type Pokemon struct {
	ID   int    `json:"pokedex_number"`
	Name string `json:"pokemon_name"`
}