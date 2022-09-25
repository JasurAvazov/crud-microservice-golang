package controller

import (
	"apelsin/models"
	"context"
)

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
