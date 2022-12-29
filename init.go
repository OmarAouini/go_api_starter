package main

import (
	"github.com/OmarAouini/go_tdd/config"
	"github.com/OmarAouini/go_tdd/logging"
)

// misc configuration initialization such as logger, environment, etc...
func Init() {
	logging.ConfigureLogger()
	config.LoadEnv()
}
