/* This service focus on reading the Environment Files of config.yaml using viper package.  */

package service

import (
	"log"

	"github.com/spf13/viper"
)

// SetupViper: Make the configuration for the viper module (config file, define paths, read config.yaml file).
func SetupViper() {
	// Set config filename
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	// Set config file path.
	viper.AddConfigPath(".")    //Normal Server Running.
	viper.AddConfigPath("./..") //In case of Testing Go files.

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Error: Couldn't Read or find the Environment File.")
	}
}

// GetEnvVariable: Find a String value inside config.yaml(Environment variable file).
func GetEnvVariable(key string) string {
	SetupViper()

	//This method returns value (string) and a boolean
	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Error: Couldn't find the string value(%v) on the Environment Variable file.", key)
	}

	return value

}

// GetInvVariable: Find an Integer value on config.yaml(Environment variable file).
func GetIntEnvVariable(key string) int {
	SetupViper()

	//This methods return value (int).
	value := viper.GetInt(key)
	if value == 0 {
		log.Fatalf("Error: Couldn't find the Integer Value(%v) on the Environment Variable File.", key)
	}

	return value
}