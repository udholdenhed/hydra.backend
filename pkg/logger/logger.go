package logger

import (
	"os"
	"sync"

	log "github.com/sirupsen/logrus"
)

// FileName name of the log file in which the logs will be written.
const FileName = "hydra-logs.json"

var (
	once   sync.Once
	logger *log.Logger
)

// Log returns a logger instance.
func Log() *log.Logger {
	once.Do(func() {
		file, err := os.OpenFile(FileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
		if err != nil {
			panic(err)
		}

		logger = log.New()
		logger.SetOutput(file)
		logger.SetFormatter(&log.JSONFormatter{})
	})
	return logger
}
