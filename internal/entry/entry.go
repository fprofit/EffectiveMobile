package entry

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func ComposeServer() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("Failed to load .env file: %w", err)
	}

	logger := InitializeLogger()

	logger.Debug("debug")
	logger.Info("info")
	return nil

}

func InitializeLogger() (logger *logrus.Logger) {
	logger = logrus.New()
	if os.Getenv("LOG_LEVEL") == "debug" {
		logger.SetLevel(logrus.DebugLevel)
	} else {
		logger.SetLevel(logrus.InfoLevel)
	}
	return logger
}
