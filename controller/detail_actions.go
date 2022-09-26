package controller

import (
	"apelsin/models"
	"context"
)

func (ctrl *Controller) CreateDetail(ctx context.Context, model models.Detail) error {
	var sessionErr error
	session, ctx, err := ctrl.store.StartSession(ctx)
	if err != nil {
		return err
	}
	defer func() { session.Close(sessionErr) }()
	repo := ctrl.store.Detail()

	if _, sessionErr = repo.Create(ctx, model); sessionErr != nil {
		return sessionErr
	}
	return nil
}

func (ctrl *Controller) GetDetail(ctx context.Context, code string) (models.Detail, error) {
	repo := ctrl.store.Detail()
	details, err := repo.Read(ctx, code)
	if err != nil {
		return models.Detail{}, err
	}
	return details, nil
}

func (ctrl *Controller) ListDetail(ctx context.Context) ([]models.Detail, error) {
	repo := ctrl.store.Detail()
	details, err := repo.ReadAll(ctx)
	if err != nil {
		return nil, err
	}
	return details, nil
}

func (ctrl *Controller) UpdateDetail(ctx context.Context, model models.Detail) error {
	var sessionErr error
	session, ctx, err := ctrl.store.StartSession(ctx)
	if err != nil {
		return err
	}
	defer func() { session.Close(sessionErr) }()
	repo := ctrl.store.Detail()

	if sessionErr = repo.Update(ctx, model); sessionErr != nil {
		return sessionErr
	}
	return nil
}

func (ctrl *Controller) DeleteDetail(ctx context.Context, id string) error {
	session, ctx, err := ctrl.store.StartSession(ctx)
	if err != nil {
		return err
	}
	var sessionErr error
	defer func() { session.Close(sessionErr) }()
	repo := ctrl.store.Detail()
	sessionErr = repo.Delete(ctx, id)
	if sessionErr != nil {
		return sessionErr
	}
	return nil
}
