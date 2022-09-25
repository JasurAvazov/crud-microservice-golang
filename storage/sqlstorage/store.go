package sqlstorage

import (
	"apelsin/pkg/logger"
	"apelsin/storage"

	"github.com/jmoiron/sqlx"
)

type Store struct {
	db  *sqlx.DB
	log logger.Logger

	customer *customerRepo
	order    *orderRepo
}

func (s *Store) Customer() storage.CustomerRepository {
	if s.customer == nil {
		s.customer = &customerRepo{s}
	}

	return s.customer
}

func (s *Store) Order() storage.OrderRepository {
	if s.order == nil {
		s.order = &orderRepo{s}
	}

	return s.order
}

func New(db *sqlx.DB, log logger.Logger) *Store {
	return &Store{
		db:  db,
		log: log,
	}
}
