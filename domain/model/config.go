package model 

type Config struct {
	Ip string `mapstructure:"IP_ADDRESS"`
	Port string `mapstructure:"PORT"`
	ReadTimeout int `mapstructure:"READ_TIMEOUT"`
	WriteTimeout int `mapstructure:"WRITE_TIMEOUT"`
	SpotifyEndpoint string `mapstructure:"SPOTIFY_API_ENDPOINT"`
	RefreshEndpoint string `mapstructure:"SPOTIFY_REFRESH_ENDPOINT"`
	RefreshToken string `mapstructure:"SPOTIFY_REFRESH_TOKEN"`
	AuthorizationToken string `mapstructure:"SPOTIFY_AUTHORIZATION_TOKEN"`
	LimitArtist string `mapstructure:"LIMIT_ARTIST"`
	PokemonCsvPath string `mapstructure:"CSV_PATH"`
	ArtistCsvPath string `mapstructure:"CSV_ARTIST_PATH"`
}
