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

func initializeDB(ctx context.Context, config EnvConfig, log *logrus.Logger) (*sqlx.DB, error) {
	log.Debug("Initializing DB connection")

	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName, config.DBSSLMode,
	)

	db, err := sqlx.Open("pgx", connectionString)
	if err != nil {
		log.WithError(err).Error("Failed to open DB connection")
		return nil, fmt.Errorf("InitializeDB sqlx open: %w", err)
	}

	log.Debug("Pinging DB")
	if err := db.PingContext(ctx); err != nil {
		log.WithError(err).Error("Failed to ping DB")
		db.Close()
		return nil, fmt.Errorf("InitializeDB sqlx ping: %w", err)
	}

	log.Info("DB connection successfully established")
	return db, nil
}

func migrateDB(db *sqlx.DB, log *logrus.Logger) error {
	log.Debug("Running database migrations")

	driver, err := pgx.WithInstance(db.DB, &pgx.Config{})
	if err != nil {
		log.WithError(err).Error("Failed to create migrate instance")
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"pgx", driver,
	)
	if err != nil {
		log.WithError(err).Error("Could not create migrate instance")
		return fmt.Errorf("Could not create migrate instance: %w", err)
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.WithError(err).Error("Failed to apply migrations")
		return err
	}

	log.Info("Database migrations applied successfully")
	return nil
}

func shutdownDBConnections(db *sqlx.DB, log *logrus.Logger) {
	log.Debug("Shutting down database connection...")

	if err := db.Close(); err != nil {
		log.Error("Error closing database connection:", err)
		return
	}

	log.Info("Database connection closed")
	return
}
