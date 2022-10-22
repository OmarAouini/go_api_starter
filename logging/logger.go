package logging

import "github.com/sirupsen/logrus"

// Logrus logger custom configuration
func ConfigureLogger() {
	// TODO customize logging format here, refer to https://github.com/sirupsen/logrus
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true, ForceColors: true})
}
