package entry

import (
	"fmt"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

func initializeEnv() (config EnvConfig, err error) {
	if err := godotenv.Load(); err != nil {
		return EnvConfig{}, fmt.Errorf("Failed to load .env file: %w", err)
	}

	if err := env.Parse(&config); err != nil {
		return EnvConfig{}, fmt.Errorf("Failed to parse env from environment variables: %w", err)
	}
	if err := validateEnvConfig(config); err != nil {
		return EnvConfig{}, err
	}

	return config, nil
}

func validateEnvConfig(config EnvConfig) error {
	switch {
	case config.DBHost == "":
		return fmt.Errorf("DBHost is a required environment variable")
	case config.DBPort == "":
		return fmt.Errorf("DBPort is a required environment variable")
	case config.DBUser == "":
		return fmt.Errorf("DBUser is a required environment variable")
	case config.DBPassword == "":
		return fmt.Errorf("DBPassword is a required environment variable")
	case config.DBName == "":
		return fmt.Errorf("DBName is a required environment variable")
	case config.DBSSLMode == "":
		return fmt.Errorf("DBSSLMode is a required environment variable")
	case config.AppPort == "":
		return fmt.Errorf("AppPort is a required environment variable")
	case config.ApiGetAge == "":
		return fmt.Errorf("ApiGetAge is a required environment variable")
	case config.ApiGetCountry == "":
		return fmt.Errorf("ApiGetCountry is a required environment variable")
	case config.ApiGetGender == "":
		return fmt.Errorf("ApiGetGender is a required environment variable")
	default:
		return nil
	}
}
