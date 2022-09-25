package sqlstorage

import (
	"apelsin/pkg/logger"
	"apelsin/storage"

	"github.com/jmoiron/sqlx"
)

type Store struct {
	db       *sqlx.DB
	log      logger.Logger
	customer *customerRepo
}

func (s *Store) Customer() storage.CustomerRepository {
	if s.customer == nil {
		s.customer = &customerRepo{s}
	}

	return s.customer
}

func New(db *sqlx.DB, log logger.Logger) *Store {
	return &Store{
		db:  db,
		log: log,
	}
}
