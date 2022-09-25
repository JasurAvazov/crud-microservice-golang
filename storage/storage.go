package storage

import (
	"apelsin/contracts"
	"apelsin/models"
	"context"
)

// Storage ...
type Storage interface {
	contracts.ISessionProvider

	District() DistrictRepository
}

type RecordRetriever interface {
	GetRecords(ctx context.Context) ([]models.Record, error)
}

// DistrictRepository ...
type DistrictRepository interface {
	RecordRetriever
	Create(context.Context, models.District) (models.District, error)
	Read(context.Context, string) (models.District, error)
	ReadAll(context.Context) ([]models.District, error)
	Update(context.Context, models.District) error
	Delete(context.Context, string) error
	ClearAll(ctx context.Context) error
}
