package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	Log *logrus.Logger
}

func NewLogger() *Logger {
	log := logrus.New()
	file, err := os.OpenFile("logs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr")
	}
	return &Logger{
		Log: log,
	}
}
