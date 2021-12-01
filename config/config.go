// This config file focus on reading the Environment Files of config.yaml using viper package.
package config

import (
	"log"

	"github.com/ErickRodriguezWize/academy-go-q42021/domain/model"

	"github.com/spf13/viper"
)

// LoadConfig: Make the init configurations and return model.Config with all the Env Values.
func LoadConfig() *model.Config {
	InitConfig()

	// Unmarshall yaml file(config.yaml) into model.Config struct.
	conf := &model.Config{}
	if err := viper.Unmarshal(conf); err != nil {
		log.Fatal("Error: Couldn't unmarshall yaml file into a struct.")
	}

	return conf
}

// InitConfig: Make the configuration for the viper module (config file, define paths, read config.yaml file).
func InitConfig() {
	// Set config filename
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	// Set config file path.
	viper.AddConfigPath(".")    //Normal Server Running.
	viper.AddConfigPath("./..") //In case of Testing Go files.

	// Read Yaml config file.
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Error: Couldn't Read or find the Environment File.")
	}
}