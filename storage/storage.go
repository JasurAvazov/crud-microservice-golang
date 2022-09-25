package storage

import (
	"apelsin/contracts"
	"apelsin/models"
	"context"
)

// Storage ...
type Storage interface {
	contracts.ISessionProvider

	Customer() CustomerRepository
	Order() OrderRepository
}

type RecordRetriever interface {
	GetRecords(ctx context.Context) ([]models.Record, error)
}

// CustomerRepository ...
type CustomerRepository interface {
	RecordRetriever
	Create(context.Context, models.Customer) (models.Customer, error)
	Read(context.Context, string) (models.Customer, error)
	ReadAll(context.Context) ([]models.Customer, error)
	Update(context.Context, models.Customer) error
	Delete(context.Context, string) error
	ClearAll(ctx context.Context) error
}

// OrderRepository ...
type OrderRepository interface {
	RecordRetriever
	Create(context.Context, models.Order) (models.Order, error)
	Read(context.Context, string) (models.Order, error)
	ReadAll(context.Context) ([]models.Order, error)
	Update(context.Context, models.Order) error
	Delete(context.Context, string) error
	ClearAll(ctx context.Context) error
}
