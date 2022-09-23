package controller

import (
	"apelsin/models"
	"context"
	"fmt"
	"io"
	"time"
)

func (ctrl *Controller) CreateNationality(ctx context.Context, model models.Nationality) error {
	var sessionErr error
	session, ctx, err := ctrl.store.StartSession(ctx)
	if err != nil {
		return err
	}
	defer func() { session.Close(sessionErr) }()
	repo := ctrl.store.Nationalities()

	if _, sessionErr = repo.Create(ctx, model); sessionErr != nil {
		return sessionErr
	}
	return nil
}

func (ctrl *Controller) GetNationality(ctx context.Context, code string) (models.Nationality, error) {
	repo := ctrl.store.Nationalities()
	nationalities, err := repo.Read(ctx, code)
	if err != nil {
		return models.Nationality{}, err
	}
	return nationalities, nil
}

func (ctrl *Controller) ImportNationalities(file io.Reader) ([]models.Nationality, error) {
	now := time.Now()
	rows, err := ctrl._parseFile(file, nationalityValidator)
	if err != nil {
		return nil, err
	}
	nationalities := make([]models.Nationality, 0)
	for i, cols := range rows {
		if i < 1 {
			ctrl.Debug(fmt.Sprintf("Row: %d;", i))
			i++
			continue
		}
		if len(cols) < 4 {
			ctrl.Error(fmt.Sprintf("Row: %d; Columns count: %d", i, len(cols)))
			break
		}
		nationality := models.Nationality{}
		nationality.Code = cols[0]
		nationality.Title.Ru = cols[1]
		nationality.Title.En = cols[2]
		nationality.Title.Uz = cols[3]
		nationality.ActivatedAt = &now
		nationality.DeactivatedAt = &now
		nationalities = append(nationalities, nationality)
		i++
	}
	return nationalities, nil
}

func (ctrl *Controller) ImportNationalitiesToDB(ctx context.Context, file io.Reader) ([]models.Nationality, error) {
	nationalities, err := ctrl.ImportNationalities(file)
	if err != nil {
		return nil, err
	}
	session, ctx, err := ctrl.store.StartSession(ctx)
	if err != nil {
		return nil, err
	}
	var sessionErr error
	defer func() { session.Close(sessionErr) }()
	repo := ctrl.store.Nationalities()
	if sessionErr = repo.ClearAll(ctx); sessionErr != nil {
		return nil, sessionErr
	}
	for i := range nationalities {
		if _, err := repo.Create(ctx, nationalities[i]); err != nil {
			sessionErr = err
			return nil, err
		}
	}
	return nationalities, nil
}

func (ctrl *Controller) ListNationalities(ctx context.Context) ([]models.Nationality, error) {
	repo := ctrl.store.Nationalities()
	nationalities, err := repo.ReadAll(ctx)
	if err != nil {
		return nil, err
	}
	return nationalities, nil
}

func (ctrl *Controller) UpdateNationality(ctx context.Context, model models.Nationality) error {
	var sessionErr error
	session, ctx, err := ctrl.store.StartSession(ctx)
	if err != nil {
		return err
	}
	defer func() { session.Close(sessionErr) }()
	repo := ctrl.store.Nationalities()

	if sessionErr = repo.Update(ctx, model); sessionErr != nil {
		return sessionErr
	}
	return nil
}

func (ctrl *Controller) DeleteNationality(ctx context.Context, code string) error {
	session, ctx, err := ctrl.store.StartSession(ctx)
	if err != nil {
		return err
	}
	var sessionErr error
	defer func() { session.Close(sessionErr) }()
	repo := ctrl.store.Nationalities()
	sessionErr = repo.Delete(ctx, code)
	if sessionErr != nil {
		return sessionErr
	}
	return nil
}
