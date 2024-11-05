package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func InitLogger() {
	Log = logrus.New()
	Log.SetOutput(os.Stdout)
	Log.SetLevel(logrus.InfoLevel) // Set to DebugLevel for more detailed logs
	Log.SetFormatter(&logrus.JSONFormatter{})
}
