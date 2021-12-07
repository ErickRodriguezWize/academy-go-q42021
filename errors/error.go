/* This file contains custom Errors use through the Proyect*/

package errors

import "errors"

type confError error

var (
	ErrUnmarshallYaml = errors.New("Error:Couldn't unmarshall taml file into a struct")
	ErrNotFoundYaml   = errors.New("Error: Couldn't find or read the Yaml file")
)

// Errors for /service/csv.go
type csvError error

var (
	ErrFileError        csvError = errors.New("Couldn't Open Csv File")
	ErrColumnParseError csvError = errors.New("First Column of CSV most have an Integer Value")
	ErrCreateError      csvError = errors.New("Couldn't Create the Csv File")
	ErrBadWrite         csvError = errors.New("Couldn't Write into the Csv File")
	ErrEndOfFile        csvError = errors.New("Couldn't Reach end of the file.")
)

// Errors for /service/pokemon.go
type pokemonError error

var ErrPokemonNotFound pokemonError = errors.New("Pokemon not Found")

// Errors for /service/spotify.go
type spotifyError error

var (
	ErrArtistNotFound   spotifyError = errors.New("Artist not found")
	ErrInvalidToken     spotifyError = errors.New("Invalid Token from Spotify")
	ErrTokenMissing     spotifyError = errors.New("Couldn't retrieve token from Spotify")
	ErrBadRequestFormat spotifyError = errors.New("Bad Request HTTP Format")
	ErrResponseError    spotifyError = errors.New("Couldn't read Response body")
	ErrUnmarshallError  spotifyError = errors.New("Couldn't map the API response Correctly")
	ErrHttpClient       spotifyError = errors.New("The Http Client couldn't make the request")
)