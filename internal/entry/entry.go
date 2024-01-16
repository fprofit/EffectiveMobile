package entry

import (
	"context"
	"os"
	"os/signal"

	"github.com/fprofit/EffectiveMobile/internal/handler"
	"github.com/fprofit/EffectiveMobile/internal/repository"
	"github.com/fprofit/EffectiveMobile/internal/service"

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
	envConfig, err := initializeEnv()
	if err != nil {
		return err
	}

	log := initializelog(envConfig)

	db, err := initializeDB(context.Background(), envConfig, log)
	if err != nil {
		return err
	}

	err = migrateDB(db, log)
	if err != nil {
		return err
	}

	repo := repository.NewRepository(db, log)

	serviceApiUrl := service.ApiUrl{
		ApiGetAge:     envConfig.ApiGetAge,
		ApiGetCountry: envConfig.ApiGetCountry,
		ApiGetGender:  envConfig.ApiGetGender}

	serv := service.NewService(repo, serviceApiUrl, log)

	handler := handler.NewHandler(serv, log)
	s := runServer(handler, envConfig, log)

	waitForShutdown()

	shutdownServer(s, log)
	shutdownDBConnections(db, log)

	return nil
}

func initializelog(envConfig EnvConfig) *logrus.Logger {
	log := logrus.New()

	if envConfig.LOGLevel == "debug" {
		log.SetLevel(logrus.DebugLevel)
	} else {
		log.SetLevel(logrus.InfoLevel)
	}

	return log
}

func waitForShutdown() {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
