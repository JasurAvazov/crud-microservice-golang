package controller

import (
	"apelsin/models"
	"context"
	"fmt"
	"io"
	"time"
)

func (ctrl *Controller) ImportRegionsToDB(ctx context.Context, file io.Reader) ([]models.Region, error) {
	regions, err := ctrl.ImportRegions(file)
	if err != nil {
		return nil, err
	}
	var sessionErr error
	session, ctx, err := ctrl.store.StartSession(ctx)
	if err != nil {
		return nil, err
	}
	defer func() { session.Close(sessionErr) }()
	repo := ctrl.store.Region()
	if sessionErr = repo.ClearAll(ctx); sessionErr != nil {
		return nil, sessionErr
	}
	for i := range regions {
		if _, err := repo.Create(ctx, regions[i]); err != nil {
			sessionErr = err
			return nil, err
		}
	}
	return regions, nil
}

func (ctrl *Controller) ImportRegions(file io.Reader) ([]models.Region, error) {
	now := time.Now()
	rows, err := ctrl._parseFile(file, regionValidator)
	if err != nil {
		return nil, err
	}
	regions := make([]models.Region, 0)
	for i, cols := range rows {
		if i < 1 {
			ctrl.Debug(fmt.Sprintf("Row: %d;", i))
			i++
			continue
		}
		if len(cols) < 5 {
			ctrl.Error(fmt.Sprintf("Row: %d; Columns count: %d", i, len(cols)))
			break
		}
		region := models.Region{}
		region.Code = cols[0]
		region.Title.Ru = cols[1]
		region.Title.En = cols[2]
		region.Title.Uz = cols[3]
		region.Rank = cols[4]
		region.ActivatedAt = &now
		regions = append(regions, region)
		i++
	}
	return regions, nil
}

func (ctrl *Controller) CreateRegion(ctx context.Context, model models.Region) error {
	var sessionErr error
	session, ctx, err := ctrl.store.StartSession(ctx)
	if err != nil {
		return err
	}
	defer func() { session.Close(sessionErr) }()
	repo := ctrl.store.Region()

	if _, sessionErr = repo.Create(ctx, model); sessionErr != nil {
		return sessionErr
	}
	return nil
}

func (ctrl *Controller) UpdateRegion(ctx context.Context, model models.Region) error {
	var sessionErr error
	session, ctx, err := ctrl.store.StartSession(ctx)
	if err != nil {
		return err
	}
	defer func() { session.Close(sessionErr) }()
	repo := ctrl.store.Region()

	if sessionErr = repo.Update(ctx, model); sessionErr != nil {
		return sessionErr
	}
	return nil
}

func (ctrl *Controller) ListRegions(ctx context.Context) ([]models.Region, error) {
	repo := ctrl.store.Region()
	regions, err := repo.ReadAll(ctx)
	if err != nil {
		return nil, err
	}
	return regions, nil
}

func (ctrl *Controller) DeleteRegion(ctx context.Context, code string) error {
	session, ctx, err := ctrl.store.StartSession(ctx)
	if err != nil {
		return err
	}
	var sessionErr error
	defer func() { session.Close(sessionErr) }()
	repo := ctrl.store.Region()
	sessionErr = repo.Delete(ctx, code)
	if sessionErr != nil {
		return sessionErr
	}
	return nil
}

func (ctrl *Controller) GetRegion(ctx context.Context, code string) (models.Region, error) {
	repo := ctrl.store.Region()
	regions, err := repo.Read(ctx, code)
	if err != nil {
		return models.Region{}, err
	}
	return regions, nil
}
