package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type Repository struct {
	db  *sqlx.DB
	log *logrus.Logger
}

func NewRepository(db *sqlx.DB, log *logrus.Logger) *Repository {
	return &Repository{
		db:  db,
		log: log,
	}
}
