package controller

import (
	"apelsin/models"
	"context"
	"fmt"
	"io"
	"time"
)

func (ctrl *Controller) ImportDistrictsToDB(ctx context.Context, file io.Reader) ([]models.District, error) {
	districts, err := ctrl.ImportDistricts(file)
	if err != nil {
		return nil, err
	}
	session, ctx, err := ctrl.store.StartSession(ctx)
	if err != nil {
		return nil, err
	}
	var sessionErr error
	defer func() { session.Close(sessionErr) }()
	repo := ctrl.store.District()
	if sessionErr = repo.ClearAll(ctx); sessionErr != nil {
		return nil, sessionErr
	}
	for i := range districts {
		if _, err := repo.Create(ctx, districts[i]); err != nil {
			sessionErr = err
			return nil, err
		}
	}

	return districts, nil
}

func (ctrl *Controller) ImportDistricts(file io.Reader) ([]models.District, error) {
	now := time.Now()
	rows, err := ctrl._parseFile(file, districtValidator)
	if err != nil {
		return nil, err
	}
	districts := make([]models.District, 0)
	for i, cols := range rows {
		if i < 1 {
			ctrl.Debug(fmt.Sprintf("Row: %d;", i))
			i++
			continue
		}
		if len(cols) < 7 {
			ctrl.Error(fmt.Sprintf("Row: %d; Columns count: %d", i, len(cols)))
			break
		}
		district := models.District{}
		district.CodeSOATO = cols[0]
		district.Code = cols[1]
		district.CodeGNI = cols[2]
		district.CodeProvince = cols[3]
		district.Title.Ru = cols[4]
		district.Title.En = cols[5]
		district.Title.Uz = cols[6]
		district.ActivatedAt = &now
		districts = append(districts, district)
		i++
	}
	return districts, nil
}

func (ctrl *Controller) DeleteDistrict(ctx context.Context, code string) error {
	session, ctx, err := ctrl.store.StartSession(ctx)
	if err != nil {
		return err
	}
	var sessionErr error
	defer func() { session.Close(sessionErr) }()
	repo := ctrl.store.District()
	sessionErr = repo.Delete(ctx, code)
	if sessionErr != nil {
		return sessionErr
	}
	return nil
}

func (ctrl *Controller) GetDistrict(ctx context.Context, code string) (models.District, error) {
	repo := ctrl.store.District()
	districts, err := repo.Read(ctx, code)
	if err != nil {
		return models.District{}, err
	}
	return districts, nil
}

func (ctrl *Controller) CreateDistrict(ctx context.Context, model models.District) error {
	var sessionErr error
	session, ctx, err := ctrl.store.StartSession(ctx)
	if err != nil {
		return err
	}
	defer func() { session.Close(sessionErr) }()
	repo := ctrl.store.District()

	if _, sessionErr = repo.Create(ctx, model); sessionErr != nil {
		return sessionErr
	}
	return nil
}

func (ctrl *Controller) UpdateDistrict(ctx context.Context, model models.District) error {
	var sessionErr error
	session, ctx, err := ctrl.store.StartSession(ctx)
	if err != nil {
		return err
	}
	defer func() { session.Close(sessionErr) }()
	repo := ctrl.store.District()

	if sessionErr = repo.Update(ctx, model); sessionErr != nil {
		return sessionErr
	}
	return nil
}

func (ctrl *Controller) ListDistrict(ctx context.Context) ([]models.District, error) {
	repo := ctrl.store.District()
	districts, err := repo.ReadAll(ctx)
	if err != nil {
		return nil, err
	}
	return districts, nil
}
