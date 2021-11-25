package errors

import "errors"

//Errors for /service/csv.go
type csvError error

var FileError csvError = errors.New("Couldn't Open Csv File")
var ColumnParseError csvError = errors.New("First Column of CSV most have an Integer Value")
var CreateError csvError = errors.New("Couldn't Create the Csv File")
var BadWrite csvError = errors.New("Couldn't Write into the Csv File")

//Errors for /service/pokemon.go
type pokemonError error

var PokemonNotFound pokemonError = errors.New("Pokemon not Found.")

//Errors for /service/spotify.go
type spotifyError error

var ArtistNotFound spotifyError = errors.New("Artist not found.")
var InvalidToken spotifyError = errors.New("Invalid Token from Spotify.")
var TokenMissing spotifyError = errors.New("Couldn't retrieve token from Spotify.")
var BadRequestFormat spotifyError = errors.New("Bad Request HTTP Format.")
var ResponseError spotifyError = errors.New("Couldn't read Response body.")
var UnmarshallError spotifyError = errors.New("Couldn't map the API response Correctly.")
