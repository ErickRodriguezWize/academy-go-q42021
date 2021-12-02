package model

import (
	"errors"
	"reflect"
)

// Config Struct will be used to "Unmarshall" all the Environment variables from config.yaml using viper.
type Config struct {
	Ip                 string `mapstructure:"IP_ADDRESS"`
	Port               string `mapstructure:"PORT"`
	ReadTimeout        int    `mapstructure:"READ_TIMEOUT"`
	WriteTimeout       int    `mapstructure:"WRITE_TIMEOUT"`
	SpotifyEndpoint    string `mapstructure:"SPOTIFY_API_ENDPOINT"`
	RefreshEndpoint    string `mapstructure:"SPOTIFY_REFRESH_ENDPOINT"`
	RefreshToken       string `mapstructure:"SPOTIFY_REFRESH_TOKEN"`
	AuthorizationToken string `mapstructure:"SPOTIFY_AUTHORIZATION_TOKEN"`
	LimitArtist        string `mapstructure:"LIMIT_ARTIST"`
	PokemonCsvPath     string `mapstructure:"CSV_PATH"`
	ArtistCsvPath      string `mapstructure:"CSV_ARTIST_PATH"`
}

// ValidateFields: Func that iterate through struct Config fields and validate if the value is empty or not. ("" or 0).
func (conf Config) ValidateFields() error {
	v := reflect.ValueOf(conf)

	for i := 0; i < v.NumField(); i++ {
		fieldName := v.Type().Field(i).Name
		fieldValue := v.Field(i)
		// Check if the field from struct has the zero value(string is "" and int is 0).
		if fieldValue.IsZero() {
			return errors.New("The Env variable " + fieldName + " is empty.")
		}

	}
	return nil
}