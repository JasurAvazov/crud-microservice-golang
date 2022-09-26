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
	category *categoryRepo
	product  *productRepo
	detail   *detailRepo
	invoice  *invoiceRepo
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

func (s *Store) Category() storage.CategoryRepository {
	if s.category == nil {
		s.category = &categoryRepo{s}
	}

	return s.category
}

func (s *Store) Product() storage.ProductRepository {
	if s.product == nil {
		s.product = &productRepo{s}
	}

	return s.product
}

func (s *Store) Detail() storage.DetailRepository {
	if s.detail == nil {
		s.detail = &detailRepo{s}
	}

	return s.detail
}

func (s *Store) Invoice() storage.InvoiceRepository {
	if s.invoice == nil {
		s.invoice = &invoiceRepo{s}
	}

	return s.invoice
}

func New(db *sqlx.DB, log logger.Logger) *Store {
	return &Store{
		db:  db,
		log: log,
	}
}
