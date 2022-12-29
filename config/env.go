package config

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

// AppConfiguration var to be accessed globally
var AppConfiguration *Configuration

// Main configuration app wrapper, contains all the env variables loaded during app startup
type Configuration struct {
	AppName   string `env:"APP_NAME,required"`
	ApiSecret string `env:"API_SECRET,required"`
	// TODO add new env vars here, refer to https://github.com/caarlos0/env
}

// Load all the environment variables and load it into the global config.AppConfiguration struct
func LoadEnv() {
	//try .env file
	loadDotEnvFile()

	//fill global configuration var
	cfg := Configuration{}
	if err := env.Parse(&cfg); err != nil {
		logrus.Fatalf("%+v\n", err)
	}

	AppConfiguration = &cfg

}

// Load the environment variables from a .env file in the root project folder, if present
func loadDotEnvFile() {
	err := godotenv.Load()
	if err != nil {
		logrus.Info(".env file not found, going production mode...")
	}
}
