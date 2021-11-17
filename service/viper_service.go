/* This service focus on reading the Environment Files of config.yaml using viper package.  */

package service

import (
	"log"

	"github.com/spf13/viper"
)

func GetEnvVariable(key string) (string) {
	//Set config filename
	viper.SetConfigName("config")
	//Set config file path. 
	viper.AddConfigPath(".")

	//Find and read config file
	err:= viper.ReadInConfig()

	if err != nil{
		log.Fatal("Error: Couldn't Read or find the Environment File.")
	}	

	//This method returns value and a boolean
	value, ok := viper.Get(key).(string)
	if !ok{
		log.Fatalf("Error: Couldn't find the key(%v) on the Environment Variable file.", key)
	}

	return value

}