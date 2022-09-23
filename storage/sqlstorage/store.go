package sqlstorage

import (
	"apelsin/pkg/logger"
	"apelsin/storage"

	"github.com/jmoiron/sqlx"
)

type Store struct {
	db          *sqlx.DB
	log         logger.Logger
	region      *regionRepo
	district    *districtRepo
	nationality *nationalityRepo
	offer       *offerRepo
}

func (s *Store) Offer() storage.OfferRepository {
	if s.offer == nil {
		s.offer = &offerRepo{s}
	}

	return s.offer
}

func (s *Store) District() storage.DistrictRepository {
	if s.district == nil {
		s.district = &districtRepo{s}
	}

	return s.district
}

func (s *Store) Nationalities() storage.NationalityRepository {
	if s.nationality == nil {
		s.nationality = &nationalityRepo{s}
	}

	return s.nationality
}

func New(db *sqlx.DB, log logger.Logger) *Store {
	return &Store{
		db:  db,
		log: log,
	}
}

func (s *Store) Region() storage.RegionRepository {
	if s.region == nil {
		s.region = &regionRepo{s}
	}

	return s.region
}
