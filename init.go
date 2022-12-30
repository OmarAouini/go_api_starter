package main

import (
	"fmt"

	"github.com/OmarAouini/go_tdd/config"
	"github.com/OmarAouini/go_tdd/logging"
)

// misc configuration initialization such as logger, environment, etc...
func Init() {
	logging.ConfigureLogger()
	config.LoadEnv()
}

func PrintAppInfo() {
	fmt.Printf("-APP NAME: %s\n-APP ENV: %s\n-DB NAME: %s\n-DB_HOST: %s\n-DB_SCHEMA: %s\n-DB_TIMEZONE: %s", config.AppConfiguration.AppName, config.AppConfiguration.AppEnv, config.AppConfiguration.DbName, config.AppConfiguration.DbHost, config.AppConfiguration.DbSchema, config.AppConfiguration.DbTimezone)
}
