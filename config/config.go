package config

import (
	"github.com/spf13/viper"
	"log"
)

type AppConfig struct {
	Server ServerConfig
	Redis  RedisConfig
}

type RedisConfig struct {
	Server string
	Port   int
}

type ServerConfig struct {
	Port int
}

var appConfig AppConfig

func Init(env string) {
	var err error

	viper.SetConfigType("yaml")
	viper.SetConfigName(env)
	viper.AddConfigPath("../config/")
	viper.AddConfigPath("config/")
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing configuration file")
	}

	marshalErr := viper.Unmarshal(&appConfig)

	if marshalErr != nil {
		log.Fatalf("Unable to decode into struct, %v \n", marshalErr)
		panic(marshalErr)
	}

	log.Printf("loaded configuration %+v", appConfig)

}

func GetConfig() AppConfig {
	return appConfig
}
