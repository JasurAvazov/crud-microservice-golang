package storage

import (
	"apelsin/contracts"
	"apelsin/models"
	"context"
)

// Storage ...
type Storage interface {
	contracts.ISessionProvider

	Region() RegionRepository
	District() DistrictRepository
	Offer() OfferRepository
	Nationalities() NationalityRepository
}

type RecordRetriever interface {
	GetRecords(ctx context.Context) ([]models.Record, error)
}

// OfferRepository ...
type OfferRepository interface {
	RecordRetriever
	Create(context.Context, models.Offer) (models.Offer, error)
	Read(context.Context, string) (models.Offer, error)
	ReadAll(context.Context) ([]models.Offer, error)
	Update(context.Context, models.Offer) error
	Delete(context.Context, string) error
	GetFullOffer(ctx context.Context, code, locale string) (models.FullOffer, error)
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

// NationalityRepository ...
type NationalityRepository interface {
	RecordRetriever
	Create(context.Context, models.Nationality) (models.Nationality, error)
	Read(context.Context, string) (models.Nationality, error)
	ReadAll(context.Context) ([]models.Nationality, error)
	Update(context.Context, models.Nationality) error
	Delete(context.Context, string) error
	ClearAll(ctx context.Context) error
}

// RegionRepository ...
type RegionRepository interface {
	RecordRetriever
	Create(context.Context, models.Region) (models.Region, error)
	Read(context.Context, string) (models.Region, error)
	ReadAll(context.Context) ([]models.Region, error)
	Update(context.Context, models.Region) error
	Delete(context.Context, string) error
	ClearAll(ctx context.Context) error
}
