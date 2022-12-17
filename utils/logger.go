package utils

import (
	log "github.com/sirupsen/logrus"
)

func Logger() *log.Logger {
	logger := log.New()
	logger.SetLevel(log.InfoLevel)
	logger.SetFormatter(&log.JSONFormatter{})

	return logger
}
