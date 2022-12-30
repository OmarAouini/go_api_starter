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
	AppName    string `env:"APP_NAME,required"`
	ApiSecret  string `env:"API_SECRET,required"`
	AppEnv     string `env:"APP_ENV,required"`
	DbName     string `env:"DB_NAME,required"`
	DbSchema   string `env:"DB_SCHEMA,required"`
	DbHost     string `env:"DB_HOST,required"`
	DbPort     string `env:"DB_PORT,required"`
	DbUsername string `env:"DB_USERNAME,required"`
	DbPassword string `env:"DB_PASSWORD,required"`
	DbSslMode  string `env:"DB_SSLMODE,required"`
	DbTimezone string `env:"DB_TIMEZONE,required"`
	DbMinConn  int    `env:"DB_MINCONN,required"`
	DbMaxConn  int    `env:"DB_MAXCONN,required"`

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
