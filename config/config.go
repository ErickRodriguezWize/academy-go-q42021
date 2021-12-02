// This config file focus on reading the Environment Files of config.yaml using viper package.
package config

import (
	"github.com/ErickRodriguezWize/academy-go-q42021/domain/model"
	conferr "github.com/ErickRodriguezWize/academy-go-q42021/errors"

	"github.com/spf13/viper"
)

// LoadConfig: Make the init configurations and return model.Config with all the Env Values.
func LoadConfig() (*model.Config, error) {
	if err := InitConfig(); err != nil {

	}

	// Unmarshall yaml file(config.yaml) into model.Config struct.
	conf := &model.Config{}
	if err := viper.Unmarshal(conf); err != nil {
		return conf, conferr.ErrUnmarshallYaml
	}

	// Check for an empty field in config struct.
	if err := conf.ValidateFields(); err != nil {
		return conf, err
	}

	return conf, nil
}

// InitConfig: Make the configuration for the viper module (config file, define paths, read config.yaml file).
func InitConfig() error {
	// Set config filename
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	// Set config file path.
	viper.AddConfigPath(".")    //Normal Server Running.
	viper.AddConfigPath("./..") //In case of Testing Go files.

	// Read Yaml config file.
	err := viper.ReadInConfig()
	if err != nil {
		return conferr.ErrNotFoundYaml
	}

	return nil
}