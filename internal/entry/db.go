package entry

import (
	"context"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/pgx"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func InitializeDB(ctx context.Context, envConfig EnvConfig, logger *logrus.Logger) (*sqlx.DB, error) {
	logger.Debug("Initializing DB connection")

	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		envConfig.DBHost, envConfig.DBPort, envConfig.DBUser, envConfig.DBPassword, envConfig.DBName, envConfig.DBSSLMode,
	)

	db, err := sqlx.Open("pgx", connectionString)
	if err != nil {
		logger.WithError(err).Error("Failed to open DB connection")
		return nil, fmt.Errorf("InitializeDB sqlx open: %w", err)
	}

	logger.Debug("Pinging DB")
	if err := db.PingContext(ctx); err != nil {
		logger.WithError(err).Error("Failed to ping DB")
		db.Close()
		return nil, fmt.Errorf("InitializeDB sqlx ping: %w", err)
	}

	logger.Info("DB connection successfully established")
	return db, nil
}

func Migrate(db *sqlx.DB, logger *logrus.Logger) error {
	logger.Debug("Running database migrations")

	driver, err := pgx.WithInstance(db.DB, &pgx.Config{})
	if err != nil {
		logger.WithError(err).Error("Failed to create migrate instance")
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"pgx", driver,
	)
	if err != nil {
		logger.WithError(err).Error("Could not create migrate instance")
		return fmt.Errorf("Could not create migrate instance: %w", err)
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		logger.WithError(err).Error("Failed to apply migrations")
		return err
	}

	logger.Info("Database migrations applied successfully")
	return nil
}
