package entry

import (
	"context"
	"fmt"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type EnvConfig struct {
	DBHost     string `env:"DB_HOST"`
	DBName     string `env:"DB_NAME"`
	DBUser     string `env:"DB_USER"`
	DBPassword string `env:"DB_PASSWORD"`
	DBPort     string `env:"DB_PORT"`
	DBSSLMode  string `env:"DB_SSLMODE"`

	LOGLevel string `env:"LOG_LEVEL"`
	AppPort  string `env:"APP_PORT"`

	ApiGetAge     string `env:"GET_AGE"`
	ApiGetCountry string `env:"GET_COUNTRY"`
	ApiGetGender  string `env:"GET_GENDER"`
}

func ComposeServer() error {
	envConfig, err := InitializeEnv()
	if err != nil {
		return err
	}

	logger := InitializeLogger(envConfig)

	db, err := InitializeDB(context.Background(), envConfig, logger)
	if err != nil {
		return err
	}

	err = Migrate(db, logger)
	if err != nil {
		return err
	}

	return nil

}

func InitializeLogger(envConfig EnvConfig) *logrus.Logger {
	logger := logrus.New()

	if envConfig.LOGLevel == "debug" {
		logger.SetLevel(logrus.DebugLevel)
	} else {
		logger.SetLevel(logrus.InfoLevel)
	}

	return logger
}

func InitializeEnv() (envConfig EnvConfig, err error) {
	if err := godotenv.Load(); err != nil {
		return EnvConfig{}, fmt.Errorf("failed to load .env file: %w", err)
	}

	if err := env.Parse(&envConfig); err != nil {
		return EnvConfig{}, fmt.Errorf("failed to parse env from environment variables: %w", err)
	}

	return envConfig, nil
}
