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
	Category() CategoryRepository
	Product() ProductRepository
	Detail() DetailRepository
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

// CategoryRepository ...
type CategoryRepository interface {
	RecordRetriever
	Create(context.Context, models.Category) (models.Category, error)
	Read(context.Context, string) (models.Category, error)
	ReadAll(context.Context) ([]models.Category, error)
	Update(context.Context, models.Category) error
	Delete(context.Context, string) error
	ClearAll(ctx context.Context) error
}

// ProductRepository ...
type ProductRepository interface {
	RecordRetriever
	Create(context.Context, models.Product) (models.Product, error)
	Read(context.Context, string) (models.Product, error)
	ReadAll(context.Context) ([]models.Product, error)
	Update(context.Context, models.Product) error
	Delete(context.Context, string) error
	ClearAll(ctx context.Context) error
}

// DetailRepository ...
type DetailRepository interface {
	RecordRetriever
	Create(context.Context, models.Detail) (models.Detail, error)
	Read(context.Context, string) (models.Detail, error)
	ReadAll(context.Context) ([]models.Detail, error)
	Update(context.Context, models.Detail) error
	Delete(context.Context, string) error
	ClearAll(ctx context.Context) error
}
