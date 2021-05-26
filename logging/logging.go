package logging

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Log is the logrus Logger
var Log *logrus.Logger

// Init allow to initialize the logger
func Init() {
	Log = logrus.New()

	switch viper.GetString("loglevel") {
	case "Info":
		Log.SetLevel(logrus.InfoLevel)
	case "Warning":
		Log.SetLevel(logrus.WarnLevel)
	case "Debug":
		Log.SetLevel(logrus.DebugLevel)
	default:
		Log.SetLevel(logrus.InfoLevel)
	}

	Log.SetOutput(os.Stdout)
}
