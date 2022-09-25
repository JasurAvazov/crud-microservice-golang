package sqlstorage

import (
	"apelsin/pkg/logger"
	"apelsin/storage"

	"github.com/jmoiron/sqlx"
)

type Store struct {
	db       *sqlx.DB
	log      logger.Logger
	district *districtRepo
}

func (s *Store) District() storage.DistrictRepository {
	if s.district == nil {
		s.district = &districtRepo{s}
	}

	return s.district
}

func New(db *sqlx.DB, log logger.Logger) *Store {
	return &Store{
		db:  db,
		log: log,
	}
}
