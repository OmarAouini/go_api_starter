package main

import (
	"github.com/OmarAouini/go_tdd/config"
	"github.com/OmarAouini/go_tdd/logging"
	"github.com/sirupsen/logrus"
)

func main() {
	logging.ConfigureLogger()
	config.LoadEnv()
	logrus.Info("hello from main")
}
